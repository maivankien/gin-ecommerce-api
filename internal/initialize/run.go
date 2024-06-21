package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()

	router := InitRouter()

	router.Run(":8080")
}
