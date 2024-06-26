package global

import (
	"github.com/maivankien/go-ecommerce-api/pkg/logger"
	"github.com/maivankien/go-ecommerce-api/pkg/settings"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
)
