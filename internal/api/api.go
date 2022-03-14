package api

import (
	"time"

	"abfw-proxy/internal/env"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginPrometheus "github.com/zsais/go-gin-prometheus"
)

func NewAPI(e *env.Env) *gin.Engine {
	if !e.Conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	if !e.Conf.Debug {
		r.Use(ginzap.RecoveryWithZap(e.Log, true))
	}

	r.Use(
		ginzap.Ginzap(e.Log, time.RFC3339, true, []string{
			"/metrics",
			"/health",
			"/ready",
		}),
	)

	ginPrometheus.NewPrometheus("").Use(r)

	r.GET("/health", Health)
	r.GET("/ready", Ready)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
