package controller

import (
	"encoding/json"
	"gostore/entity"
	"gostore/helper"
	"gostore/service"
	"net/http"
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
	ctx := r.Context()
	result, err := tr.Service.GetTransactions(ctx)
	if err != nil {
		helper.BuildError(w, err)
		return
	}

	if result == nil || len(result) < 1 {
		helper.BuildResponseSuccess(w, result, nil, "No results found")
		return
	}

	helper.BuildResponseSuccess(w, result, nil, "")
}

func (tr *TransactionController) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		result, err := tr.Service.GetTransactionById(ctx, id)
		if err != nil {
			helper.BuildError(w, err)
			return
		}

		helper.BuildResponseSuccess(w, result, nil, "")
	}
}

func (tr *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var transaction entity.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		helper.BuildError(w, err)
		return
	}
	defer r.Body.Close()

	err = tr.Service.CreateTransaction(ctx, &transaction)
	if err != nil {
		helper.BuildError(w, err)
		return
	}

	helper.BuildResponseSuccess(w, nil, nil, "Success create new transaction")
}
