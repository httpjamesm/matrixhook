package main

import (
	"matrixhook/env"
	"matrixhook/matrix"
	"sync"

	"matrixhook/webserver"
)

func main() {
	env.LoadEnv()

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		matrix.ConnectBot()
		wg.Done()
	}()

	go func() {
		webserver.StartWebserver()
		wg.Done()
	}()

	wg.Wait()
}
