package JsonSoapProxyLib

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type IEnvelope interface {
	NewEnv() interface{}
}

type IRequest interface {
	GetQuery() string
	NewRequest() interface{}
}

func MakeRequest(url string, query string, v interface{}) bool {
	buf := []byte(query)
	body := bytes.NewBuffer(buf)

	myClient := &http.Client{}
	r, err := myClient.Post(url, "text/xml", body)
	if err != nil {
		fmt.Println(err.Error())
	}
	response, _ := ioutil.ReadAll(r.Body)

	responseStr := strings.Replace(string(response), "soap:", "", -1)

	responseStr = strings.Replace(string(responseStr), "utf-16", "utf-8", -1)

	err = xml.Unmarshal([]byte(responseStr), v)

	if err != nil {
		fmt.Println("can not parse value from xml: ", responseStr, "error:", err.Error())
		return false

	}
	return true
}
