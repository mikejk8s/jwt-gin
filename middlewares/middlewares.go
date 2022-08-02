package middlewares

import (
	"github.com/gin-gonic/gin"
	"jwt-gin/utils/token"
	"log"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		log.Println("User not logged in")
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
		return
	}
	c.Next()
}
