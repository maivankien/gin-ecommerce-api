package initialize

import (
	"github.com/maivankien/go-ecommerce-api/global"
	"github.com/maivankien/go-ecommerce-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
