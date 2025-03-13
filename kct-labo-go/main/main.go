package main

import (
	"kct-labo-go/kct-labo-go/router"
)

func main() {
	r := router.SetupRouter()

	r.Run(":4545")
}
