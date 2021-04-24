/*
Metadata Service. Application layer
*/
package application

import (
	link_store "github.com/batazor/shortlink/internal/services/link/infrastructure/store"
)

type Service struct {
	Store *link_store.LinkStore
}
