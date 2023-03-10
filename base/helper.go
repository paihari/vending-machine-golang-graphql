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


func GetRenterByName(renterName string, db *pg.DB ) model.Renter {

	var renter model.Renter
	db.Model(&renter).Where("name = ?", renterName).Select()
	return renter
}

func GetRenterByUUID(uuid string, db *pg.DB ) model.Renter {

	var renter model.Renter
	db.Model(&renter).Where("uuid = ?", uuid).Select()
	return renter
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



func GetCloudEstatePoliciesByName(policyNames []*string, db *pg.DB ) []*model.CloudEstatePolicy {
	var cloudEstatePolicys []*model.CloudEstatePolicy

	db.Model(&cloudEstatePolicys).Where("name = ?", policyNames[0]).Select()
	return cloudEstatePolicys
}

func GetUserByName(userName string, db *pg.DB) model.User {

	var user model.User
	db.Model(&user).Where("name = ?", userName).Select()
	return user
}

func GetUserByUUID(uuid string, db *pg.DB) model.User {

	var user model.User
	db.Model(&user).Where("uuid = ?", uuid).Select()
	return user
}

func GetResidentByName(residentName string, db *pg.DB) model.Resident {

	var resident model.Resident
	db.Model(&resident).Where("name = ?", residentName).Select()
	return resident
}

func GetResidentByUUID(uuid string, db *pg.DB) model.Resident {

	var resident model.Resident
	db.Model(&resident).Where("uuid = ?", uuid).Select()
	return resident
}

func GetTagByName(tagName string, db *pg.DB) model.Tag {

	var tag model.Tag
	db.Model(&tag).Where("name = ?", tagName).Select()
	return tag
}

func GetTagByUUID(uuid string, db *pg.DB) model.Tag {

	var tag model.Tag
	db.Model(&tag).Where("uuid = ?", uuid).Select()
	return tag
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