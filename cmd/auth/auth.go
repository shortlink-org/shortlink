/*
Auth service
*/
package main

import (
	"fmt"
	"net/url"

	kratos "github.com/ory/kratos-client-go/client"
	"github.com/ory/kratos-client-go/client/health"
	"github.com/ory/kratos-client-go/client/version"
)

func main() {
	adminURL, err := url.Parse("http://localhost:4434")
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := kratos.NewHTTPClientWithConfig(
		nil,
		&kratos.TransportConfig{
			Schemes:  []string{adminURL.Scheme},
			Host:     adminURL.Host,
			BasePath: adminURL.Path,
		},
	)

	vers, err := admin.Version.GetVersion(version.NewGetVersionParams())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vers.GetPayload().Version)

	ok, err := admin.Health.IsInstanceAlive(health.NewIsInstanceAliveParams())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ok.Payload.Status)
}
