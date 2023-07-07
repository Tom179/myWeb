package jwt

import "github.com/gin-gonic/gin"

func ParseJWTMiddleWare() gin.HandlerFunc {
	return GetJWT().parseToken
}
