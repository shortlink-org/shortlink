//go:generate go run github.com/ServiceWeaver/weaver/cmd/weaver generate .

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

type app struct {
	weaver.Implements[weaver.Main]
	reverser weaver.Ref[Reverser]
	hello    weaver.Listener
}

func (app *app) Main(ctx context.Context) error {
	fmt.Printf("hello listener available on %v\n", app.hello)

	// Serve the /hello endpoint.
	http.Handle("/hello", weaver.InstrumentHandlerFunc("hello",
		func(w http.ResponseWriter, r *http.Request) {
			name := r.URL.Query().Get("name")
			if name == "" {
				name = "World"
			}
			reversed, err := app.reverser.Get().Reverse(ctx, name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "Hello, %s!\n", reversed)
		}))

	return http.Serve(app.hello, nil)
}

func main() {
	// Get a network listener on address "localhost:12345".
	err := weaver.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
