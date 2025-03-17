package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"kct-labo-go/kct-labo-go/controller/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
*
go 에서는 테스트 코드를 항상 클래스와 동일한 패키지에서 작성하고
테스트하고자 하는 클래스명 뒤에 _test 을 붙혀서 파일명을 작성한다
그리고 테스트 명은 항상 Test${메소드명} 으로 작성한다
*/
func TestGetPing(t *testing.T) {
	r := gin.New()
	r.GET("/test/ping", GetPing)

	userId := "3052731353"
	mpaId := "16982698336781016760331400303886"

	req, _ := http.NewRequest("GET", "/test/ping?userId=3052731353&mpaId=16982698336781016760331400303886", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("Response code is %v", resp.Code)
	}

	println(resp.Body.String())

	assert.Equal(t, resp.Code, 200)

	//여기서
	//json.Unmarshal([]byte(resp.Body.String()), &responseGo)

	var response dto.PongRequest
	json.Unmarshal([]byte(resp.Body.String()), &response)

	assert.Equal(t, response.UserId, userId)
	assert.Equal(t, response.MpaId, mpaId)
}
