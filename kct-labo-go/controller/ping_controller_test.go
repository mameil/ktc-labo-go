package controller

import (
	"github.com/gin-gonic/gin"
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

	req, _ := http.NewRequest("GET", "/test/ping", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("Response code is %v", resp.Code)
	}

	println(resp.Body.String())
}
