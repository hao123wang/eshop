package svc

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 微服务所需的所有依赖都在这里创建
type ServiceContext struct {
	DbConn *gorm.DB
}

func NewServiceContext() *ServiceContext {
	svc := &ServiceContext{}

	// 开发环境，从 .env 文件中获取配置信息
	if err := godotenv.Load(); err != nil {
		panic("未找到.env文件")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PWD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	// mysql连接
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}

	svc.DbConn = db
	return svc
}
