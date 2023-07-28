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
	result, err := tr.Service.GetTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, result, "Success get all transactions")
}

func (tr *TransactionController) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		result, err := tr.Service.GetTransactionById(id)
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
	userId := ctx.Value(middleware.GOSTORE_USERID).(int)

	var transaction entity.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := tr.Service.CreateTransaction(userId, &transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Success create new transaction by user %d!", userId)
	helper.ResponseWrite(w, result, message)
}

// func (tr *TransactionController) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	userId := ctx.Value(middleware.GOSTORE_USERID).(int)
// 	if id, s := helper.IdVarsMux(w, r); s {
// 		var transaction entity.Transaction
// 		err := json.NewDecoder(r.Body).Decode(&transaction)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		defer r.Body.Close()

// 		fmt.Println(transaction.UpdateAt)
// 		if err := tr.Service.UpdateTransaction(id, &transaction, userId); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		data := map[string]any{
// 			"total":      transaction.Total,
// 			"status":     transaction.Status,
// 			"updated_at": transaction.UpdateAt,
// 		}
// 		message := fmt.Sprintf("Success update transaction %d!", id)
// 		helper.ResponseWrite(w, data, message)
// 	}
// }

func (tr *TransactionController) SoftDeleteTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(middleware.GOSTORE_USERID).(int)

	if id, s := helper.IdVarsMux(w, r); s {
		if err := tr.Service.SoftDeleteTransaction(id, userId); err != nil {
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
	fmt.Println(userRole)

	if userRole != "Admin" {
		http.Error(w, "Access denied!", http.StatusUnauthorized)
		return
	}

	if id, s := helper.IdVarsMux(w, r); s {
		if err := tr.Service.RestoreDeletedTransaction(id); err != nil {
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
		if _, err := tr.Service.GetTransactionById(id); err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := tr.Service.DeleteTransaction(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success delete permanently transaction %d!", id)
		helper.ResponseWrite(w, id, message)
	}
}
