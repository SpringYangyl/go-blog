package service

import "go-blog/dao"

type UserService interface {
	Register(user *dao.User) bool
}