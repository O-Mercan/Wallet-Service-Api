package model

import (
	"gorm.io/gorm"
)

//Transaction - defines our transaction structure
type Transaction struct {
	gorm.Model
	TransactionType string
	Amount          float64
	WalletID        int
}

type transactionService struct {
	DB *gorm.DB
}

type TransactionService interface {
	GetTransactionByID(ID uint) (Transaction, error)
	GetTransactions() (*[]Transaction, error)
	GetTransactionsByWalletID(ID uint) (*[]Transaction, error)
	PostTransaction(transaction Transaction) (Transaction, error)
	PutTransaction(ID uint, NewTransaction Transaction) (Transaction, error)
	DeleteTransaction(ID uint) (Transaction, error)
}

//GetTransactionByID - retrieves comments by their ID from the database
func (s *transactionService) GetTransactionByID(ID uint) (Transaction, error) {
	var transaction Transaction
	if result := s.DB.First(&transaction); result.Error != nil {

	}
	return transaction, nil
}

//GetTransactions - Get all transactions from the database
func (s *transactionService) GetTransactions() (*[]Transaction, error) {
	var transaction []Transaction
	if result := s.DB.Find(&transaction); result.Error != nil {

	}
	return &transaction, nil

}

//GetTransactionsByWalletID - Get transaction by walletid from the database
func (s *transactionService) GetTransactionsByWalletID(ID uint) (*[]Transaction, error) {
	var transaction []Transaction
	if result := s.DB.Where("WalletID = ?", ID).Find(&transaction); result.Error != nil {

	}
	return &transaction, nil
}

//PostTransaction - add a new transaction to the database
func (s *transactionService) PostTransaction(transaction Transaction) (Transaction, error) {

	if result := s.DB.Save(&transaction); result.Error != nil {

	}
	return transaction, nil

}

//PutTransaction - Update a row from the database
func (s *transactionService) PutTransaction(ID uint, newTransaction Transaction) (Transaction, error) {
	var transaction Transaction
	transaction, err := s.GetTransactionByID(ID)
	if err != nil {
		return Transaction{}, err
	}
	if result := s.DB.Model(&transaction).Updates(newTransaction); result.Error != nil {

		return Transaction{}, err
	}
	return transaction, nil

}

//DeleteTransaction - Delete a row from database
func (s *transactionService) DeleteTransaction(ID uint) (Transaction, error) {
	var transaction Transaction
	if result := s.DB.Delete(&transaction, ID); result.Error != nil {

	}
	return transaction, nil

}

func NewTransactionService(db *gorm.DB) *userService {
	return &userService{
		DB: db,
	}
}
