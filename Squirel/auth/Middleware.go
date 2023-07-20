package auth

import (
	"fmt"
	"golang/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func MiddlewareAdmin(c *gin.Context) {
	TokenString := c.Request.Header.Get("Authorization")
	if TokenString == "" {
		c.JSON(401, gin.H{"message": "Silahkan Login"})
		c.AbortWithStatus(401)
		return
	}

	token, _ := jwt.Parse(TokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return JWT_KEY, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(401, gin.H{"message": "Silahkan Login Kembali"})
			c.AbortWithStatus(401)
			return
		}
		var account Account
		model.DB.First(&account, claims["Id"])
		if account.Id == 0 {
			c.AbortWithStatus(401)
			return
		}
		c.Set("account", account)
		c.Next()
	} else {
		c.AbortWithStatus(401)
		return
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if origin := c.Request.Header.Get("Origin"); origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Menu-Slug, X-Origin-Path, X-Request-Id,VerificationCode,XMLHttpRequest")
			c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}
		}
		c.Next()
	}
}
