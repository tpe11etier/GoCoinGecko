package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/tpe11etier/gocoingecko/types"
)

var baseURL = "https://api.coingecko.com/api/v3"

type Client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{httpClient: httpClient}
}

// Ping /ping endpoint
func (c *Client) Ping() (*types.Ping, error) {
	url := fmt.Sprintf("%s/ping", baseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *types.Ping
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MakeReq HTTP Request Helper
func (c *Client) MakeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	resp, err := sendReq(req, c.httpClient)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Send Request after Make
func sendReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}


func (c *Client) GetSimplePrice(ids []string, vsCurrencies []string) (*map[string]map[string]float32, error) {
	params := url.Values{}
	idsParam := strings.Join(ids[:], ",")
	vsCurrenciesParam := strings.Join(vsCurrencies[:], ",")

	params.Add("ids", idsParam)
	params.Add("vs_currencies", vsCurrenciesParam)

	url := fmt.Sprintf("%s/simple/price?%s", baseURL, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	t := make(map[string]map[string]float32)
	err = json.Unmarshal(resp, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}


func (c *Client) GetCoinsMarkets(vsCurrency string, options *types.CoinsMarketOptions) (*types.CoinsMarket, error) {
	params := url.Values{}

	params.Add("vs_currency", vsCurrency)
	
	if options != nil {
		params.Add("ids", options.Ids)
		params.Add("category", options.Category)
		params.Add("order", options.Order)
		params.Add("per_page", strconv.Itoa(options.PerPage))
		params.Add("page", strconv.Itoa(options.Page))
		params.Add("sparkline", strconv.FormatBool(options.Sparkline))
		params.Add("price_change_percentage", options.PriceChangePercentage)
	}

	url := fmt.Sprintf("%s/coins/markets?%s", baseURL, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	var data *types.CoinsMarket
	err = json.Unmarshal(resp, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
	


}