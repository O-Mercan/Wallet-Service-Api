package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/o-mercan/Wallet-Service-Api/internal/model"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) AddNewWallet(w http.ResponseWriter, r *http.Request) {
	var wallet model.Wallet
	if err := json.NewDecoder(r.Body).Decode(&wallet); err != nil {
		log.Error("Error retrieving Wallet By ID. AddNewWallet handler")
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}
	wallet, err := h.Service.AddNewWallet(wallet)
	if err != nil {
		log.Error("Error retrieving Wallet By ID. AddNewWallet handler")
		sendErrorResponse(w, "Failed to create new wallet", err)
	}

	if err := sendOkResponse(w, wallet); err != nil {
		log.Panic(err)
	}
}

func (h *Handler) DepositWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error("Unable to parse UINT from ID. DepositWallet Handler")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
	}
	var wlt model.Wallet
	if err := json.NewDecoder(r.Body).Decode(&wlt); err != nil {
		log.Error("Faile to decode json body. DepositWallet Handler")
		sendErrorResponse(w, "Faile to decode json body", err)
	}

	wallet, err := h.Service.DepositWallet(uint(i), wlt)
	if err != nil {
		log.Error("Failed to update a wallet. DepositWallet Handler")
		sendErrorResponse(w, "Failed to update a wallet", err)
	}

	if err := sendOkResponse(w, wallet); err != nil {
		log.Panic(err)
	}
}

func (h *Handler) WithdrawWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error("Unable to parse UINT from ID. DepositWallet Handler")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
	}
	var wlt model.Wallet
	if err := json.NewDecoder(r.Body).Decode(&wlt); err != nil {
		log.Error("Faile to decode json body. DepositWallet Handler")
		sendErrorResponse(w, "Faile to decode json body", err)
	}

	wallet, err := h.Service.DepositWallet(uint(i), wlt)
	if err != nil {
		log.Error("Failed to update a wallet. DepositWallet Handler")
		sendErrorResponse(w, "Failed to update a wallet", err)
	}

	if err := sendOkResponse(w, wallet); err != nil {
		log.Panic(err)
	}
}

func (h *Handler) GetCurrentBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error("Unable to parse UINT from ID, GetCurrentBalance handler")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)

	}
	wallet, err := h.Service.GetCurrentBalance(uint(i))
	if err != nil {
		log.Error("Error retrieving Wallet By ID. GetCurrentBalance handler")
		sendErrorResponse(w, "Error retrieving Wallet By ID", err)
		return
	}

	if err := sendOkResponse(w, wallet); err != nil {
		log.Panic(err)
	}
}

func (h *Handler) GetTransactionsReport(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetWalletByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error("Unable to parse UINT from ID, GetWallet handler")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)

	}
	wlt, err := h.Service.GetWalletByID(uint(i))
	if err != nil {
		log.Error("Error retrieving Wallet By ID. GetWallet handler")
		sendErrorResponse(w, "Error retrieving Wallet By ID", err)
		return
	}

	if err := sendOkResponse(w, wlt); err != nil {
		log.Panic(err)
	}
}

func (h *Handler) GetWalletsByUsersID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error("Unable to parse UINT from ID, GetWalletsByUserID handler")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)

	}
	user, err := h.Service.GetWalletsByUsersID(uint(i))
	if err != nil {
		log.Error("Error retrieving Wallet By ID. GetWalletsByUserID handler")
		sendErrorResponse(w, "Error retrieving Wallet By ID", err)
		return
	}

	if err := sendOkResponse(w, user); err != nil {
		log.Panic(err)
	}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	user, err := h.Service.GetUsers()
	if err != nil {
		log.Error("Failed to get products, GetUsers handler")
		sendErrorResponse(w, "Failed to get wallets", err)
		return
	}

	if err := sendOkResponse(w, user); err != nil {
		log.Panic(err)
	}
}
func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error("Unable to parse UINT from ID, GetUserByID handler")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)

	}
	wlt, err := h.Service.GetUserByID(uint(i))
	if err != nil {
		log.Error("Error retrieving Wallet By ID. GetUserByID handler")
		sendErrorResponse(w, "Error retrieving Wallet By ID", err)
		return
	}

	if err := sendOkResponse(w, wlt); err != nil {
		log.Panic(err)
	}
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
