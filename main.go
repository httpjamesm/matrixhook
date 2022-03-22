package main

import (
	"matrixhook/env"
	"matrixhook/matrix"
	"matrixhook/webserver"
)

func main() {
	//Load in environmental variables from .env
	env.LoadEnv()

	//Start matrix client
	matrix.ConnectBot()
	
	//Start webserver
	webserver.StartWebserver()
}
