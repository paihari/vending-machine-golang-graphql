package base

import (
	"os"
	"github.com/paihari/vending-machine-golang-graphql/graph/model"
	pg "github.com/go-pg/pg/v10"

)

func GetClientByName(clientName string) model.Client {

	var client model.Client
	GetDb().Model(&client).Where("name = ?", clientName).Select()
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

