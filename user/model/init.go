package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	conn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = conn
	db, _ := DB.DB()
	//db.LogMode(true)
	//if gin.Mode() == "release" {
	//	db.LogMode(false)
	//}
	//默认不加复数
	//db.SingularTable(true)
	//设置连接池
	//空闲
	db.SetMaxIdleConns(20)
	//打开
	db.SetMaxOpenConns(100)
	//超时
	db.SetConnMaxLifetime(time.Second * 30)
	//DB = db
	migration()
}
