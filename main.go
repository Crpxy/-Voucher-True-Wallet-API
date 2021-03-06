package main

import (
	"net/http"
  "log"
  "io/ioutil"
  "time"
  "encoding/json"
  "bytes"
)

var (
  voucher = "VOUCHER HERE"
  mobile = "MOBILE HERE"
)

func main() {
  url := "https://gift.truemoney.com/campaign/vouchers/" + voucher + "/redeem"
  
  requestBody, err := json.Marshal(map[string]string{
    "mobile": mobile,
    "voucher_hash": voucher,
  })

  if err != nil {
    log.Fatalln(err)
  }
  
  tr := &http.Transport{
    ForceAttemptHTTP2: true,
    MaxIdleConns: 0,
    MaxConnsPerHost: 0,
    MaxResponseHeaderBytes: 65535,
    MaxIdleConnsPerHost: 65535,
    IdleConnTimeout: 10 * time.Second,
    TLSHandshakeTimeout: 10 * time.Second,
    ExpectContinueTimeout: 10 * time.Second,
    WriteBufferSize: 65535,
  }
  
  client := http.Client{
    Timeout: 10 * time.Second,
    Transport: tr,
  }
  
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
  request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36 Edg/87.0.664.66")
  request.Header.Set("Content-type", "application/json")

  if err != nil {
    log.Fatalln(err)
  }

  resp, err := client.Do(request)

  if err != nil {
    log.Fatalln(err)
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    log.Fatalln(err)
  }

  log.Println(string(body))
}
