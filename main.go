package main

import (
	"fmt"
	// "fmt"
	"log"

	// "net/http"

	gecko "github.com/tpe11etier/gocoingecko/client"
)

const API_BASE_URL = "https://api.coingecko.com/api/v3"

func main() {
	cg := gecko.NewClient(nil)
	res, err := cg.GetSimplePrice([]string{"bitcoin,cardano"}, []string{"usd"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

  }