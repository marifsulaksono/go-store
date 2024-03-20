package utils

import (
	"context"
	"fmt"
	"gostore/entity"
	"gostore/utils/helper"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

const (
	SMTP_HOST = "smtp.gmail.com"
	SMTP_PORT = 587
)

func EmailSender(ctx context.Context, receiver string, subject string, message string) error {
	log.Println(message)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("SMTP_SENDER_NAME"))
	mailer.SetHeader("To", receiver)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(SMTP_HOST, SMTP_PORT, os.Getenv("SMTP_AUTH_USER"), os.Getenv("SMTP_AUTH_PASS"))
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	log.Println("Success send mail...")
	return nil
}

func TransactionMailBuilder(ctx context.Context, tx entity.Transaction) error {
	receiver := ctx.Value(helper.GOSTORE_USEREMAIL).(string)
	message := fmt.Sprintf(`Hi, %s <br />
	Berikut adalah informasi pembayaran untuk transakti Anda: <br />
	Link pembayaran : %s <br />
	Total : %d <br /><br />
	
	Mohon untuk dapat segera menyelesaikan transaksi Anda.`, ctx.Value(helper.GOSTORE_USERNAME), tx.PaymentUrl, tx.Total)

	err := EmailSender(ctx, receiver, "Transaction reminder", message)
	return err
}
