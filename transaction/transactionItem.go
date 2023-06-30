package transaction

import (
	"encoding/json"
	"fmt"
	"gostore/config"
	"gostore/entity"
	"gostore/helper"
	"net/http"
)

func GetTransactionItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var transactionItem []entity.TransactionItemResponse
		err := config.DB.Preload("Item").Find(&transactionItem).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, transactionItem, "Success get all transaction item!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func GetTransactionItemId(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, s := helper.IdVarsMux(w, r)
		if !s {
			return
		}

		var transactionItem entity.TransactionItemResponse
		err := config.DB.Where("id = ?", id).Preload("Item").First(&transactionItem).Error
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success get transaction item %d", transactionItem.Id)
		helper.ResponseWrite(w, transactionItem, message)
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func CreateTransactionItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var transactionItem entity.TransactionItem
		err := json.NewDecoder(r.Body).Decode(&transactionItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if err := config.DB.Create(&transactionItem).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, transactionItem, "Success create transaction item!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func UpdateTransactionItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		id, s := helper.IdVarsMux(w, r)
		if !s {
			return
		}

		var transactionItem entity.TransactionItem
		err := json.NewDecoder(r.Body).Decode(&transactionItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		transaction_item := entity.TransactionItem{}
		if err := config.DB.Where("id = ?", id).First(&transaction_item).Error; err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := config.DB.Model(&transaction_item).Updates(transactionItem).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, transaction_item, "Success update transaction item!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func DeleteTransactionItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		id, s := helper.IdVarsMux(w, r)
		if !s {
			return
		}

		checkItem := entity.TransactionItem{}
		if err := config.DB.Where("id = ?", id).First(&checkItem).Error; err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := config.DB.Where("id = ?", id).Delete(&checkItem).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, id, "Success delete transaction item!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

// {
// "transaction_id": 1,
// "item_id": 4,
// "qty": 3,
// "price": 1000000,
// "subtotal": 1000000
// }
