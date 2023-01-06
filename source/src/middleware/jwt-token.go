package middleware

import (
	"errors"
	"net/http"
	"os"
	"time"

	"server_temp/source/src/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenToken(account string) (string, error) {
	c := model.JWTClaims{
		Account: account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(24) * time.Hour).Unix(),
			Issuer:    "server",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseToken(tokenString string) (*model.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Authorization is null in Header",
			})
			c.Abort()
			return
		}

		mc, err := ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Invalid Token.",
			})
			c.Abort()
			return
		}

		c.Set("account", mc.Account)

		c.Next()
	}
}
