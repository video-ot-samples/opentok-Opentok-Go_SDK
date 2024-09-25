package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Set the OpenTok session creation URL
	URL := "https://api.opentok.com/session/create"

	// Replace "jwt_string" with your actual JSON web token
	jsonWebToken := ""
	// Define the data string to be sent in the request body
	dataStr := "location=140.228.54.78"

	// Create a new HTTP request with the POST method and request body
	req, err := http.NewRequest("POST", URL, bytes.NewBufferString(dataStr))
	if err != nil {
		panic(err)
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-OPENTOK-AUTH", jsonWebToken)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Print the complete response JSON string
	fmt.Println("Response:", string(body))
}
