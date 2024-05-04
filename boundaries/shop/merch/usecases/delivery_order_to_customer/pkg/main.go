package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	ctx := context.Background()

	viper.SetDefault("HTTP_SERVER_PORT", "3500")
	viper.SetDefault("HTTP_SERVER_TIMEOUT", "30s")

	daprHost := os.Getenv("DAPR_HOST")
	if daprHost == "" {
		daprHost = "http://localhost"
	}

	daprHttpPort := os.Getenv("DAPR_HTTP_PORT")
	if daprHttpPort == "" {
		daprHttpPort = viper.GetString("HTTP_SERVER_PORT")
	}

	client := &http.Client{
		Timeout: viper.GetDuration("HTTP_SERVER_TIMEOUT"),
	}

	for i := 1; i <= 20; i++ { //nolint:revive // ignore
		order := `{"orderId":` + strconv.Itoa(i) + "}"
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, daprHost+":"+daprHttpPort+"/orders", strings.NewReader(order))
		if err != nil {
			log.Fatal(err.Error())
		}

		// Adding app id as part of the header
		req.Header.Add("Dapr-App-Id", "order-processor")

		// Invoking a service
		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err.Error())
		}

		// Read the response
		result, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Order passed:", string(result))
	}
}
