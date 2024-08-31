package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
	`"id": "d-16088952994",
	"username": "jokdjcag",
	"password": "d2ewjr847ln5",
	"proxy_address": "31.223.189.90",
	"port": 6356,
	"valid": true,
	"last_verification": "2024-08-31T00:23:51.753164-07:00",
	"country_code": "PL",
	"city_name": "Warsaw",
	"asn_name": "Hostroyale Technologies Pvt Ltd",
	"asn_number": 44144,
	"high_country_confidence": false,
	"created_at": "2024-08-26T20:28:46.823984-07:00"


*/

type ProxyResult struct {
	Id                    string `json:"id"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	ProxyAddress          string `json:"proxy_address"`
	Port                  int    `json:"port"`
	Valid                 bool   `json:"valid"`
	LastVerification      string `json:"last_verification"`
	CountryCode           string `json:"country_code"`
	CityName              string `json:"city_name"`
	ASNNumber             int    `json:"asn_number"`
	HighCountryConfidence bool   `json:"high_country_confidence"`
	CreatedAt             string `json:"created_at"`
}

type ResponseBody struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []ProxyResult `json:"results"`
}

func main() {
	var (
		req *http.Request
	)
	req, err := http.NewRequest(http.MethodGet, "https://proxy.webshare.io/api/v2/proxy/list/?mode=direct&page_size=100&page=1", nil)
	if err != nil {
		panic(err)
	}
	httpClient := http.DefaultClient
	req.Header.Set("Authorization", "Token 5yxyhppmpwtdwkq78f5gctx18lgjq3du11k7zh3w")

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var input ResponseBody
	if err = json.Unmarshal(respBody, &input); err != nil {
		log.Println("can't unmarshal. Err: ", err)
		return
	}
	// socks5h://belet_proxy:asd43fdsa@94.136.185.149:1080
	var resultStr string
	for _, result := range input.Results {
		if len(resultStr) == 0 {
			resultStr = fmt.Sprintf("[ { \"URL\": \"socks5h://%s:%s@%s:%d\", \n \"countryCode\": \"%s\" }", result.Username, result.Password, result.ProxyAddress, result.Port, result.CountryCode)
		} else {
			resultStr = fmt.Sprintf("%s,\n{ \"URL\": \"socks5h://%s:%s@%s:%d\", \n \"countryCode\": \"%s\" }", resultStr, result.Username, result.Password, result.ProxyAddress, result.Port, result.CountryCode)
		}
	}
	fmt.Println(resultStr)
}
