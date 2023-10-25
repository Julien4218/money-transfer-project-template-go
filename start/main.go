package main

import (
	"context"
	"log"
	"os"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"

	"money-transfer-project-template-go/app"
)

// @@@SNIPSTART money-transfer-project-template-go-start-workflow
func main() {
	sourceAccount := "85-150"
	targetAccount := "43-812"
	if len(os.Args) > 1 {
		sourceAccount = os.Args[1]
		if len(os.Args) > 2 {
			targetAccount = os.Args[2]
		}
	}
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}

	defer c.Close()

	input := app.PaymentDetails{
		SourceAccount: sourceAccount,
		TargetAccount: targetAccount,
		Amount:        250,
		ReferenceID:   "12345",
	}

	id, _ := uuid.NewRandom()
	options := client.StartWorkflowOptions{
		ID:        "pay-invoice-" + id.String(),
		TaskQueue: app.MoneyTransferTaskQueueName,
	}

	log.Printf("Starting transfer from account %s to account %s for %d", input.SourceAccount, input.TargetAccount, input.Amount)

	we, err := c.ExecuteWorkflow(context.Background(), options, app.MoneyTransfer, input)
	if err != nil {
		log.Fatalln("Unable to start the Workflow:", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	// var result string

	// err = we.Get(context.Background(), &result)

	// if err != nil {
	// 	log.Fatalln("Unable to get Workflow result:", err)
	// }

	// log.Println(result)
}

// @@@SNIPEND
