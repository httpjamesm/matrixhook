package webserver

import (
	"fmt"
	"matrixhook/env"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartWebserver() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	server := &http.Server{
		Handler: r,
		Addr:    env.WebserverHost + ":" + env.WebserverPort,
	}

	fmt.Println("[Webserver] Starting server on port " + env.WebserverPort)
	server.ListenAndServe()
}
