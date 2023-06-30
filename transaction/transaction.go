package transaction

import (
	"encoding/json"
	"fmt"
	"gostore/config"
	"gostore/entity"
	"gostore/helper"
	"net/http"
)

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var transaction []entity.TransactionResponse
		err := config.DB.Preload("TransactionItem").Find(&transaction).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, transaction, "Success get all transaction data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func GetTransactionById(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, s := helper.IdVarsMux(w, r)
		if !s {
			return
		}

		var transaction entity.Transaction
		err := config.DB.Where("id = ?", id).Preload("TransactionItem").Preload("TransactionItem.Item").First(&transaction).Error
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success get transaction item %d", transaction.Id)
		helper.ResponseWrite(w, transaction, message)
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusInternalServerError)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var transaction entity.Transaction
		err := json.NewDecoder(r.Body).Decode(&transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if err := config.DB.Create(&transaction).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, transaction.UserId, "Success create transaction!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		id, s := helper.IdVarsMux(w, r)
		if !s {
			return
		}

		var transaction entity.Transaction
		err := json.NewDecoder(r.Body).Decode(&transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		trx := entity.Transaction{}
		if err := config.DB.Where("id = ?", id).First(&trx).Error; err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := config.DB.Model(&trx).Updates(transaction).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, transaction.UserId, "Success update transaction!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		id, s := helper.IdVarsMux(w, r)
		if !s {
			return
		}

		transaction := entity.Transaction{}
		if err := config.DB.Where("id = ?", id).First(&transaction).Error; err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := config.DB.Where("id = ?", id).Delete(&transaction).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, transaction.UserId, "Success delete transaction!")
		return
	}
	http.Error(w, "Method isn't valid", http.StatusBadRequest)
}

// {
//     "id": 1,
//     "date": "2023-06-29T19:20:00Z",
//     "total": 3000000,
//     "status": "unpaid",
//     "user_id": 3
// }
