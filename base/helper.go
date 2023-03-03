package base

import (
	"os"
	pg "github.com/go-pg/pg/v10"
	"github.com/paihari/vending-machine-golang-graphql/graph/model"
)




func GetCloudEstateByName(name string, db *pg.DB ) model.CloudEstate {
	var cloudEstate model.CloudEstate
	db.Model(&cloudEstate).Where("name = ?", name).Select()
	return cloudEstate
}

func GetCloudEstateByUUID(uuid string, db *pg.DB ) model.CloudEstate {
	var cloudEstate model.CloudEstate
	db.Model(&cloudEstate).Where("uuid = ?", uuid).Select()
	return cloudEstate
}


func GetFederalByUUID(uuid string, db *pg.DB ) model.Federal {
	var federal model.Federal
	db.Model(&federal).Where("uuid = ?", uuid).Select()
	return federal
}


func GetFederalByName(name string, db *pg.DB ) model.Federal {
	var federal model.Federal
	db.Model(&federal).Where("name = ?", name).Select()
	return federal
}


func GetClientByName(clientName string, db *pg.DB ) model.Client {

	var client model.Client
	db.Model(&client).Where("name = ?", clientName).Select()
	return client
}

func GetClientByUUID(uuid string, db *pg.DB ) model.Client {

	var client model.Client
	db.Model(&client).Where("uuid = ?", uuid).Select()
	return client
}

func GetClassByName(className string, db *pg.DB ) model.Class {

	var class model.Class
	db.Model(&class).Where("name = ?", className).Select()
	return class
}

func GetClassByUUID(uuid string, db *pg.DB ) model.Class {

	var class model.Class
	db.Model(&class).Where("uuid = ?", uuid).Select()
	return class
}

func GetStageByName(stageName string, db *pg.DB) model.Stage {

	var stage model.Stage
	db.Model(&stage).Where("name = ?", stageName).Select()
	return stage
}

func GetStageByUUID(uuid string, db *pg.DB) model.Stage {

	var stage model.Stage
	db.Model(&stage).Where("uuid = ?", uuid).Select()
	return stage
}


func GetCloudProviderByName(cloudProviderName string, db *pg.DB) model.CloudProvider {

	var cloudProvider model.CloudProvider
	db.Model(&cloudProvider).Where("name = ?", cloudProviderName).Select()
	return cloudProvider
}

func GetCloudProviderByUUID(uuid string, db *pg.DB) model.CloudProvider {

	var cloudProvider model.CloudProvider
	db.Model(&cloudProvider).Where("uuid = ?", uuid).Select()
	return cloudProvider
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