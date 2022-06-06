package db

import (
	"../model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	dsn := "host=localhost user=postgres password=1 dbname=SmartPostTest port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Print("DB connected\n")
	}
	err = db.AutoMigrate(&model.Admin{})
	if err != nil {
		return
	}

}

func DbManager() *gorm.DB {
	return db
}
