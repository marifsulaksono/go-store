package utils

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"gostore/entity"
	"net/http"
	"os"
	"time"
)

const (
	MidtransKey  = "MIDTRANS_SERVER_KEY"
	MidtransLink = "MIDTRANS_SNAP_SANDBOX_LINK"
)

func PaymentGenerator(ctx context.Context, transaction entity.Transaction, user entity.UserResponse, items []entity.ItemDetails) (string, error) {
	paymentRequest := entity.PaymentRequestPayload{
		TransactionDetails: entity.TransactionDetail{
			OrderId:  transaction.Id,
			GrossAmt: transaction.Total,
		},
		ItemDetail: items,
		Expiry: &entity.ExpiryDetails{
			Start:    time.Now().Format("2006-01-02 15:04:05 +0700"),
			Duration: 24,
			Unit:     "hours",
		},
		CustomerDetail: &entity.CostumerDetails{
			FirstName: "buyer - ",
			LastName:  user.Name,
			Email:     user.Email,
			Phone:     user.Phonenumber,
		},
	}

	payloadRequest, err := json.Marshal(paymentRequest)
	if err != nil {
		return "", err
	}

	authString := base64.StdEncoding.EncodeToString([]byte(os.Getenv(MidtransKey) + ":"))
	request, err := http.NewRequest(http.MethodPost, os.Getenv(MidtransLink), bytes.NewBuffer(payloadRequest))
	if err != nil {
		return "", err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+authString)

	// hit midtrans API enpoint with the prepared request
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusCreated {
		return "", errors.New("payment doesn't created")
	}

	var responseJSON map[string]any
	if err := json.NewDecoder(response.Body).Decode(&responseJSON); err != nil {
		return "", err
	}

	return fmt.Sprint(responseJSON["redirect_url"]), nil
}
