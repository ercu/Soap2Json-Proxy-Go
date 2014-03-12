Soap2Json-Proxy-Go
============

This app is a standalone proxy for Golang to request Soap WebServices using Json interface. 

You make a json request and get response as json. The server transforms json to Soap Message and make the request to the webservice. 

You need to create Request and Envelope structs from soap message xml, as seen in example tests/ExampleRequestResponse.go.

Example Web Server


    func main() {
  	    soapRequests := map[string]JsonSoapProxyLib.RequestLookupItem{
		"/list": JsonSoapProxyLib.RequestLookupItem{&ListManagementRequest{}, &ListManagementEnvelope{}, "http://URL_FOR_SERVER/ListService.asmx"},
		}

    	JsonSoapProxyLib.InitSoapRequests(soapRequests)
    	http.HandleFunc("/", JsonSoapProxyLib.Handler)
	    http.ListenAndServe(":8081", nil)
	 }


