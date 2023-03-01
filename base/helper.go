package base

import (
	"os"
	pg "github.com/go-pg/pg/v10"
	"github.com/paihari/vending-machine-golang-graphql/graph/model"
)

func GetClientByName(clientName string, db *pg.DB ) model.Client {

	var client model.Client
	db.Model(&client).Where("name = ?", clientName).Select()
	return client
}

func GetCloudProviderByName(cloudProviderName string, db *pg.DB) model.CloudProvider {

	var cloudProvider model.CloudProvider
	db.Model(&cloudProvider).Where("name = ?", cloudProviderName).Select()
	return cloudProvider
}

func GetClassByName(className string, db *pg.DB) model.Class {

	var class model.Class
	db.Model(&class).Where("name = ?", class).Select()
	return class
}

func GetStageByName(stageName string, db *pg.DB) model.Stage {

	var stage model.Stage
	db.Model(&stage).Where("name = ?", stage).Select()
	return stage
}

func GetDb() *pg.DB {

	connStr := os.Getenv("DB_URL")
	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)
	defer db.Begin()
	return db;
}