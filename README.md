<h1>KTC 프로젝트를 위한 성능 테스트용 Go 프로젝트</h1>

해당 프로젝트는 gin 프레임워크를 채택 <br>
진행하는 GO 버전은 1.24.1(25.03 기준 최신) <br>
사용하는 GIN 버전은 1.10.0(25.03 기준 최신)

### TODO 공통으로 구현해야하는 항목들
- [ ] Request / Response 로깅 > 완료
- [ ] 에러 로깅 및 서버가 죽지 않도록 적용 > 완료
- [ ] Graceful Shutdown >

<br>

### 프로젝트 생성

```aiignore
--go mod 파일을 생성한다 > go mod 란 build grade 처럼 프로젝트에서 사용하는 의존성을 관리하는 곳
go mod init kct-labo-go

--gin 프레임워크를 사용하기 위해 필요한 패키지를 다운로드 받는다
go get github.com/gin-gonic/gin
```
<br>

### 프로젝트 구성
http 요청 > router > middleware > controller > service <br>
- controller
  - url 요청에 따라서 필요한 서비스를 호출하는 곳
- middleware
  - 라우팅을 통과하기 전처리나 요청에 대한 응답 후처리를 수행하는 계층
  - 현재 구현 사항
    - 로깅처리
- router
  - 클라이언트의 HTTP 요청을 특정 url 경로에 연결된 핸들러 함수로 매핑해주는곳
  - api 를 정리하고 미들웨어를 달아주는 곳
- main
  - 프로젝트를 구동시키기 위한 메인 클래스가 있는 곳
- service
  - 프로젝트에서 실제 비즈니스 로직이 수행되는 곳
- utils
  - 프로젝트에서 필요한 유틸성 클래스들을 포함한 곳