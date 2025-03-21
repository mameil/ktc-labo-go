package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"kct-labo-go/kct-labo-go/utils"
	"log"
	"net/http"
	"os"
	"time"
)

func KonaLoggingMiddleware() gin.HandlerFunc {

	// 로그 파일 열기
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("로그 파일을 열 수 없습니다: %v", err)
	}

	// Logger 설정 (파일 출력)
	logger := log.New(logFile, "", log.LstdFlags)

	return func(c *gin.Context) {
		startTime := time.Now()

		requester := c.GetHeader("X-KM-CALLER")
		if requester == "" {
			requester = "CLIENT"
		}

		//tid, userId, mpaId 를 로그에 출력하기 위해 설정
		utils.SetLogContext(c)

		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // body 복구
			requestBody = string(bodyBytes)
		}

		//Request 을 로깅한다
		logger.Printf("[%s-REQ] %s %s %s",
			requester,
			c.Request.Method,
			c.Request.URL.String(),
			requestBody,
		)

		//응답 기록용 ResponseWriter 래핑
		loggingResponseWriter := &responseWriter{ResponseWriter: c.Writer, statusCode: http.StatusOK, body: &bytes.Buffer{}}

		//Request 을 처리한다
		c.Next()

		//Response 을 로깅한다
		duration := time.Since(startTime) //소요시간 계산
		logger.Printf("[%s-RES] %s %s %d %dms %s",
			requester,
			c.Request.Method,
			c.Request.URL.String(),
			loggingResponseWriter.statusCode,
			duration.Milliseconds(),
			loggingResponseWriter.body.String(),
		)

		utils.ClearLogContext(c)

	}

}

// responseWriter 의 응답을 기록하기 위한 Custom ResponseWriter
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(body []byte) (int, error) {
	rw.body.Write(body)
	return rw.ResponseWriter.Write(body)
}
