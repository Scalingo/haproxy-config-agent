package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/haproxy-config-agent/config"
	"github.com/Scalingo/haproxy-config-agent/env"
	"github.com/Scalingo/haproxy-config-agent/web"
)

func main() {
	err := env.InitConfig()
	if err != nil {
		panic(err)
	}

	log := logger.Default()
	log.WithFields(logrus.Fields{
		"pid":  env.Config.PID,
		"path": env.Config.Path,
	}).Info("Config loaded")

	ctx := logger.ToCtx(context.Background(), log)
	configManager := web.ConfigManager{
		Reloader: config.NewConfigReloader(env.Config.PID, syscall.SIGHUP),
		Writer:   config.NewConfigWriter(env.Config.Path),
	}

	router := web.NewRouter(ctx, configManager)
	log.Infof("Listening on :%s", env.Config.Port)
	go http.ListenAndServe(":"+env.Config.Port, router)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Info("Shutting down!")
}
