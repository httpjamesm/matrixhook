package main

import (
	"matrixhook/env"
	"matrixhook/matrix"
)

func main() {
	env.LoadEnv()

	matrix.ConnectBot()
}
