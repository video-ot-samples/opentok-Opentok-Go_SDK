package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const (
	apiKey    = ""
	apiSecret = ""
	jsonWebToken = ""
)

type SessionResponse struct {
	SessionID string `json:"session_id"`
}

func createSession() string {
	// Set the OpenTok session creation URL
	URL := "https://api.opentok.com/session/create"

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

	// Parse the JSON response
	var sessionResponses []SessionResponse
	err = json.Unmarshal(body, &sessionResponses)
	if err != nil {
		panic(err)
	}

	if len(sessionResponses) == 0 {
		panic("no session ID found in response")
	}

	return sessionResponses[0].SessionID
}


func generateToken(sessionID string) string {
	// Token Params
	timeNow := time.Now().Unix()
	expire := timeNow + 86400 // 1 day expiration
	role := "publisher"
	data := "bob"

	// Calculation
	rand.Seed(time.Now().UnixNano())
	nonce := rand.Int63()
	data = url.QueryEscape(data)
	dataString := fmt.Sprintf("session_id=%s&create_time=%d&expire_time=%d&role=%s&connection_data=%s&nonce=%d", sessionID, timeNow, expire, role, data, nonce)

	// Encryption
	mac := hmac.New(sha1.New, []byte(apiSecret))
	mac.Write([]byte(dataString))
	hash := mac.Sum(nil)

	// Encoding
	sig := hex.EncodeToString(hash)
	preCoded := fmt.Sprintf("partner_id=%s&sig=%s:%s", apiKey, sig, dataString)
	token := "T1==" + base64.StdEncoding.EncodeToString([]byte(preCoded))

	return token
}

func main() {
	sessionID := createSession()
	fmt.Println("Generated Session ID:", sessionID)
	token := generateToken(sessionID)
	fmt.Println("Generated Token:", token)
}
