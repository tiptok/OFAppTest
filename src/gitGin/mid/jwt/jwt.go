package jwt

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//https://segmentfault.com/a/1190000013297828
//https://juejin.im/post/5b4dd73be51d4518ef2cd571
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		code := http.StatusOK
		if token == "" {
			code = http.StatusUnauthorized
		} else {
			claims, err := ParseJWTToken(token)
			if err != nil {
				code = http.StatusUnauthorized
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = http.StatusUnauthorized
			}
		}

		if code != http.StatusOK {
			c.JSON(code, gin.H{
				"code": code,
				"msg":  "not authorized",
				"data": "",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

var jwtSecret = []byte("123456")

type UserTokenClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//解析 UserTokenClaims
func ParseJWTToken(token string) (*UserTokenClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claim, ok := tokenClaims.Claims.(*UserTokenClaims); ok && tokenClaims.Valid {
			log.Printf("ParseJWTToken:%s -> %v", token, claim)
			return claim, nil
		}
	}

	return nil, err
}

func GenerateToken(username, password string) (string, error) {
	now := time.Now()
	expireTime := now.Add(3 * time.Hour)

	claims := UserTokenClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "jwt",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
