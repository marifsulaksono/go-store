package controller

import (
	"encoding/json"
	"gostore/entity"
	"gostore/service"
	"gostore/utils/response"
	"net/http"

	"github.com/gorilla/mux"
)

type TransactionController struct {
	Service service.TransactionService
}

func NewTransactionController(s service.TransactionService) *TransactionController {
	return &TransactionController{
		Service: s,
	}
}

func (tr *TransactionController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		message string
	)

	result, err := tr.Service.GetTransactions(ctx)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if result == nil || len(result) < 1 {
		message = "No results found"
	}

	response.BuildSuccesResponse(w, result, nil, message)
}

func (tr *TransactionController) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	result, err := tr.Service.GetTransactionById(ctx, id)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, result, nil, "")
}

func (tr *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var transaction entity.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	err = tr.Service.CreateTransaction(ctx, &transaction)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success create new transaction")
}
