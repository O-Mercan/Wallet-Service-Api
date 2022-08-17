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
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/v1/wallets/add", h.AddNewWallet).Methods("POST")
	h.Router.HandleFunc("/api/v1/wallets/deposit/{id}", h.DepositWallet).Methods("POST")
	h.Router.HandleFunc("/api/v1/wallets/withdraw/{id}", h.WithdrawWallet).Methods("POST")
	h.Router.HandleFunc("/api/v1/wallets/balance/{id}", h.GetCurrentBalance).Methods("GET")
	h.Router.HandleFunc("/api/v1/transactions/reports", h.GetTransactionsReport).Methods("GET")
	/////////////////////////////////////////
	h.Router.HandleFunc("/api/v1/wallets/{id}", h.GetWalletByID).Methods("GET")
	//h.Router.HandleFunc("/api/v1/users", h.PostUser).Methods("POST")
	//h.Router.HandleFunc("/api/v1/users", h.GetUsers).Methods("GET")
	//h.Router.HandleFunc("/api/v1/wallets/post", h.PostWallet).Methods("POST")
	h.Router.HandleFunc("/api/v1/wallets/get", h.GetWallets).Methods("GET")

}
