package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/heyyyyyyi/go-biz-admin/utils"
	"net/http"
)

func IsAuthenticate(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")
	if _, err := utils.ParseJwt(cookie); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated"},
		)
	}
	return
}
