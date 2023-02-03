package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
*/
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("gorm:gorm@tcp(localhost:9911)/gorm?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to open connect database")
	}
	// 迁移schema
	db.AutoMigrate(&Product{})
	// create
	db.Create(&Product{Code: "D42", Price: 100})

}
