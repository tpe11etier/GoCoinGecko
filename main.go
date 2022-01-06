package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const APIBASEURL = "https://api.coingecko.com/api/v3"

type CoinGeckoAPI struct {
	apiBaseURL string
	HTTPClient *http.Client 
}

func NewCoinGeckoAPI() *CoinGeckoAPI {
	return &CoinGeckoAPI{
		apiBaseURL: APIBASEURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type Coin struct {
	coin map[string]interface{}

}

type errorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	coin map[string]interface{} 
}

func (c *CoinGeckoAPI) GetSimplePrice(ctx context.Context) (*Coin, error) {
	req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/simple/price?ids=ethereum%2Ccardano&vs_currencies=usd",nil)
    fmt.Printf("Request Body %v", req)
	if err != nil {
        return nil, err
    }
	req = req.WithContext(ctx)

	res := Coin{}
    if err := c.sendRequest(req, &res); err != nil {
        return nil, err
    }
	return &res, nil
}


func (c *CoinGeckoAPI) sendRequest(req *http.Request, v interface{}) error {
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    req.Header.Set("Accept", "application/json; charset=utf-8")
    // req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

    res, err := c.HTTPClient.Do(req)
    if err != nil {
        return err
    }

    defer res.Body.Close()

    if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
        var errRes errorResponse
        if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
            return errors.New(errRes.Message)
        }

        return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
    }

    fullResponse := Coin{}

    if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
        return err
    }

    return nil
}


func getSimplePrice() string {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=ethereum%2Ccardano&vs_currencies=usd"

	resp, getErr := http.Get(url)
	
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	
	if readErr != nil {
		log.Fatal(readErr)
	}

	return string(body)
}

func main() {
	// fmt.Println("You are in 'main'")
	// s := getSimplePrice()	
	// c := Coin{}
	// if err := json.Unmarshal([]byte(s), &c.coin); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", c)

	var client = NewCoinGeckoAPI()
	ctx := context.Background()
	res, err := client.GetSimplePrice(ctx) 
	if err != nil {
		fmt.Println(err)
	}
	// println(res)
	fmt.Printf("%+v", res)


}