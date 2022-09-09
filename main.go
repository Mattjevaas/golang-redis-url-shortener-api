package main

import (
	"URLShortener/injector"
	"fmt"
)

func main() {

	r := injector.InitializeServer()
	if err := r.Run(); err == nil {
		msg := fmt.Sprintf("Failed to start server: Error %v", err)
		panic(msg)
	}
}
