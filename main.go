package main

import (
	"context"
	"net/http"
	"syscall"

	"github.com/Scalingo/go-internal-tools/logger"
	"github.com/Scalingo/haproxy-config-agent/config"
	"github.com/Scalingo/haproxy-config-agent/env"
	"github.com/Scalingo/haproxy-config-agent/web"
	"github.com/sirupsen/logrus"
)

func main() {
	err := env.InitConfig()
	if err != nil {
		panic(err)
	}

	log := logger.Default()
	log.SetLevel(logrus.InfoLevel)
	log.WithFields(logrus.Fields{
		"pid":  env.Config.TargetPID,
		"path": env.Config.Path,
	}).Info("Config loaded")

	ctx := logger.ToCtx(context.Background(), log)
	configManager := web.ConfigManager{
		Reloader: config.NewConfigReloader(env.Config.TargetPID, syscall.SIGHUP),
		Writer:   config.NewConfigWriter(env.Config.Path),
	}

	router := web.NewRouter(ctx, configManager)
	log.Infof("Listening on :%s", env.Config.Port)
	http.ListenAndServe(":"+env.Config.Port, router)
}
