package main

import ( 
	"github.com/tpe11etier/GoCoinGecko/client"
	"net/http"
	"log"
)

const API_BASE_URL = "https://api.coingecko.com/api/v3"

func main() {
	// c should be re-used for further calls
	  c, err := gocoingecko.NewClient(API_BASE_URL,10)
	  response := c.Send(c, http.MethodGet)
	//   response := gocoingecko.SendRequest(c, http.MethodGet)
	  log.Println("Response Body:", string(response))
  }