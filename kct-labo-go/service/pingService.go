package service

import (
	"log"
)

func DoPing(userId string, mpaId string) {
	log.Println("PING!!!!")
}

func MakeError() {
	panic("something went wrong.....")
}

func DoPong(userId string, mpaId string) {
	log.Println("PONG!!!!")
}
