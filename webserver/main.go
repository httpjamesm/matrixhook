package webserver

import (
	"fmt"
	"matrixhook/env"
	"matrixhook/matrix"
	"net/http"

	"github.com/gin-gonic/gin"
	"maunium.net/go/mautrix/id"
)

type sendMessageRequest struct {
	Content string `json:"content" binding:"required"`
}

func StartWebserver() {
	//Disable GIN debug mode
	if !env.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CheckPresharedKey())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	r.POST("/", func(c *gin.Context) {
		// get body of request as text
		var json sendMessageRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Invalid body",
			})
			return
		}

		// send message to matrix
		matrix.MatrixClient.SendText(id.RoomID(env.RoomId), json.Content)

		c.JSON(200, gin.H{
			"success": true,
			"message": "Message sent",
		})
	})

	server := &http.Server{
		Handler: r,
		Addr:    env.WebserverHost + ":" + env.WebserverPort,
	}

	if env.DebugMode {
		fmt.Println("[Webserver] Starting server on port " + env.WebserverPort)
	}
	server.ListenAndServe()
}
