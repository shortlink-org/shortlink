/*
Report application

Make reports for users
*/
package main

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"

	metadata_di "github.com/shortlink-org/shortlink/internal/boundaries/link/metadata/di"
	"github.com/shortlink-org/shortlink/internal/boundaries/marketing/report/shared"
	"github.com/shortlink-org/shortlink/internal/boundaries/marketing/report/workflow"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-report")

	// Init a new service
	service, cleanup, err := metadata_di.InitializeMetaDataService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			service.Log.Error(r.(string)) //nolint:forcetypeassert // simple type assertion
		}
	}()

	// Stop the service gracefully.
	defer cleanup()

	// create a namespace
	c1, err := client.NewNamespaceClient(client.Options{})
	if err != nil {
		service.Log.Error("unable to create Temporal client", field.Fields{
			"error": err.Error(),
		})
		panic(err)
	}

	t := time.Now()
	duration := time.Since(t) + time.Hour*24*1
	err = c1.Register(context.Background(), &workflowservice.RegisterNamespaceRequest{
		Namespace:                        client.DefaultNamespace,
		WorkflowExecutionRetentionPeriod: &duration,
	})
	if err != nil {
		service.Log.Warn("unable to create Temporal namespace", field.Fields{
			"error": err.Error(),
		})
	}

	// get struct logger
	structLogger, err := logger.NewStructLogger(service.Log)
	if err != nil {
		service.Log.Error("unable to create StructLogger", field.Fields{
			"error": err.Error(),
		})
		panic(err)
	}

	// Create the client object just once per process
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
		Logger:   structLogger,
	})
	if err != nil {
		service.Log.Error("unable to create Temporal client", field.Fields{
			"error": err.Error(),
		})
		panic(err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: shared.GreetingTaskQueue,
	}

	// Start the Workflow
	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, workflow.GreetingWorkflow, name)
	if err != nil {
		service.Log.Error("unable to start Workflow", field.Fields{
			"error": err.Error(),
		})
		panic(err)
	}

	// Get the results
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		service.Log.Error("unable to get Workflow result", field.Fields{
			"error": err.Error(),
		})
		panic(err)
	}

	printResults(greeting, we.GetID(), we.GetRunID(), service.Log)
}

func printResults(greeting string, workflowID, runID string, log logger.Logger) {
	log.Info("Workflow completed", field.Fields{
		"WorkflowID": workflowID,
		"RunID":      runID,
		"Greeting":   greeting,
	})
}
