package http

import (
	"github.com/gorilla/mux"
	wallet_service "github.com/o-mercan/Wallet-Service-Api/internal/wallet"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Router  *mux.Router
	Service *wallet_service.Service
}

func NewHandler(service *wallet_service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SetUpRoutes() {
	log.Info("Routes are setting")
	h.Router = mux.NewRouter().StrictSlash(true)

	h.Router.HandleFunc("/api/v1/wallets/add", h.AddNewWallet).Methods("POST")
	h.Router.HandleFunc("/api/v1/wallets/{id}/deposit", h.DepositWallet).Methods("POST")
	h.Router.HandleFunc("/api/v1/wallets/{id}/withdraw", h.WithdrawWallet).Methods("POST")
	h.Router.HandleFunc("/api/v1/wallets/balance/{id}", h.GetCurrentBalance).Methods("GET")
	h.Router.HandleFunc("/api/v1/transactions/reports", h.GetTransactionsReport).Methods("GET")
	h.Router.HandleFunc("/api/v1/wallets/{id}", h.GetWalletByID).Methods("GET")
	h.Router.HandleFunc("/api/v1/users", h.GetUsers).Methods("GET")
	h.Router.HandleFunc("/api/v1/users/wallets/{id}", h.GetWalletsByUsersID).Methods("GET")
	h.Router.HandleFunc("/api/v1/user/{id}", h.GetUserByID).Methods("GET")

}
