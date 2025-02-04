package server

import (
	"net/http"

	"github.com/nikashlabs/hishab/internal/api"
	"github.com/nikashlabs/hishab/internal/server/config"
	"go.uber.org/zap"
)

func NewServer(
	logger *zap.Logger,
	config *config.Config,
) http.Handler {
	mux := http.NewServeMux()
	api.AddRoutes(
		mux,
		logger,
		config,
	)

	var handler http.Handler = mux
	// handler = someMiddleware(handler)
	// handler = someMiddleware2(handler)
	// handler = someMiddleware3(handler)
	return handler
}
