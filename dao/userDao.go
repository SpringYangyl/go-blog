package dao

import (
	"log"
)

func GetUserByName(name string) (User,error){
	user :=User{}
	if err := Db.Where("username = ?",name).First(&user).Error;err != nil{
		return user,err
	}
	return user,nil
}
func Register(user *User) bool {
	if err := Db.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
type User struct {
	Id int64
	Username string
	Password string
	Role string
}
func (user User) TableName() string {
	return "user"
}