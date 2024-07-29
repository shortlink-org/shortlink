package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/usecases/cart"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	workflowID := "CART-" + fmt.Sprintf("%d", time.Now().Unix())

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "CART_TASK_QUEUE",
	}

	state := v1.NewCartState(uuid.New())
	_, err = c.ExecuteWorkflow(context.Background(), options, cart.Workflow, state)
	if err != nil {
		log.Fatalln("unable to execute workflow", err)
	}

	update := v1.AddToCartSignal{Route: v1.ADD_TO_CART, Item: v1.NewCartItem(uuid.New(), 1)}
	err = c.SignalWorkflow(context.Background(), workflowID, "", "ADD_TO_CART_CHANNEL", update)

	resp, err := c.QueryWorkflow(context.Background(), workflowID, "", "getCart")
	if err != nil {
		log.Fatalln("Unable to query workflow", err)
	}
	var result interface{}
	if err := resp.Get(&result); err != nil {
		log.Fatalln("Unable to decode query result", err)
	}

	// Prints a message similar to:
	// 2021/03/31 15:43:54 Received query result Result map[Email: Items:[map[ProductId:0 Quantity:1]]]
	log.Println("Received query result", "Result", result)
}
