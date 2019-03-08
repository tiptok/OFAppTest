package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:admin@(127.0.0.1:3306)/TopDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Open DB Error:", err.Error())
	} else {
		fmt.Println("db connect:", db)
	}
}
