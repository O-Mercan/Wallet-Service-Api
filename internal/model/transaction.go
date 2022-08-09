package model

import (
	"github.com/jinzhu/gorm"
)

//Transaction - defines our transaction structure
type Transaction struct {
	gorm.Model
	TransactionType string
	Amount          float64
	WalletID        int
}

type Service struct {
	DB *gorm.DB
}

type TransactionService interface {
	//GetTransactionByID
	//GetTransactions
	//PostTransaction
	//PutTransaction
	//DeleteTransaction
}

//GetTransactionByID - retrieves comments by their ID from the database
func GetTransactionByID(ID uint) (Transaction, error) {
	return Transaction{}, nil
}

//GetTransaction - Get all transactions from the database
func GetTransactions() (Transaction, error) {
	return Transaction{}, nil

}

//PostTransaction - add a new transaction to the database
func PostTransaction(transaction Transaction) (Transaction, error) {
	return Transaction{}, nil

}

//PutTransaction - Update a row from the database
func PutTransaction(ID uint, newTransaction Transaction) (Transaction, error) {
	return Transaction{}, nil

}

//DeleteTransaction - Delete a row from database
func DeleteTransaction(ID uint) (Transaction, error) {
	return Transaction{}, nil

}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
