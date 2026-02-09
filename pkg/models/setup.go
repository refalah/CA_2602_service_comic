package models

import (
	"github.com/refalah/CA_2602_service_comic/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func init(){
	config.Connect()
	db = config.GetDB()

	_ = Comic{}
	_ = CreativeTeam{}
	_ = Creator{}
	_ = Role{}

	// db.AutoMigrate(&Comic{}, &CreativeTeam{}, &Creator{}, &Role{})
}