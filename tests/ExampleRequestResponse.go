package JsonSoapProxyLib

import (
	"fmt"
)

type ListManagementEnvelope struct {
	Body struct {
		Response struct {
			Result struct {
				TransactionHeader struct {
					RequestDateTime  string `xml:"request_datetime", json:"request_datetime"`
					ResponseDatetime string `xml:"response_datetime", json:"response_datetime"`
				} `xml:"transaction_header"`

				TransactionBody struct {
					ListItems struct {
						ListItem []struct {
							Name      string `xml:"list_item_name", json:"list_item_name"`
							PromtCpin string `xml:"prompt_cpin", json:"prompt_cpin"`
							Value1    string `xml:"list_item_value_1", json:"list_item_value_1"`
							Value2    string `xml:"list_item_value_2", json:"list_item_value_2"`
						} `xml:"list_item", json:"list_item"`
					} `xml:"list_items", json:"list_items"`
				} `xml:"transaction_body", json:"transaction_body"`
			} `xml:"ViewListResult"`
		} `xml:"ViewListResponse"`
	} `xml:"Body"`
}

func (env *ListManagementEnvelope) NewEnv() interface{} {
	return new(ListManagementEnvelope)
}

type ListManagementRequest struct {
	ClientId        string `json:"clientId"`
	DateTime        string `json:"dateTime"`
	ReferenceNo     string `json:"referenceNo"`
	SendSms         string `json:"sendSms"`
	SendSmsLanguage string `json:"sendSmsLanguage"`
	Msisdn          string `json:"msisdn"`
	ListType        string `json:"listType"`
}

func (request *ListManagementRequest) NewRequest() interface{} {
	return new(ListManagementRequest)
}
func (request *ListManagementRequest) GetQuery() string {

	return fmt.Sprintf(`<?xml version="1.0" encoding="utf-16"?><soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
  <soap:Body>
    <ViewList xmlns="http://phaymobile.cardtekgroup/Tmm/">
      <ViewListRequest xmlns="http://phaymobile.cardtekgroup/Tmm/ListMngmnt">
        <transaction_header>
          <client_id xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">%s</client_id>
          <request_datetime xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">%s</request_datetime>
          <request_reference_no xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">%s</request_reference_no>
          <send_sms xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">%s</send_sms>
          <send_sms_language xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">%s</send_sms_language>
        </transaction_header>
        <transaction_body>
          <msisdn>%s</msisdn>
          <list_type>%s</list_type>
        </transaction_body>
      </ViewListRequest>
    </ViewList>
  </soap:Body>
</soap:Envelope>`, request.ClientId, request.DateTime, request.ReferenceNo, request.SendSms, request.SendSmsLanguage, request.Msisdn, request.ListType)

}
func (l *ListManagementEnvelope) GetResult() string {
	return ""
}
