package database

import (
	"fmt"

	"github.com/o-mercan/Wallet-Service-Api/config.go"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	log.Info("Setting up new database")

	dbConfig := config.NewDBConfig()

	dbUserName := dbConfig.DBUsername
	dbPassword := dbConfig.DBPassword
	dbHost := dbConfig.DBHost
	dbTable := dbConfig.DBTable
	dbPort := dbConfig.DBPort

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUserName, dbPassword, dbTable, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
}
