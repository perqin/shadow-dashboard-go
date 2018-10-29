package db

import (
	"github.com/jinzhu/gorm"
	"github.com/perqin/shadow-dashboard-go/models"
)

var Db *gorm.DB

func OpenDb() (err error) {
	Db, err = gorm.Open("sqlite3", "data/main.db")
	if err != nil {
		return
	}
	Db.AutoMigrate(&models.Subscription{})
	Db.AutoMigrate(&models.Node{})
	return
}

func CloseDb() {
	if Db != nil {
		Db.Close()
	}
}
