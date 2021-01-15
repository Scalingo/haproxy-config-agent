package web

import (
	"context"

	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/haproxy-config-agent/env"
)

func NewRouter(ctx context.Context, c ConfigManager) *handlers.Router {
	log := logger.Get(ctx)
	r := handlers.NewRouter(log)
	r.Use(handlers.ErrorMiddleware)
	r.Use(handlers.AuthMiddleware(func(user, password string) bool {
		return user == env.Config.User && password == env.Config.Password

	}))
	r.HandleFunc("/config", c.PostConfig)
	return r
}
