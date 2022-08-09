package model

import (
	"gorm.io/gorm"
)

//Wallet - defines our wallet structure
type Wallet struct {
	gorm.Model
	TransactionType string
	Amount          float64
	WalletID        int
}

type walletService struct {
	DB *gorm.DB
}

type WalletService interface {
	GetWalletByID(ID uint) (Wallet, error)
	GetWallets() (*[]Wallet, error)
	GetWalletByUserID(userID uint) (Wallet, error)
	PostWallet(wallet Wallet) (Wallet, error)
	PutWallet(ID uint, newWallet Wallet) (Wallet, error)
	DeleteWallet(ID uint) (Wallet, error)
}

//GetWalletByID - retrieves comments by their ID from the database
func (s *walletService) GetWalletByID(ID uint) (Wallet, error) {
	var wallet Wallet
	if result := s.DB.First(&wallet); result.Error != nil {

	}
	return wallet, nil
}

//GetWallets - Get all transactions from the database
func (s *walletService) GetWallets() (*[]Wallet, error) {
	var wallet []Wallet
	if result := s.DB.Find(&wallet); result.Error != nil {

	}
	return &wallet, nil
}

//GetWalletByUserID- Get wallet by user id from the database
func (s *walletService) GetWalletByUserID(userID uint) (Wallet, error) {
	var wallet Wallet
	if result := s.DB.Where("UserID = ?", userID).First(&wallet); result.Error != nil {

	}
	return wallet, nil
}

//PostWallet - add a new transaction to the database
func (s *walletService) PostWallet(wallet Wallet) (Wallet, error) {

	if result := s.DB.Save(&wallet); result.Error != nil {

	}
	return wallet, nil
}

//PutWallet - Update a row from the database
func (s *walletService) PutWallet(ID uint, newWallet Wallet) (Wallet, error) {
	var wallet Wallet
	wallet, err := s.GetWalletByID(ID)
	if err != nil {
		return Wallet{}, err
	}
	if result := s.DB.Model(&wallet).Updates(newWallet); result.Error != nil {

		return Wallet{}, err
	}
	return wallet, nil
}

//DeleteWallet - Delete a row from database
func (s *walletService) DeleteWallet(ID uint) (Wallet, error) {
	var wallet Wallet
	if result := s.DB.Delete(&wallet, ID); result.Error != nil {

	}
	return wallet, nil
}

func NewWalletService(db *gorm.DB) *walletService {
	return &walletService{
		DB: db,
	}
}
