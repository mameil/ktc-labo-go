package controller

import (
	"github.com/gin-gonic/gin"
	"kct-labo-go/kct-labo-go/service"
)

// ㅋㅋ go 에서는 기본적으로 파일을 만들 때 snake case를 사용
// 타입명 / 함수명 / 필드명에 CamelCase를 사용
func GetPing(c *gin.Context) {
	service.DoPing("3052731353", "16982698336781016760331400303886")
}

func GetPingError(c *gin.Context) {
	service.MakeError()
}
