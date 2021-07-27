package request

import (
	"github.com/golang-jwt/jwt"
)

// Custom claims structure
type CustomClaims struct {
	Username    int64
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}
