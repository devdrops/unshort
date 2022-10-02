package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Error: no arguments. Please provide a URL.")
	}
	url := os.Args[1]

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 10
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100
	httpClient := &http.Client{
		Timeout: time.Second * 5,
		Transport: transport,
	}

	res, err := httpClient.Get(url)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	fmt.Printf("REFERER:\t%s\n", res.Request.Header.Get("Referer"))
	fmt.Printf("LOCATION:\t%s\n", res.Request.URL)

	os.Exit(0)
}
