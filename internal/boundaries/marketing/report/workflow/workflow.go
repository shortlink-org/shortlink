package workflow

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"github.com/shortlink-org/shortlink/internal/boundaries/marketing/report/activity"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, activity.ComposeGreeting, name).Get(ctx, &result)

	return result, err
}
