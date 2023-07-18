package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goweb02/Database/mysql"
	"goweb02/Database/redis"
	"goweb02/config"
	"time"
)

type JWT struct {
	SecretKey []byte
	Duration  int64 //分钟
}

var InterNalJwt *JWT

type MyClaims struct { //载荷要携带哪些信息？我这里把非敏感信息都携带了
	ID       int
	Username string
	Email    string
	jwt.StandardClaims
	//实现StandardClaims即可
}

func GetJWT() *JWT { //第一次调用生成新的jwt，后面就用这个jwt
	if InterNalJwt == nil {
		fmt.Println("jwt实例为空,新建对象")
		return &JWT{[]byte(config.JWTsecretKey), config.JWTexpireTime}
	}
	fmt.Println("jwt已经被实例化")
	return InterNalJwt
}

func (this *JWT) CreateToken(user mysql.User) string {

	claims := MyClaims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "myWeb",
			NotBefore: time.Now().Unix(),           // 签名生效时间
			IssuedAt:  time.Now().Unix(),           // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: expireAtTime(this.Duration), // 签名过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString(this.SecretKey) //根据载荷前面并生成jwt
	if err != nil {
		fmt.Println("jwt签名失败:", err)
		return ""
	}
	fmt.Println(token)

	return jwt
}

func expireAtTime(minute int64) int64 {
	timenow := time.Now()
	expire := time.Duration(minute) * time.Minute
	return timenow.Add(expire).Unix()
}

func (this *JWT) parseToken(c *gin.Context) { //中间件函数，？？？？1.检验jwt登出黑名单？？？，如果未登出： 2.解析jwt
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "token为空",
		}) //不再执行后面的接口函数
		return
	}

	exists, err := redis.RDB.Exists(redis.CTX, tokenString).Result()
	if err != nil {
		fmt.Println("查询失败")
	}
	if exists == 1 { //存在该键
		fmt.Println("jwt存在黑名单中，视作失效")
		c.AbortWithStatusJSON(200, "token已经登出")
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return this.SecretKey, nil
	})
	if err != nil {
		fmt.Println("token无效")
		c.AbortWithStatusJSON(200, gin.H{
			"message": "token无效",
		})
		return
	}
	fmt.Println(token.Claims) //token.Claims就是解析出的载荷
	c.Set("claims", token.Claims)
	//c.JSON(200, token.Claims)
	c.Next()
}

func (this *JWT) RefreshToken() {

}
