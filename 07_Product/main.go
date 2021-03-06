package main

import (
	"fmt"
	"log"

	ecg "github.com/codef-io/easycodefgo"
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
	codef := &ecg.Codef{
		PublicKey: publicKey,
	}

	// 데모 클라이언트 정보 설정
	// - 데모 서비스 가입 후 코드에프 홈페이지에 확인 가능(https://codef.io/#/account/keys)
	// - 데모 서비스로 상품 조회 요청시 필수 입력 항목
	codef.SetClientInfoForDemo(demoClientID, demoClientSecret)

	// 정식 클라이언트 정보 설정
	// - 정식 서비스 가입 후 코드에프 홈페이지에 확인 가능(https://codef.io/#/account/keys)
	// - 정식 서비스로 상품 조회 요청시 필수 입력 항목
	codef.SetClientInfo(clientID, clientSecret)

	// 요청 파라미터 설정
	// - 각 상품별 파라미터를 설정(https://developer.codef.io/products)
	parameter := map[string]interface{}{
		"connectedId":  "8PQI4dQ......hKLhTnZ",
		"organization": "0004",
	}

	// 코드에프 정보 조회 요청
	// - 서비스타입(0:정식, 1:데모, 2:샌드박스)
	// 개인 보유계좌 조회 (https://developer.codef.io/products/bank/common/p/account)
	productURL := "/v1/kr/bank/p/account/account-list"
	result, err := codef.RequestProduct(productURL, ecg.TypeSandbox, parameter)
	if err != nil {
		log.Fatalln(err)
	}

	// 결과 출력
	fmt.Println(result)
}
