package service

import (
	"pep/global"
	"time"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateJsonBlackListRecord
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func CreateJsonBlackListRecord(jwt string) (err error) {
	timer := time.Duration(global.GVA_CONFIG.JWT.ExpiresTime) * time.Second
	return global.GVA_REDIS.Set(jwt, "exist", timer).Err()
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func IsBlacklist(jwt string) bool {
	exist, _ := global.GVA_REDIS.Get(jwt).Result()
	return exist != ""
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(userName).Result()
	return err, redisJWT
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: userName string
//@return: err error, redisJWT string

func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GVA_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GVA_REDIS.Set(userName, jwt, timer).Err()
	return err
}
