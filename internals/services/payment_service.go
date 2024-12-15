package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Fake credentials
// TODO: replace with real credentials and store in environment variables
const (
	password  = "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMjQxMjE1MDMxMDA4"
	ShortCode = 174379
)

// ProcessStkPush processes a STK push payment
func ProcessStkPush(phoneNumber string, amount int) (string, error) {

	// API endpoint and method
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	method := "POST"

	// Dynamic payload structure
	payload := map[string]interface{}{
		"BusinessShortCode": ShortCode,
		"Password":          password,
		"Timestamp":         "20241215031008",
		"TransactionType":   "CustomerPayBillOnline",
		"Amount":            amount,
		"PartyA":            phoneNumber,
		"PartyB":            ShortCode,
		"PhoneNumber":       phoneNumber,
		"CallBackURL":       "https://mydomain.com/path",
		"AccountReference":  "CompanyXLTD",
		"TransactionDesc":   "Payment of X",
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Create an HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	// Add headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer srADXA5t9cwR4MeWkVSrn4qree4T")

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Read and print the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
