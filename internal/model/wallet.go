package model

import (
	"github.com/jinzhu/gorm"
)

//Wallet - defines our wallet structure
type Wallet struct {
	gorm.Model
	TransactionType string
	Amount          float64
	WalletID        int
}

type Service struct {
	DB *gorm.DB
}

type WalletService interface {
	//GetWalletByID
	//GetWallets
	//PostWallet
	//PutWallet
	//DeleteWallet
}

//GetWalletByID - retrieves comments by their ID from the database
func GetWalletByID(ID uint) (Wallet, error) {
	return Wallet{}, nil
}

//GetWallets - Get all transactions from the database
func GetWallets() (Wallet, error) {
	return Wallet{}, nil
}

//PostWallet - add a new transaction to the database
func PostWallet(wallet Wallet) (Wallet, error) {
	return Wallet{}, nil
}

//PutWallet - Update a row from the database
func PutWallet(ID uint, newWallet Wallet) (Wallet, error) {
	return Wallet{}, nil
}

//DeleteWallet - Delete a row from database
func DeleteWallet(ID uint) (Wallet, error) {
	return Wallet{}, nil
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
