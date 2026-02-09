package models

import (
	"github.com/refalah/go-comics/pkg/config"
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