package gocoingecko

import (
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	// "encoding/json"
)

const API_BASE_URL = "https://api.coingecko.com/api/v3"

func NewClient(apiBaseURL string, timeout int) (*Client, error) {
	return &Client{
		Client: &http.Client{},
		ApiBaseURL: API_BASE_URL,
		Timeout: timeout,
	}, nil
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)
}
func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func sendRequest(client *http.Client, method string) []byte {
	endpoint := "https://api.coingecko.com/api/v3/simple/price?ids=ethereum%2Ccardano&vs_currencies=usd"
    //   values := map[string]string{"ids": "cardano"}
	// jsonData, err := json.Marshal(values)

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	return body
}

