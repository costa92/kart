package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	app, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()
	app.Run()
}
