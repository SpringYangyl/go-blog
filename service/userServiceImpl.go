package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-blog/dao"
	"strconv"
	"time"
)

var OneDayOfHours = 60 * 60 * 24
var Secret = "tiktok"
func Register(user *dao.User){
	username := user.Username
	password := user.Password
	role := user.Role
	u,_:=dao.GetUserByName(username)
	if username == u.Username{
		fmt.Println("用户已经存在")
		return
	}
	password = EnCoder(password)
	newUser := dao.User{
		Username: username,
		Password: password,
		Role: role,
	}
	dao.Register(&newUser)
	fmt.Println("注册成功")



}
func EnCoder(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
func NewToken(u dao.User) string {
	expiresTime := time.Now().Unix() + int64(OneDayOfHours)
	//fmt.Printf("expiresTime: %v\n", expiresTime)
	id64 := u.Id
	//fmt.Printf("id: %v\n", strconv.FormatInt(id64, 10))
	claims := jwt.StandardClaims{
		Audience:  u.Username,
		ExpiresAt: expiresTime,
		Id:        strconv.FormatInt(id64, 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    u.Role,
		NotBefore: time.Now().Unix(),
		Subject:   "token",
	}
	var jwtSecret = []byte(Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		//token = "Bearer " + token
		println("generate token success!\n")
		return token
	} else {
		println("generate token fail\n")
		return "fail"
	}
}

func ParseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}


