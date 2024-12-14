package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Fake credentials
// TODO: replace with real credentials and store in environment variables
const (
	password  = "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMjQxMjE0MTgyOTQ0"
	ShortCode = 174379
)

// ProcessStkPush processes a STK push payment
func ProcessStkPush(phoneNumber string, amount int) (string, error) {
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf(`{
    "BusinessShortCode": %d,
    "Password": %s,
    "Timestamp": "20160216165627",
    "TransactionType": "CustomerPayBillOnline",
    "Amount": %d,
    "PartyA": 254708374149,
    "PartyB": 174379,
    "PhoneNumber": %s,
    "CallBackURL": "https://mydomain.com/path",
    "AccountReference": "CompanyXLTD",
    "TransactionDesc": "Payment of X"
  }`, ShortCode, password, amount, phoneNumber))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer Odn8HiE1c2JeVjqWt1oh7OByDB4a")

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return string(body), nil
}
