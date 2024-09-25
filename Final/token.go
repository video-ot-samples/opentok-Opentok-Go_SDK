package main

import (
    "encoding/base64"
    "crypto/hmac"
    "crypto/sha1"
    "encoding/hex"
    "fmt"
    "math/rand"
    "net/url"
    "time"
)

const (
    apiKey    = ""
    apiSecret = ""
    sessionId = ""
)

func main() {
    // Token Params
    timeNow := time.Now().Unix()
    expire := timeNow + 86400 // 1 day expiration
    role := "publisher"
    data := "bob"

    // Calculation
    rand.Seed(time.Now().UnixNano())
    nonce := rand.Int63()
    data = url.QueryEscape(data)
    dataString := fmt.Sprintf("session_id=%s&create_time=%d&expire_time=%d&role=%s&connection_data=%s&nonce=%d", sessionId, timeNow, expire, role, data, nonce)

    // Encryption
    mac := hmac.New(sha1.New, []byte(apiSecret))
    mac.Write([]byte(dataString))
    hash := mac.Sum(nil)

    // Encoding
    sig := hex.EncodeToString(hash)
    preCoded := fmt.Sprintf("partner_id=%s&sig=%s:%s", apiKey, sig, dataString)
    token := "T1==" + base64.StdEncoding.EncodeToString([]byte(preCoded))

    // Token Achieved
    fmt.Println("Generated Token:", token)
}
