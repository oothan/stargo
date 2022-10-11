package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Currencylayer struct {
	Success   bool   `json:"success"`
	Terms     string `json:"terms"`
	Privacy   string `json:"privacy"`
	Timestamp int    `json:"timestamp"`
	Source    string `json:"source"`
	Quotes    struct {
		Usdeur float64 `json:"USDEUR"`
		Usdinr float64 `json:"USDINR"`
	} `json:"quotes"`
}

func main() {
	AccessKey := "rv4uoAIG3wtDjHX7ThPEXJMFYEF7uOrc"
	url := fmt.Sprintf("http://apilayer.net/api/live?access_key=%s&currencies=EUR,INR&source=USD", AccessKey)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Currencylayer

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println(record.Quotes.Usdeur, record.Quotes.Usdinr)
	// We want to know what 100 Indian Rupee (INR) is worth in Euros (EUR).
	// The formula is USDEUR/USDINR*100
	fmt.Println("Rs. 100 is worth in Euros = ", (record.Quotes.Usdeur / record.Quotes.Usdinr * 100.0))
}
