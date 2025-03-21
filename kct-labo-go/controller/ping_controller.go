package controller

import (
	"github.com/gin-gonic/gin"
	"kct-labo-go/kct-labo-go/controller/dto"
	"kct-labo-go/kct-labo-go/service"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// go 에서는 기본적으로 파일을 만들 때 snake case를 사용 >> 벌써 이상해
// 타입명 / 함수명 / 필드명에 CamelCase를 사용
func GetPing(c *gin.Context) {

	// 로그 파일 열기 (쓰기, 추가 모드)
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("로그 파일을 열 수 없습니다: %v", err)
	}
	defer logFile.Close()

	// 로그 출력을 파일로 변경
	log.SetOutput(logFile)

	ch := make(chan string)

	// 비동기적으로 값을 채널에 전송
	go func() {
		ch <- "Hello World"

		rand.Seed(time.Now().UnixNano()) // 랜덤 시드 설정

		// 100자리 랜덤 문자열 생성 및 출력
		randomString := generateRandomString(100)
		log.Println(randomString)

		time.Sleep(10 * time.Millisecond) // 10ms 동안 대기

		close(ch)
	}()

	recv := <-ch
	log.Println(recv)

	// 채널에서 값을 읽고 응답 반환
	c.JSON(http.StatusOK, gin.H{"message": recv})
}

func PostPong(c *gin.Context) {
	//객체를 미리 선언해두고
	var req dto.PongRequest

	//아래 메소드를 통해서 req 라는 데이터에 gin.Context 에서 BODY 값을 뽑아다가 넣어준다
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "data": err.Error()})
	}

	service.DoPong(req.UserId, req.MpaId)

	c.JSON(http.StatusOK, gin.H{"status": "200", "data": "success"})
}

func GetPingError(c *gin.Context) {
	service.MakeError()
}
