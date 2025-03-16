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
