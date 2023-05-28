package database

import (
	"github.com/diiineeei/IntegracaoJson/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var InstanceMySQL *gorm.DB
var InstancePostgres *gorm.DB
var dbError error

func ConnectMySQL(connectionString string) {
	InstanceMySQL, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database MySQL!")
}
func ConnectPostgres(connectionString string) {
	InstancePostgres, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database Postgres!")
}

func Migrate() {
	InstanceMySQL.AutoMigrate(&models.Contacts{}, &models.User{})
	InstancePostgres.AutoMigrate(&models.Contacts{}, &models.User{})
	log.Println("Database Migration Completed!")
}
