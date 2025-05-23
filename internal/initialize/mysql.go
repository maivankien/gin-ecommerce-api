package initialize

import (
	"fmt"
	"time"

	"github.com/maivankien/go-ecommerce-api/global"
	"github.com/maivankien/go-ecommerce-api/internal/model"
	"github.com/maivankien/go-ecommerce-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	config := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	var s = fmt.Sprintf(dsn, config.Username, config.Password, config.Host, config.Port, config.DBName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrorPanic(err, "InitMysql initialization failed")

	global.Mysql = db

	SetPool()
	genTableDAO()
	MigrateTable()
}

func SetPool() {
	config := global.Config.Mysql
	mysqlDB, err := global.Mysql.DB()

	if err != nil {
		fmt.Println("SetPool failed:", err)
	}
	mysqlDB.SetConnMaxIdleTime(time.Duration(config.MaxIdleConns))
	mysqlDB.SetMaxOpenConns(config.MaxOpenConns)
	mysqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime))
}

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(global.Mysql)

	g.GenerateModel("go_crm_user")

	g.Execute()
}

func MigrateTable() {
	err := global.Mysql.AutoMigrate(
		&po.User{},
		&po.Role{},
		&model.GoCrmUserV2{},
	)

	if err != nil {
		global.Logger.Error("MigrateTable failed", zap.Error(err))
	}
}
