package app

import (
	"context"
	"fmt"
	"log"
	"time"
)

// @@@SNIPSTART money-transfer-project-template-go-activity-withdraw
func Withdraw(ctx context.Context, data PaymentDetails) (string, error) {
	txn := NrApp.StartTransaction("Withdraw")
	defer txn.End()

	log.Printf("Withdrawing $%d from account %s.\n\n",
		data.Amount,
		data.SourceAccount,
	)

	log.Printf("Waiting 50 seconds...")
	time.Sleep(50 * time.Second)
	log.Printf("Waiting done")

	referenceID := fmt.Sprintf("%s-withdrawal", data.ReferenceID)
	bank := BankingService{"bank-api.example.com"}
	confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount, referenceID)
	return confirmation, err
}

// @@@SNIPEND

// @@@SNIPSTART money-transfer-project-template-go-activity-deposit
func Deposit(ctx context.Context, data PaymentDetails) (string, error) {
	txn := NrApp.StartTransaction("Deposit")
	defer txn.End()
	log.Printf("Depositing $%d into account %s.\n\n",
		data.Amount,
		data.TargetAccount,
	)

	referenceID := fmt.Sprintf("%s-deposit", data.ReferenceID)
	bank := BankingService{"bank-api.example.com"}
	// Uncomment the next line and comment the one after that to simulate an unknown failure
	// confirmation, err := bank.DepositThatFails(data.TargetAccount, data.Amount, referenceID)
	confirmation, err := bank.Deposit(data.TargetAccount, data.Amount, referenceID)
	return confirmation, err
}

// @@@SNIPEND

// @@@SNIPSTART money-transfer-project-template-go-activity-refund
func Refund(ctx context.Context, data PaymentDetails) (string, error) {
	txn := NrApp.StartTransaction("Refund")
	defer txn.End()
	log.Printf("Refunding $%v back into account %v.\n\n",
		data.Amount,
		data.SourceAccount,
	)

	referenceID := fmt.Sprintf("%s-refund", data.ReferenceID)
	bank := BankingService{"bank-api.example.com"}
	confirmation, err := bank.Deposit(data.SourceAccount, data.Amount, referenceID)
	return confirmation, err
}

// @@@SNIPEND
