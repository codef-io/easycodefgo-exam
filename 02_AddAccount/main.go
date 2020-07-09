// easycodefgo를 활용한 계정 추가 등록 샘플 코드

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
		"loginType":    "0",
	}

	pwd, err := ecg.EncryptRSA("password", codef.PublicKey)
	if err != nil {
		log.Fatalln(err)
	}
	account["password"] = pwd

	// 인증서 정보 가져오기
	certFile, err := ecg.EncodeToFileString("CERT_FILE_PATH")
	if err != nil {
		log.Fatalln(err)
	}
	keyFile, err := ecg.EncodeToFileString("KEY_FILE_PATH")
	if err != nil {
		log.Fatalln(err)
	}

	// 인증서 정보 셋팅
	account["derFile"] = certFile
	account["keyFile"] = keyFile

	accountList = append(accountList, account)
	parameter := map[string]interface{}{
		"accountList": accountList,
		"connectedId": "8PQI4dQ......hKLhTnZ",
	}

	// 요청
	result, err := codef.AddAccount(ecg.TypeSandbox, parameter)
	if err != nil {
		log.Fatalln(result)
	}

	// 결과 확인
	fmt.Println(result)
}
