package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"log"
	"os"

	pg "github.com/go-pg/pg/v10"
	"github.com/paihari/vending-machine-golang-graphql/base"
	"github.com/paihari/vending-machine-golang-graphql/graph/model"
	"github.com/paihari/vending-machine-golang-graphql/graph/awscompose"
)

// CreateResident is the resolver for the createResident field.
func (r *mutationResolver) CreateResident(ctx context.Context, input model.NewResident) (*model.Resident, error) {
	residentCid, err := awscompose.CreateResidentAccount(input.Name)
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
		return nil, err
	}

	client := base.GetClientByName(input.ClientName)
	cloudProvider := base.GetCloudProviderByName(input.CloudProviderName)
	class := base.GetClassByName(input.ClassName)
	stage := base.GetStageByName(input.Name)

	// TODO: Get the Values user context
	var createdBy, updatedBy string
	createdBy = "VEND"
	updatedBy = "VEND"

	resident := model.Resident{
		Name:            input.Name,
		Description:     input.Description,
		PurchaseOrderID: input.PurchaseOrderID,
		Client:          client.Name,
		CloudProvider:   cloudProvider.Name,
		ResidentCid:     residentCid,
		RootCid:         cloudProvider.RootCid,
		Class:           class.Name,
		Stage:           stage.Name,
		CreatedBy:       createdBy,
		UpdatedBy:       updatedBy,
	}

	connStr := os.Getenv("DB_URL")
	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)
	defer db.Close()

	_, error := db.Model(&resident).Insert()

	if error != nil {
		return nil, fmt.Errorf("error inserting new Resident: %v", error)
	}

	return &resident, nil
}

// Residents is the resolver for the residents field.
func (r *queryResolver) Residents(ctx context.Context) ([]*model.Resident, error) {
	var residents []*model.Resident

	connStr := os.Getenv("DB_URL")
	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	defer db.Close()

	error := db.Model(&residents).Select()
	if error != nil {
		return nil, error
	}

	return residents, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
