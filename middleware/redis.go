package middleware

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

}
func SetToken(val string) {
	rdb.Set(Ctx, "token", val, 0)

}
func GetToken() string {
	token := rdb.Get(Ctx, "token")
	return token.Val()
}
func SetKey(key string, val any) {
	rdb.Set(Ctx, key, val, 0)
}
func GetKey(key string) (any, error) {
	val := rdb.Get(Ctx, key)
	return val.Result()
}
func Incre(key string) {
	rdb.Incr(Ctx, key)
}
