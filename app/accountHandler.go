package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedramaghasian/hexagonal-in-go/dto"
	"github.com/pedramaghasian/hexagonal-in-go/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h *AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)

		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
func (h *AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]
	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		request.AccountId = accountId
		account, appError := h.service.MakeTransaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}
