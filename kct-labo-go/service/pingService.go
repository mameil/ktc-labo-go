package service

import "log"

func DoPing(userId string, mpaId string) {
	log.Printf("U[%s] M[%s] PING!!!!\n", userId, mpaId)
}
