package JsonSoapProxyLib

import (
	"encoding/json"
	_ "errors"
	"fmt"
	_ "unicode/utf8"
)

type RequestLookupItem struct {
	Request IRequest
	Env     IEnvelope
	Url     string
}

type SoapRequest struct {
	rmap map[string]RequestLookupItem
}

type SoapRequester interface {
	GetResponse(string, string) (string, error)
}

func NewSoapRequest(requestList map[string]RequestLookupItem) *SoapRequest {
	w := new(SoapRequest)
	w.rmap = requestList
	return w
}

func (w *SoapRequest) GetResponse(requestType string, requestBody string) (string, error) {

	return w.RequestWebService(requestBody, w.rmap[requestType].Request, w.rmap[requestType].Env, w.rmap[requestType].Url)
}

func (w *SoapRequest) RequestWebService(bodyUnscaped string, request IRequest, env IEnvelope, url string) (string, error) {
	newRequest := request.NewRequest()
	newEnv := env.NewEnv()
	err := json.Unmarshal([]byte(bodyUnscaped), newRequest)
	if err != nil {
		fmt.Println("can not parse json, ", err, bodyUnscaped)
		return "", err
	}

	MakeRequest(url, newRequest.(IRequest).GetQuery(), newEnv)

	response := struct {
		Result  bool
		Message string
		Data    interface{}
	}{
		true,
		"Successful",
		newEnv,
	}

	jsonByte, err := json.Marshal(response)
	jsonStr := string(jsonByte)

	if err != nil {
		return "", err
	}

	return jsonStr, nil

}
