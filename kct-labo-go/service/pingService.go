package service

import (
	"log"
)

func DoPing() {
	ch := make(chan string)
	go func() {
		ch <- "PING!!!!"
		close(ch)
	}()
	log.Printf("%v", <-ch)
}

func MakeError() {
	panic("something went wrong.....")
}

func DoPong(userId string, mpaId string) {
	log.Println("PONG!!!!")
}
