package gocoingecko


func main() {
	// c should be re-used for further calls
	  c := httpClient()
	  response := sendRequest(c, http.MethodGet)
	  log.Println("Response Body:", string(response))
  }