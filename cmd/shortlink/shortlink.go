package main

import (
	"context"
	"github.com/batazor/shortlink/pkg/api/http"
	log "github.com/batazor/shortlink/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// Add context
	ctx := context.Background()
	ctx = log.WithLogger(ctx, *logger)

	// start HTTP-server
	err := http.Run(ctx)
	if err != nil {
		logger.Panic(err.Error())
	}
}
