package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	Username    int64
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}
