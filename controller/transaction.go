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

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	result, err := service.GetTransaction()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, result, "Success get all transaction")
}

func GetTransactionById(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		result, err := service.GetTransactionById(id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success get transaction item %d", id)
		helper.ResponseWrite(w, result, message)
	}
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(middleware.GOSTORE_USERID)

	var transaction entity.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := service.CreateTransaction(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Success create new transaction by user %d!", userId)
	helper.ResponseWrite(w, transaction, message)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		var transaction entity.Transaction
		err := json.NewDecoder(r.Body).Decode(&transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if err := service.UpdateTransaction(id, transaction); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success update transaction %d!", id)
		helper.ResponseWrite(w, transaction, message)
	}
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		if _, err := service.GetTransactionById(id); err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := service.DeleteTransaction(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success delete transaction %d!", id)
		helper.ResponseWrite(w, id, message)
	}
}
