package webserver

import (
	"matrixhook/env"

	"github.com/gin-gonic/gin"
)

func CheckPresharedKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if env preshared key is set to "none"
		if env.PresharedKey == "none" {
			c.Next()
			return
		}

		// check authorization header for preshared key
		if c.Request.Header.Get("authorization") != env.PresharedKey {
			c.AbortWithStatus(418)
			return
		}

		c.Next()
	}
}
