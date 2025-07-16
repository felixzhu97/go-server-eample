package database

import (
	"go-server-example/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init 初始化数据库连接
func Init() error {
	var err error
	
	// 使用SQLite作为示例数据库
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	
	if err != nil {
		return err
	}
	
	// 自动迁移数据库表
	err = DB.AutoMigrate(
		&models.User{},
		&models.Post{},
	)
	
	if err != nil {
		return err
	}
	
	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
} 