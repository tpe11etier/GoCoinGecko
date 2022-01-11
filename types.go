package gocoingecko

import "net/http"


type (
	//Represents a Coingecko API
	Client struct {
		Client *http.Client
		ApiBaseURL string
		Timeout  int
	}
)