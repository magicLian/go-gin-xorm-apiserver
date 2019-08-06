package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//EIToken := c.Request.FormValue("EIToken")
		//
		//if EIToken == "" {
		//	noAuth(c, "EIToken not found")
		//	return
		//}
		//c.Next()
		return
	}
}

func noAuth(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": msg,
	})
	c.Abort()
}
