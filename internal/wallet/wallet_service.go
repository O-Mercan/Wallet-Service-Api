package wallet

import (
	"encoding/json"
	"fmt"
	model "github.com/o-mercan/Wallet-Service-Api/internal/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type Service struct {
	DB *gorm.DB
}

type WalletService interface {
	AddNewWallet(wallet model.Wallet) (model.Wallet, error)
	DepositWallet(ID uint, newWallet model.Wallet) (model.Wallet, error)
	WithdrawWallet(ID uint, newWallet model.Wallet) (model.Wallet, error)
	GetCurrentBalance(ID uint) (model.Wallet, error)
	GetTransactionsReport() (*[]model.Transaction, error)
	GetWalletByID(ID uint) (model.Wallet, error)
	GetUsers() (*[]model.User, error)
	GeWalletsByUsersID(ID uint) (model.Wallet, error)
	GetUserByID(ID uint) (model.User, error)
}

func (s *Service) AddNewWallet(wallet model.Wallet) (model.Wallet, error) {
	if result1 := s.DB.Limit(1).FirstOrCreate(&wallet, model.Wallet{UserID: wallet.UserID, Currency: wallet.Currency}); result1.Error != nil {
		log.WithFields(log.Fields{
			"Function": "AddNewWallet",
		}).Error(result1.Error.Error())
		fmt.Println("You have already this account, Choose different Currency")
		return model.Wallet{}, nil
	}

	s.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&wallet)
	return wallet, nil
}

func (s *Service) DepositWallet(ID uint, newWallet model.Wallet) (model.Wallet, error) {
	var wallet model.Wallet
	var transaction model.Transaction
	wlt, err := s.GetWalletByID(ID)
	if err != nil {
		return model.Wallet{}, err
	}
	depositAmount := newWallet.Amount
	if result1 := s.DB.Model(&wlt).Update("amount", gorm.Expr("amount + ?", depositAmount)); result1.Error != nil {
		return model.Wallet{}, nil
	}
	s.DB.Create(&transaction)
	return wallet, nil
}

func (s *Service) WithdrawWallet(ID uint, newWallet model.Wallet) (model.Wallet, error) {
	var wallet model.Wallet
	var transaction model.Transaction
	wlt, err := s.GetWalletByID(ID)
	if err != nil {
		return model.Wallet{}, err
	}
	withdrawAmount := newWallet.Amount
	if result1 := s.DB.Model(&wlt).Update("amount", gorm.Expr("amount - ?", withdrawAmount)); result1.Error != nil {
		return model.Wallet{}, nil
	}
	s.DB.Create(transaction)

	return wallet, nil
}

func (s *Service) GetCurrentBalance(ID uint) (model.Wallet, error) {
	var wallet model.Wallet
	if result := s.DB.First(&wallet, ID); result.Error != nil {
		return model.Wallet{}, nil
	}

	return model.Wallet{Amount: wallet.Amount}, nil
}

func (s *Service) GetTransactionsReport() (*[]model.Transaction, error) {
	var transaction []model.Transaction
	if result := s.DB.Preload("Wallet").Preload("Wallet.UserID").Find(&transaction); result.Error != nil {
		return &transaction, nil
	}

	return &transaction, nil
}

func (s *Service) GetWalletByID(ID uint) (model.Wallet, error) {
	var wallet model.Wallet
	if result := s.DB.First(&wallet, ID); result.Error != nil {
		log.WithFields(log.Fields{
			"WalletID":  wallet.ID,
			"requestID": ID,
			"Function":  "GetWallet",
		}).Error(result.Error.Error())
		return model.Wallet{}, nil
	}
	return wallet, nil
}

func (s *Service) GetWalletsByUsersID(ID uint) (model.User, error) {
	var user model.User
	if result := s.DB.First(&user, ID); result.Error != nil {
		return model.User{}, nil
	}

	return model.User{Wallets: user.Wallets}, nil
}

func (s *Service) GetUserByID(ID uint) (model.User, error) {
	var user model.User
	if result := s.DB.First(&user, ID); result.Error != nil {
		log.WithFields(log.Fields{
			"UserID":    user.ID,
			"requestID": ID,
			"Function":  "GetUserByID",
		}).Error(result.Error.Error())
		return model.User{}, nil
	}
	return user, nil
}

func (s *Service) GetUsers() (*[]model.User, error) {
	var user []model.User
	if result := s.DB.Find(&user); result.Error != nil {
		log.WithFields(log.Fields{
			"Function": "GetUsers",
		}).Error(result.Error.Error())
		return &user, nil
	}
	return &user, nil
}

type Response struct {
	Message string
	Error   string
}

func sendOkResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		log.Panic(err)
	}
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
