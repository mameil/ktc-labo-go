package main

import (
	"kct-labo-go/kct-labo-go/router"
	"kct-labo-go/kct-labo-go/utils"
)

func main() {
	//logging 설정
	utils.InitLogger()

	r := router.SetupRouter()

	r.Run(":4545")
}
