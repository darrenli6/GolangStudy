package models

import (
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/define"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB() {
	dsn := define.MysqlDSN + "/platform?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	err = db.AutoMigrate(&DeviceBasic{}, &ProductBasic{}, &UserBasic{})
	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	DB = db
}
