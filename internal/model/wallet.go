package model

import (
	"gorm.io/gorm"
)

//Wallet - defines our wallet structure
type Wallet struct {
	gorm.Model
	Amount   float64
	Currency string
	UserID   uint
}

type User struct {
	gorm.Model
	Name    string
	Surname string
	Email   string
	Wallets []Wallet `gorm:"foreignKey:UserID"`
}
