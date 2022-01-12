package gocoingecko

import (
	// "fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"encoding/json"
)

const API_BASE_URL = "https://api.coingecko.com/api/v3"

func NewClient(apiBaseURL string, timeout int) (*Client, error) {
	return &Client{
		Client: &http.Client{},
		ApiBaseURL: API_BASE_URL,
		Timeout: timeout,
	}, nil
}


func HttpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

// Send makes a request to the API, the response body will be
// unmarshalled into v, or if v is an io.Writer, the response will
// be written to it without decoding
func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	// Set default headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en_US")

	// Default values for headers
	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}
	// if c.returnRepresentation {
	// 	req.Header.Set("Prefer", "return=representation")
	// }

	resp, err = c.Client.Do(req)
	// c.log(req, resp)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp := &ErrorResponse{Response: resp}
		data, err = ioutil.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			json.Unmarshal(data, errResp)
		}

		return errResp
	}
	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		io.Copy(w, resp.Body)
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

func SendRequest(client *http.Client, method string) []byte {
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

