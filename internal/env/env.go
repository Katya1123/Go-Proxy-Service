package env

import (
	"go.uber.org/zap"

	"abfw-proxy/config"
)

type Env struct {
	Log  *zap.Logger
	Conf *config.Config
}
