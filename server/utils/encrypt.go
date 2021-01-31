package utils

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

// SetPassword 根据明文密码和加盐值生成密码
func SetPassword(password string, salt string) (verify string, err error) {
	var rb []byte
	rb, err = scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		return
	}
	verify = hex.EncodeToString(rb)
	return
}