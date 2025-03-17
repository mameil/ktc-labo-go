package controller

import (
	"github.com/gin-gonic/gin"
	"kct-labo-go/kct-labo-go/controller/dto"
	"kct-labo-go/kct-labo-go/service"
	"kct-labo-go/kct-labo-go/utils"
	"log"
	"net/http"
)

// go 에서는 기본적으로 파일을 만들 때 snake case를 사용 >> 벌써 이상해
// 타입명 / 함수명 / 필드명에 CamelCase를 사용
func GetPing(c *gin.Context) {
	//query param 을 가져오는 방법
	userId := c.Query("userId")
	mpaId := c.Query("mpaId")

	//쿼리파람으로 들어온 값을 검증한다
	if utils.IsNilOrEmpty(&userId) {
		c.JSON(400, gin.H{"status": "400", "data": "userId is nil or empty"})
	}
	if utils.IsNilOrEmpty(&mpaId) {
		c.JSON(400, gin.H{"status": "400", "data": "mpaId is nil or empty"})
	}

	log.Printf("userId : %s, mpaId : %s", userId, mpaId)

	//서비스를 호출한다
	service.DoPing(userId, mpaId)

	//JSON 메소드는 gin 에서 제공하는 메소드로, http 응답코드 + 응답할 내용을리터한다
	c.JSON(http.StatusOK, gin.H{"status": "200", "data": "success"})
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
