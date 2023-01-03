//go:generate wire
//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package csi_di

import (
	"context"

	"github.com/google/wire"

	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
)

// Service - heplers
type Service struct {
	Ctx context.Context
	Log logger.Logger
}

// Context =============================================================================================================
func NewContext() (context.Context, func(), error) {
	ctx := context.Background()

	cb := func() {
		ctx.Done()
	}

	return ctx, cb, nil
}

// CSI =================================================================================================================
var CSISet = wire.NewSet(di.DefaultSet, NewSCIDriver)

func NewSCIDriver(log logger.Logger, ctx context.Context) (*Service, error) {
	return &Service{
		Ctx: ctx,
		Log: log,
	}, nil
}

func InitializeSCIDriver() (*Service, func(), error) {
	panic(wire.Build(CSISet))
}
