package subaccount

import (
	"Rave-go/rave"
	"Rave-go/rave/helper"
	// "strconv"
	"go/types"
)

type CreateSubaccount interface {
	CreateSubaccount(data CreateSubaccountData) (error error, response map[string]interface{})
}

type ListSubaccount interface {
	ListSubaccount(data ListSubaccountData) (error error, response map[string]interface{})
} 

type FetchSubaccount interface {
	FetchSubaccount(id string) (error error, response map[string]interface{})
}

type SubaccountInterface interface {
	CreateSubaccount
	ListSubaccount
	FetchSubaccount	
}

type CreateSubaccountData struct {
	AccountBank                   string         `json:"account_bank"`
	AccountNumber                 string         `json:"account_number"`
	BusinessName                  string         `json:"business_name"`
	BusinessEmail                 string         `json:"business_email"`
	BusinessContact               string         `json:"business_contact"`
	BusinessMobile                string         `json:"business_mobile"`
	BusinessContactMobile         string         `json:"business_contact_mobile"`
	Seckey                        string         `json:"seckey"`
	Meta                          types.Slice    `json:"meta"`
	SplitType                     string         `json:"split_type"`
	SplitValue                    string         `json:"split_value"`
}

type ListSubaccountData struct {
	AccountBank               string         `json:"account_bank"`
	AccountNumber             string         `json:"account_number"`
	BankName                  string         `json:"bank_name"`
	Seckey                    string         `json:"seckey"`
}

type Subaccount struct {
	rave.Rave
}

var noresponse = map[string]interface{}{
	"": "",
}

func (s Subaccount) CreateSubaccount(data CreateSubaccountData) (error error, response map[string]interface{}) {
	data.Seckey = s.GetSecretKey()
	url := s.GetBaseURL() + s.GetEndpoint("subaccount", "create")
	err, response := helper.MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (s Subaccount) ListSubaccount(data ListSubaccountData) (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"seckey":    s.GetSecretKey(),
		"account_number": data.AccountNumber,
		"account_bank": data.AccountBank,
		"bank_name": data.BankName,
	}
	url := s.GetBaseURL() + s.GetEndpoint("subaccount", "list")
	err, response := helper.MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (s Subaccount) FetchSubaccount(id string) (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"seckey": s.GetSecretKey(),	
	}
	url := s.GetBaseURL() + s.GetEndpoint("subaccount", "fetch") + "/" + id
	err, response := helper.MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (s Subaccount) DeleteSubaccount(id string) (error error, response map[string]interface{}) {
	paymentData := map[string]string{
		"seckey": s.GetSecretKey(),	
		"id": id,
	}

	url := s.GetBaseURL() + s.GetEndpoint("subaccount", "delete")
	err, response := helper.MakePostRequest(paymentData, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}