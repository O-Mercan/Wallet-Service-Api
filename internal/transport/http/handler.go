package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/o-mercan/Wallet-Service-Api/internal/model"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Router  *mux.Router
	Service *model.TransactionService
}

func NewHandler(service *model.TransactionService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SetupRoutes() {
	log.Info("Routes are setting")
	h.Router = mux.NewRouter()

	h.Router.Handlefunc("api/transactions/{id}", h.GetTransactionByID).Methods("GET")
	h.Router.Handlefunc("api/transactions/{id}", h.GetTransactionByID).Methods("GET")
	h.Router.Handlefunc("api/transactions", h.GetTransactionsByWalletID).Methods("GET")
	h.Router.Handlefunc("api/transactions/{id}", h.PostTransaction).Methods("POST")
	h.Router.Handlefunc("api/transactions/{id}", h.PutTransaction).Methods("PUT")
	h.Router.Handlefunc("api/transactions/{id}", h.DeleteTransaction).Methods("DELETE")

}

func (h *Handler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error("Unable to parse UINT from ID, GetTransaction handler")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)

	}
	transaction, err := h.Service.GetTransactionByID(uint(i))
	if err != nil {
		log.Error("Error retrieving transaction By ID. GetTransaction handler")
		sendErrorResponse(w, "Error retrieving transaction By ID", err)
		return
	}

	if err := sendOkResponse(w, transaction); err != nil {
		log.Panic(err)
	}
}

func (h *Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {

	transaction, err := h.Service.GetTransactions()
	if err != nil {
		log.Error("Failed to get transaction, GetTransaction handler")
		sendErrorResponse(w, "Failed to get transactions", err)
		return
	}

	if err := sendOkResponse(w, transaction); err != nil {
		log.Panic(err)
	}
}

func (h *Handler) GetTransactionsByWalletID(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) PostTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction model.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		log.Error("Error retrieving transaction By ID. PostTransaction handler")
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}
	transaction, err := h.Service.PostTransaction(transaction)
	if err != nil {
		log.Error("Error retrieving transaction By ID. PostTransaction handler")
		sendErrorResponse(w, "Failed to create new transaction", err)
	}

	if err := sendOkResponse(w, transaction); err != nil {
		log.Panic(err)
	}
}

func (h *Handler) PutTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error("Unable to parse UINT from ID. PutTranaction Handler")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
	}
	var transaction model.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		log.Error("Faile to decode json body. PutTransaction Handler")
		sendErrorResponse(w, "Faile to decode json body", err)
	}

	transaction2, err := h.Service.PutTransaction(uint(i), transaction)
	if err != nil {
		log.Error("Failed to update a transaction PutTransaction Handler")
		sendErrorResponse(w, "Failed to update a transaction", err)
	}

	if err := sendOkResponse(w, transaction2); err != nil {
		log.Panic(err)
	}
}

func (h *Handler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error("convert Error, DeleteTransaction handler")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
	}

	transaction, err := h.Service.DeleteTransaction(uint(i))
	if err != nil {
		log.WithFields(log.Fields{
			"Function": "DeleteTransaction",
		}).Error("ID doesn't exist")
		sendErrorResponse(w, "Failed to delete a transaction", err)
	}

	if err := sendOkResponse(w, transaction); err != nil {
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
