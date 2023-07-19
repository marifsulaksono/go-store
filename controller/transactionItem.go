package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/service"
	"net/http"
)

type TransactionItemController struct {
	Service service.TransactionItemService
}

func NewTransactionItemController(s service.TransactionItemService) *TransactionItemController {
	return &TransactionItemController{
		Service: s,
	}
}

func (ti *TransactionItemController) GetTransactionItems(w http.ResponseWriter, r *http.Request) {
	result, err := ti.Service.GetTransactionItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, result, "Success get all transaction items")
}

func (ti *TransactionItemController) GetTransactionItemId(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		result, err := ti.Service.GetTransactionItemById(id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success get transaction item %d", id)
		helper.ResponseWrite(w, result, message)
	}
}

func (ti *TransactionItemController) CreateTransactionItem(w http.ResponseWriter, r *http.Request) {
	var transactionItem entity.TransactionItem
	err := json.NewDecoder(r.Body).Decode(&transactionItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := ti.Service.CreateTransactionItem(&transactionItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, transactionItem, "Success create new transaction item")
}

func (ti *TransactionItemController) UpdateTransactionItem(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		var transactionItem entity.TransactionItem
		err := json.NewDecoder(r.Body).Decode(&transactionItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if err := ti.Service.UpdateTransactionItem(id, &transactionItem); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success update transaction item %d", id)
		helper.ResponseWrite(w, transactionItem, message)
	}
}

func (ti *TransactionItemController) SoftDeleteTransactionItem(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		if err := ti.Service.SoftDeleteTransactionItem(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success delete transaction item %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (ti *TransactionItemController) RestoreDeletedTransactionItem(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		if err := ti.Service.RestoreDeletedTransactionItem(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success restore transaction item %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (ti *TransactionItemController) DeleteTransactionItem(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		if err := ti.Service.DeleteTransactionItem(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success delete permanently transaction item %d", id)
		helper.ResponseWrite(w, id, message)
	}
}
