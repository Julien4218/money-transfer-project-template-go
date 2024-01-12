package main

import (
	"log"
	"money-transfer-project-template-go/app"
	"os"

	"go.temporal.io/sdk/worker"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	licenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")
	app.NrApp, _ = newrelic.NewApplication(
		newrelic.ConfigAppName("temporal-money-transfer"),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	c, err := app.GetTemporalClient()
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	w := worker.New(c, app.GenericWorkflowTaskQueueName, worker.Options{})

	w.RegisterWorkflow(app.GenericWorkflow)
	w.RegisterActivity(app.Withdraw)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
