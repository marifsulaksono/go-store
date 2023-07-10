package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/service"
	"net/http"
)

func GetTransactionItem(w http.ResponseWriter, r *http.Request) {
	result, err := service.GetTransactionItem()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, result, "Success get all transaction item")
}

func GetTransactionItemId(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		result, err := service.GetTransactionItemById(id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success get transaction item %d", id)
		helper.ResponseWrite(w, result, message)
	}
}

func CreateTransactionItem(w http.ResponseWriter, r *http.Request) {
	var transactionItem entity.TransactionItem
	err := json.NewDecoder(r.Body).Decode(&transactionItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := service.CreateTransactionItem(transactionItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, transactionItem, "Success create new transaction item")
}

func UpdateTransactionItem(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		var transactionItem entity.TransactionItem
		err := json.NewDecoder(r.Body).Decode(&transactionItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if err := service.UpdateTransactionItem(id, transactionItem); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success update transaction item %d", id)
		helper.ResponseWrite(w, transactionItem, message)
	}
}

func DeleteTransactionItem(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		if _, err := service.GetTransactionItemById(id); err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := service.DeleteTransactionItem(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success delete transaction item %d", id)
		helper.ResponseWrite(w, id, message)
	}
}
