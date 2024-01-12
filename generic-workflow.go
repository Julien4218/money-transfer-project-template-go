package app

import (
	"fmt"
	"money-transfer-project-template-go/app/shared"
	"strconv"
	"strings"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func GenericWorkflow(ctx workflow.Context, definition shared.WorkflowDefinition) (string, error) {
	txn := NrApp.StartTransaction("GenericWorkflow")
	defer txn.End()

	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        3,
		NonRetryableErrorTypes: []string{"InvalidAccountError", "InsufficientFundsError"},
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)

	inputs := make(map[string]interface{})
	for _, input := range definition.Inputs {
		switch strings.ToLower(input.Type) {
		case "int":
			ival, _ := strconv.Atoi(input.Default)
			inputs[input.Name] = ival
		case "bool":
			bval, _ := strconv.ParseBool(input.Default)
			inputs[input.Name] = bval
		default:
			inputs[input.Name] = input.Default
		}
	}

	for _, step := range definition.Steps {
		// queueName := fmt.Sprintf("%s-%s", step.Type, step.Version)
		var result string
		err := workflow.ExecuteActivity(ctx, step.Type, inputs).Get(ctx, &result)
		if err != nil {
			return "", err
		}
		NrApp.RecordLog(newrelic.LogData{
			Message: fmt.Sprintf("%s executed successfully, result:%s", step.Type, result),
		})
	}

	// // Withdraw money.
	// var withdrawOutput string
	// withdrawErr := workflow.ExecuteActivity(ctx, Withdraw, input).Get(ctx, &withdrawOutput)
	// if withdrawErr != nil {
	// 	return "", withdrawErr
	// }

	// // Deposit money.
	// var depositOutput string
	// depositErr := workflow.ExecuteActivity(ctx, Deposit, input).Get(ctx, &depositOutput)
	// if depositErr != nil {
	// 	// The deposit failed; put money back in original account.

	// 	var result string
	// 	refundErr := workflow.ExecuteActivity(ctx, Refund, input).Get(ctx, &result)
	// 	if refundErr != nil {
	// 		return "",
	// 			fmt.Errorf("Deposit: failed to deposit money into %v: %v. Money could not be returned to %v: %w",
	// 				input.TargetAccount, depositErr, input.SourceAccount, refundErr)
	// 	}

	// 	return "", fmt.Errorf("Deposit: failed to deposit money into %v: Money returned to %v: %w",
	// 		input.TargetAccount, input.SourceAccount, depositErr)
	// }

	result := fmt.Sprintf("Workflow complete")
	return result, nil
}
