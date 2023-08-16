package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/middleware"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, result, "Success get all transactions")
}

func (tr *TransactionController) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		result, err := tr.Service.GetTransactionById(ctx, id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success get transaction item %d", id)
		helper.ResponseWrite(w, result, message)
	}
}

func (tr *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var transaction entity.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := tr.Service.CreateTransaction(ctx, &transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, result, "Success create new transaction")
}

func (tr *TransactionController) SoftDeleteTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(middleware.GOSTORE_USERID).(int)

	if id, s := helper.IdVarsMux(w, r); s {
		if err := tr.Service.SoftDeleteTransaction(ctx, id, userId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success delete transaction %d!", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (tr *TransactionController) RestoreDeletedTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userRole := ctx.Value(middleware.GOSTORE_USERROLE)

	if userRole != "Admin" {
		http.Error(w, "Access denied!", http.StatusUnauthorized)
		return
	}

	if id, s := helper.IdVarsMux(w, r); s {
		if err := tr.Service.RestoreDeletedTransaction(ctx, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success restore transaction %d!", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (tr *TransactionController) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userRole := ctx.Value(middleware.GOSTORE_USERROLE)

	if userRole != "Admin" {
		http.Error(w, "Access denied!", http.StatusUnauthorized)
		return
	}

	if id, s := helper.IdVarsMux(w, r); s {
		if _, err := tr.Service.GetTransactionById(ctx, id); err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := tr.Service.DeleteTransaction(ctx, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success delete permanently transaction %d!", id)
		helper.ResponseWrite(w, id, message)
	}
}
