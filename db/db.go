package db

import (
	envPostgres "app/config/postgres"
	"app/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		envPostgres.EnvHost(),
		envPostgres.EnvUser(),
		envPostgres.EnvPassword(),
		envPostgres.EnvDbName(),
		envPostgres.EnvPort(),
	)
	//dsn := envPostgres.EnvCloudURL()

	fmt.Print(dsn)

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

func GetDbManager() *gorm.DB {
	return db
}
