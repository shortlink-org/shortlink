package http_chi

import (
	"context"
	"github.com/batazor/shortlink/pkg/internal/store"
)

type API struct {
	store store.DB
	ctx   context.Context
}

type addRequest struct {
	Url      string
	Describe string
}

type getRequest struct {
	Hash     string
	Describe string
}

type deleteRequest struct {
	Hash string
}
