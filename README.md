# easycodefgo-exam

### About
[esaycodefgo](https://github.com/codef-io/easycodefgo) 라이브러리 사용 예제입니다.  

### Quick Start
```bash
$ git clone https://github.com/codef-io/easycodefgo-exam.git

$ go get github.com/codef-io/easycodefgo
```

### Use it

각 예제를 사용할 때 인증 정보와 암호화 키 정보를 셋팅해서 사용합니다.

```go
const (
	demoClientID     = ""
	demoClientSecret = ""

	clientID     = ""
	clientSecret = ""

	publicKey = ""
)
```

예제들은 기본적으로 샌드박스 타입으로 요청되며 데모 또는 정식 버전 테스트를 원하실 경우 요청 타입을 변경해서 사용합니다.

```go
...

result, err := codef.RequestProduct(ecg.TypeSandbox)

...
```
