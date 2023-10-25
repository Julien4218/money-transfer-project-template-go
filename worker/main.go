package main

import (
	"log"
	"money-transfer-project-template-go/app"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// @@@SNIPSTART money-transfer-project-template-go-worker
func main() {
	licenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")
	app.NrApp, _ = newrelic.NewApplication(
		newrelic.ConfigAppName("temporal-money-transfer"),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, app.MoneyTransferTaskQueueName, worker.Options{})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(app.MoneyTransfer)
	w.RegisterActivity(app.Withdraw)
	w.RegisterActivity(app.Deposit)
	w.RegisterActivity(app.Refund)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}

// @@@SNIPEND
