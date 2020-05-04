// easycodefgo를 활용한 Account 목록 조회
package main

import (
	"fmt"
	"log"

	ecg "gitlab.codef.io/codef-io-dev/easycodefgo.git"
)

const (
	demoClientID     = ""
	demoClientSecret = ""

	clientID     = ""
	clientSecret = ""

	publicKey = ""
)

func main() {
	// 코드에프 인스턴스 생성
	codef := &ecg.Codef{}

	// 데모 클라이언트 정보 설정
	// - 데모 서비스 가입 후 코드에프 홈페이지에 확인 가능(https://codef.io/#/account/keys)
	// - 데모 서비스로 상품 조회 요청시 필수 입력 항목
	codef.SetClientInfoForDemo(demoClientID, demoClientSecret)

	// 정식 클라이언트 정보 설정
	// - 정식 서비스 가입 후 코드에프 홈페이지에 확인 가능(https://codef.io/#/account/keys)
	// - 정식 서비스로 상품 조회 요청시 필수 입력 항목
	codef.SetClientInfo(clientID, clientSecret)

	// 요청 파라미터 설정
	parameter := map[string]interface{}{
		"connectedId": "8PQI4dQ......hKLhTnZ",
	}

	// 요청
	result, err := codef.DeleteAccount(ecg.TypeSandbox, parameter)
	if err != nil {
		log.Fatalln(err)
	}

	// 결과 확인
	fmt.Println(result)
}
