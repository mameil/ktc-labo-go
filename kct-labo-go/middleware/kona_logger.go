package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"kct-labo-go/kct-labo-go/utils"
	"log"
	"net/http"
	"time"
)

func KonaLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		//todo check kona common header values
		requester := c.GetHeader("X-Kona-Request-Id")
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
		//todo check kona common logging message format
		log.Printf("[%s-REQ] %s %s %s",
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
		log.Printf("[%s-RES] %s %s %d %dms %s",
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
