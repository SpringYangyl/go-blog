package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var Db *gorm.DB
func Init(){
	var err error
	dsn := "root:root@tcp(123.56.91.14:3306)/goBlog?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println("err:",err.Error())
	}

}


