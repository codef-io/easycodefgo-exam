package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

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
		"organization":    "0001",
		"userName":        "이름",
		"identity":        "199101011",
		"phoneNo":         "01000000000",
		"telecom":         "0",
		"timeout":         "120",
		"authMethod":      "0",
		"applicationType": "0",
		"phoneNo1":        "",
	}

	// 코드에프 정보 조회 요청
	// - 서비스타입(0:정식, 1:데모, 2:샌드박스)
	productURL := "/v1/kr/public/ft/do-not-call/set-register" // 공정거래위원회 수신거부 등록/해제 신청 URL
	serviceType := ecg.TypeSandbox

	// * 정보 조회 요청 메소드 사용
	result, err := codef.RequestProduct(productURL, serviceType, parameter)
	if err != nil {
		log.Fatalln(err)
	}

	// 결과 확인
	fmt.Println(result)

	// 추가인증 정보 설정
	fmt.Println("보안숫자 입력 : ")
	var secureNo string
	_, err = fmt.Scanf("%s", &secureNo)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("보안숫자 입력 : " + secureNo)

	parameter["secureNo"] = secureNo
	parameter["secureNoRefresh"] = "0"
	setUpTwoWayInfo(parameter, result)

	// 코드에프 추가인증 요청
	result, err = codef.RequestCertification(productURL, serviceType, parameter)
	if err != nil {
		log.Fatalln(err)
	}

	// 결과 확인
	fmt.Println(result)

	// 추가인증 정보 설정
	fmt.Println("SMS인증숫자 입력 : ")
	var smsAuthNo string
	_, err = fmt.Scanf("%s", &smsAuthNo)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("SMS인증숫자 입력 : " + smsAuthNo)

	parameter["smsAuthNo"] = smsAuthNo
	setUpTwoWayInfo(parameter, result)

	// 코드에프 추가인증 요청
	result, err = codef.RequestCertification(productURL, serviceType, parameter)
	if err != nil {
		log.Fatalln(err)
	}

	// 결과 확인
	fmt.Println(result)
}

func setUpTwoWayInfo(parameter map[string]interface{}, result string) {
	jsonData := map[string]interface{}{}
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(result)))
	decoder.UseNumber()
	if err := decoder.Decode(&jsonData); err != nil {
		log.Fatalln(err)
	}

	data := jsonData["data"].(map[string]interface{})
	jobIndex, err := strconv.ParseUint(string(data["jobIndex"].(json.Number)), 10, 32)
	if err != nil {
		log.Fatalln(err)
	}
	threadIndex, err := strconv.ParseUint(string(data["threadIndex"].(json.Number)), 10, 32)
	if err != nil {
		log.Fatalln(err)
	}
	twoWayTimestamp, err := strconv.ParseUint(string(data["twoWayTimestamp"].(json.Number)), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	twoWayInfo := map[string]interface{}{
		"jobIndex":        jobIndex,
		"threadIndex":     threadIndex,
		"jti":             data["jti"],
		"twoWayTimestamp": twoWayTimestamp,
	}

	parameter["twoWayInfo"] = twoWayInfo
	parameter["is2Way"] = true
}
