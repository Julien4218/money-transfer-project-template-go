package app

import "github.com/newrelic/go-agent/v3/newrelic"

// @@@SNIPSTART money-transfer-project-template-go-shared-task-queue
const MoneyTransferTaskQueueName = "TRANSFER_MONEY_TASK_QUEUE"

// @@@SNIPEND

// @@@SNIPSTART money-transfer-project-template-go-transferdetails
type PaymentDetails struct {
	SourceAccount string
	TargetAccount string
	Amount        int
	ReferenceID   string
}

// @@@SNIPEND

var NrApp *newrelic.Application
