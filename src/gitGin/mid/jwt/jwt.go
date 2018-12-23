package jwt

import (
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//https://segmentfault.com/a/1190000013297828
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		code := http.StatusOK
		if token == "" {
			code = http.StatusUnauthorized
		} else {
			_, err := ParseJWTToken(token)
			if err != nil {
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

var jwtSecret = []byte("")

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
