// easycodefgo를 활용한 계정 등록 샘플 코드

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
	// - 계정관리 파라미터를 설정(https://developer.codef.io/cert/account/cid-overview)
	accountList := []map[string]interface{}{}
	account := map[string]interface{}{
		"countryCode":  "KR",
		"businessType": "BK",
		"clientType":   "P",
		"organization": "0004",
		"loginType":    "1",
		"id":           "user_id",
	}

	pwd, err := ecg.EncryptRSA("password", codef.PublicKey)
	if err != nil {
		log.Fatalln(err)
	}
	account["password"] = pwd

	accountList = append(accountList, account)
	parameter := map[string]interface{}{
		"accountList": accountList,
	}

	// 요청
	result, err := codef.CreateAccount(ecg.TypeSandbox, parameter)
	if err != nil {
		log.Fatalln(err)
	}

	// 결과 확인
	fmt.Println(result)
}
