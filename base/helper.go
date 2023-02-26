package base

import (
	"fmt"
	"os"

	pg "github.com/go-pg/pg/v10"
	"github.com/paihari/vending-machine-golang-graphql/graph/model"
)

func GetClientByName(clientName string) model.Client {

	var client model.Client
	GetDb().Model(&client).Where("name = ?", clientName).Select()
	fmt.Println("Client NAme")
	fmt.Println(client)
	return client
}

func GetCloudProviderByName(cloudProviderName string) model.CloudProvider {

	var cloudProvider model.CloudProvider
	GetDb().Model(&cloudProvider).Where("name = ?", cloudProviderName).Select()
	return cloudProvider
}

func GetClassByName(className string) model.Class {

	var class model.Class
	GetDb().Model(&class).Where("name = ?", class).Select()
	return class
}

func GetStageByName(stageName string) model.Stage {

	var stage model.Stage
	GetDb().Model(&stage).Where("name = ?", stage).Select()
	return stage
}

func GetDb() pg.DB {

	connStr := os.Getenv("DB_URL")
	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)
	defer db.Close()
	return *db;
}