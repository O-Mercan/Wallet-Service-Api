package model

import (
	"gorm.io/gorm"
)

//Wallet - defines our wallet structure
type Wallet struct {
	gorm.Model
	Amount        float64
	Currency      string
	UserID        uint
	TransactionID uint
}

type User struct {
	gorm.Model
	Name    string
	Surname string
	Email   string
	Wallets []Wallet `gorm:"foreignKey:UserID"`
}
type Transaction struct {
	gorm.Model
	Wallets []Wallet `gorm:"foreignKey:TransactionID"`
}
