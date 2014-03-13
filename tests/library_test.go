package JsonSoapProxyLib

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func Test_CanParseListManagementResponse(t *testing.T) {

	var listManagement ListManagementEnvelope
	responseStr := `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
  <soap:Body>
    <ViewListResponse xmlns="http://phaymobile.cardtekgroup/Tmm/">
      <ViewListResult>
        <transaction_header xmlns="http://phaymobile.cardtekgroup/Tmm/ListMngmnt">
          <request_datetime xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">0001-01-01T00:00:00</request_datetime>
          <send_sms xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">Y</send_sms>
          <send_sms_language xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">aar</send_sms_language>
          <client_id xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">10000001</client_id>
          <response_datetime xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">2014-02-24T11:01:55.3739948+02:00</response_datetime>
          <request_reference_no xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">123456</request_reference_no>
        </transaction_header>
        <transaction_body xmlns="http://phaymobile.cardtekgroup/Tmm/ListMngmnt">
          <retrieval_reference_no xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">100000906983</retrieval_reference_no>
          <sms_noti xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">
            <toField>905308810677</toField>
            <valueField>MasterCard Mobile 24.02.2014 11:01:55: Your Accouts:  Account Name: My Mastercard:      Mobile Account Number:  520019******8447  (Funding Acct U02 Account Name: GarantiVPOS:      Mobile Account Number:  428220******8015  (Funding Acct U02.  Ref: 123456</valueField>
          </sms_noti>
          <list_items>
            <list_item xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">
              <list_item_name>My Mastercard</list_item_name>
              <prompt_cpin>Y</prompt_cpin>
              <list_item_value_1>520019******8447</list_item_value_1>
              <list_item_value_2>U02</list_item_value_2>
            </list_item>
            <list_item xmlns="http://phaymobile.cardtekgroup/Tmm/CommonTypes">
              <list_item_name>GarantiVPOS</list_item_name>
              <prompt_cpin>Y</prompt_cpin>
              <list_item_value_1>428220******8015</list_item_value_1>
              <list_item_value_2>U02</list_item_value_2>
            </list_item>
          </list_items>
        </transaction_body>
      </ViewListResult>
    </ViewListResponse>
  </soap:Body>
</soap:Envelope>`
	err := xml.Unmarshal([]byte(responseStr), &listManagement)
	if err != nil {
		fmt.Println("can not parse value from xml: ", responseStr, "error:", err.Error())
		t.Error("Can not parse from xml")

	}

	var emptyLM ListManagementEnvelope
	if listManagement.Body.Response.Result.TransactionHeader == emptyLM.Body.Response.Result.TransactionHeader {
		t.Error("Listmanagement is not filled")
	}

	if listManagement.Body.Response.Result.TransactionHeader.RequestDateTime != "0001-01-01T00:00:00" {
		t.Error("RequestDateTime can not be parsed")
	}

	listItemCount := len(listManagement.Body.Response.Result.TransactionBody.ListItems.ListItem)

	if listItemCount != 2 {

		t.Error("ListItem count is not 2, it is ", listItemCount, listManagement)
	}
}

func Test_CanGenerateSoapRequest(t *testing.T) {

	listManagementRequest := ListManagementRequest{}

	listManagementRequest.ClientId = "12345"
	listManagementRequest.DateTime = "11112223345"
	listManagementRequest.ReferenceNo = "4555"
	listManagementRequest.SendSms = "Y"
	listManagementRequest.SendSmsLanguage = "eng"

	listManagementRequest.Msisdn = "905051234567"
	listManagementRequest.ListType = "Phone"
	queryResult := listManagementRequest.GetQuery()

	if queryResult == "" {
		t.Error("can not generate list management soap request", queryResult)
	}
}
