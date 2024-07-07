package global

import (
	"github.com/maivankien/go-ecommerce-api/pkg/logger"
	"github.com/maivankien/go-ecommerce-api/pkg/settings"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mysql  *gorm.DB
)
