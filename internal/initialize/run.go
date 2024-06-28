package initialize

import (
	"github.com/maivankien/go-ecommerce-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitLogger()
	global.Logger.Info("Logger init success", zap.String("success", "true"))
	InitMysql()
	InitRedis()

	router := InitRouter()

	router.Run(":8080")
}
