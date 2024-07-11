package global

import (
	"github.com/maivankien/go-ecommerce-api/pkg/logger"
	"github.com/maivankien/go-ecommerce-api/pkg/settings"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mysql  *gorm.DB
	Rdb    *redis.Client
)
