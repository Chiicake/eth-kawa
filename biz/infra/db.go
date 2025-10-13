package infra

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var GlobalDB *gorm.DB

func InitDB() {
	// init a sqlite database as a global db
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出目标
		logger.Config{
			SlowThreshold: time.Second, // 慢查询阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 彩色输出
		},
	)

	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		Logger: newLogger, // 使用自定义日志配置
	})

	sqlDB, err := db.DB()
	if err != nil {
		panic("获取数据库实例失败: " + err.Error())
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)    // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)   // 最大打开连接数
	sqlDB.SetConnMaxLifetime(10) // 连接的最大生存期（分钟）

	if err != nil {
		panic(err)
	}
	GlobalDB = db
}
