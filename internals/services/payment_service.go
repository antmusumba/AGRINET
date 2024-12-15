package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ProcessStkPush processes a STK push payment using hardcoded values
func ProcessStkPush(phoneNumber string, amount int) (string, error) {

	// Define the STK Push request URL
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	method := "POST"

	// Create the request payload with hardcoded values
	payload := strings.NewReader(fmt.Sprintf(`{
    "BusinessShortCode": 174379,
    "Password": "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMjQxMjE1MDYxODQ5",
    "Timestamp": "20241215061849",
    "TransactionType": "CustomerPayBillOnline",
    "Amount": %d,
    "PartyA": "%s",
    "PartyB": 174379,
    "PhoneNumber": "%s",
    "CallBackURL": "https://mydomain.com/path",
    "AccountReference": "AgriNet",
    "TransactionDesc": "Payment of Seedlings"
  }`, amount, phoneNumber, phoneNumber))

	// Create an HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}

	// Add headers to the request
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer MxIhRFOXIWnZ3DfJAceBDKzWdXnb")

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Return the response as a string
	return string(body), nil
}
