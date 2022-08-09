package database

import (
	"github.com/o-mercan/Wallet-Service-Api/internal/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}, &model.Wallet{}, &model.Transaction{}); err != nil {
		return err
	}
	return nil
}
