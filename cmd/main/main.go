package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"abfw-proxy/config"
	"abfw-proxy/internal/api"
	"abfw-proxy/internal/env"

)

// @title ABFW-PROXY API
// @version 0.0.1
// @query.collection.format multi.
func main() {
	logger, syn := log.New(&log.Opts{
		Development: true,
		LogLevel:    "INFO",
		LogConfig:   nil,
	})
	defer syn()

	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	e := env.Env{
		Log:  logger,
		Conf: cfg,
	}

	handler := api.NewAPI(&e)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Addr, cfg.Server.Port),
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("can't run server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("graceful shutdown start...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("graceful shutdown successful")
}
