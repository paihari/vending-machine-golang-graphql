package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"os"

	pg "github.com/go-pg/pg/v10"
	"github.com/paihari/vending-machine-golang-graphql/graph/model"
)

// CreateResident is the resolver for the createResident field.
func (r *mutationResolver) CreateResident(ctx context.Context, input model.NewResident) (*model.Resident, error) {
	connStr := os.Getenv("DB_URL")
	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)
	defer db.Close()

	// client := base.GetClientByName(input.ClientName)
	var client model.Client
	db.Model(&client).Where("name = ?", input.ClientName).Select()

	//cloudProvider := base.GetCloudProviderByName(input.CloudProviderName)

	var cloudProvider model.CloudProvider
	db.Model(&cloudProvider).Where("name = ?", input.CloudProviderName).Select()

	// class := base.GetClassByName(input.ClassName)
	var class model.Class
	db.Model(&class).Where("name = ?", input.ClassName).Select()

	// stage := base.GetStageByName(input.Name)
	var stage model.Stage
	db.Model(&stage).Where("name = ?", input.StageName).Select()

	var createdBy, updatedBy string
	createdBy = "VEND"
	updatedBy = "VEND"

	// residentCid, err := awscompose.CreateResidentAccount(input.Name, input.EmailAddress)
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve labels: %v", err)
	// 	return nil, err
	// }

	residentCid := "SOME CID"

	resident := model.Resident{
		Name:            input.Name,
		Description:     input.Description,
		PurchaseOrderID: input.PurchaseOrderID,
		EmailAddress:    input.EmailAddress,
		Client:          client.Name,
		CloudProvider:   cloudProvider.Name,
		ResidentCid:     residentCid,
		RootCid:         cloudProvider.RootCid,
		Class:           class.Name,
		Stage:           stage.Name,
		CreatedBy:       createdBy,
		UpdatedBy:       updatedBy,
	}

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
