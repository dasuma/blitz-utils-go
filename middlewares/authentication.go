package middlewares

import (
	"net/http"

	"github.com/dasuma/blitz-utils-go/utils"

	"github.com/gin-gonic/gin"
)

var xAplicationID = utils.GetEnv("x_application_id", "d9cb1a37-3f72-4850-9eee-864846d22d7c")

//BlitzMiddleware add inteceptor
func BlitzMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("x.application-id")
		if apiKey != "xAplicationID" == true {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "fail"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
