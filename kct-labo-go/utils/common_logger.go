package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// T, U, M 값을 저장하는 구조체 (Request 단위)
type LogContext struct {
	T string
	U string
	M string
}

// 전역 변수 (각 요청마다 T, U, M 값을 저장)
var logContextMap = struct {
	sync.RWMutex
	contexts map[*gin.Context]LogContext
}{contexts: make(map[*gin.Context]LogContext)}

// 요청 단위로 T, U, M 값을 저장하는 함수
func SetLogContext(c *gin.Context) {
	logContextMap.Lock()
	defer logContextMap.Unlock()
	logContextMap.contexts[c] = LogContext{
		T: c.GetHeader("X-KM-Correlation-ID"),
		U: c.GetHeader("X-KM-UserId"),
		M: c.GetHeader("X-KM-User-MpaId"),
	}
}

// 요청이 끝나면 Context 삭제
func ClearLogContext(c *gin.Context) {
	logContextMap.Lock()
	defer logContextMap.Unlock()
	delete(logContextMap.contexts, c)
}

// 현재 Context 기반으로 로그 메시지 포맷팅
func formatLog(msg string) string {
	// 현재 시간
	timestamp := time.Now().Format("2006/01/02 15:04:05")

	// Context에서 T, U, M 값 가져오기
	logContextMap.RLock()
	defer logContextMap.RUnlock()

	// 기본값 설정 (없는 경우 빈 배열)
	tValue, uValue, mValue := "T[]", "U[]", "M[]"
	for _, ctx := range logContextMap.contexts {
		if ctx.T != "" {
			tValue = fmt.Sprintf("T[%s]", ctx.T)
		}
		if ctx.U != "" {
			uValue = fmt.Sprintf("U[%s]", ctx.U)
		}
		if ctx.M != "" {
			mValue = fmt.Sprintf("M[%s]", ctx.M)
		}
		break // 한 개만 가져오면 됨 (현재 요청 기준)
	}

	// 최종 로그 메시지
	return fmt.Sprintf("%s %s %s %s %s", timestamp, tValue, uValue, mValue, msg)
}

// 기본 log.Println(), log.Printf()도 자동 적용하도록 변경
type customLogWriter struct{}

func (w customLogWriter) Write(p []byte) (n int, err error) {
	prefixedMsg := formatLog(string(p))
	return os.Stdout.Write([]byte(prefixedMsg))
}

// Logger 초기화 (기본 log.Println()에도 적용)
func InitLogger() {
	log.SetFlags(0)                  // 기본 로그 포맷 제거
	log.SetOutput(customLogWriter{}) // 모든 로그에 자동 Prefix 추가
}

// 에러 로그를 찍고 Stack Trace를 남기는 함수
func LogError(err error) {
	if err != nil {
		log.Printf("ERROR: %s\n%s", err.Error(), debug.Stack())
	}
}
