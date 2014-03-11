package JsonSoapProxyLib

import (
	_ "encoding/json"
	_ "errors"
	"fmt"
	_ "github.com/nu7hatch/gouuid"
	"io/ioutil"
	_ "log"
	"net/http"
	"net/url"
	_ "unicode/utf8"
)

var soapRequester SoapRequester

func InitSoapRequests(requestList map[string]RequestLookupItem) {
	soapRequester = NewSoapRequest(requestList)
	fmt.Println("httpcontroller init is called")
	fmt.Println(soapRequester)
}

func Handler(rw http.ResponseWriter, r *http.Request) {
	if soapRequester == nil {
		fmt.Println("soapRequester is nil")
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	urlpath, err := url.QueryUnescape(string(r.URL.Path))
	if err != nil {
		fmt.Println(err)
		return
	}

	bodyUnscaped, err := url.QueryUnescape(string(body))
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("bodyUnscaped: ", bodyUnscaped)

	var responseStr string

	responseStr, err = soapRequester.GetResponse(urlpath, bodyUnscaped)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Fprint(rw, responseStr)
}
