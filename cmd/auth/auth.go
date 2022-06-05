/*
Auth service
*/
package main

import (
	"fmt"
	"net/url"

	kratos "github.com/ory/kratos-client-go"
)

func main() {
	adminURL, err := url.Parse("http://localhost:4434")
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = kratos.NewAPIClient(
		&kratos.Configuration{
			Host:             adminURL.Host,
			Scheme:           adminURL.Scheme,
			DefaultHeader:    nil,
			UserAgent:        "",
			Debug:            false,
			Servers:          nil,
			OperationServers: nil,
			HTTPClient:       nil,
		},
	)

	vers := kratos.NewVersion()
	fmt.Printf(`%s\n`, vers)

	health := kratos.NewHealthStatus()
	fmt.Printf(`%s`, health.Status)
}
