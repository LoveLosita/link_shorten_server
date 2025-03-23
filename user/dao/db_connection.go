package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB() error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/link_shorten?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s"
	//第一部分：连接数据库，并检测其连接正常性
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //链接数据库
	if err != nil {
		return err
	}
	sqlDB, err := Db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping() //检测数据库连接是否正常
	if err != nil {
		return err
	}
	return nil
}
