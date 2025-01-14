/*
	vuldb_api_demo - Go VulDB API Demo

	License: GPL-3.0
    	Required Dependencies: 
        * bytes
        * fmt
        * io/ioutil
        * net/http
    	Optional Dependencies: None
*/

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// API URL
	url := "https://vuldb.com/?api"

	// Headers for authentication
	personalAPIKey := "" // Enter your personal API key here
	userAgent := "VulDB API Advanced Go Demo Agent"

	// Request body parameters
	recent := "5"		// Default is 5
	details := "0"		// Default is 0
	id := ""		// Example: "290848", Default: id := ""
	cve := ""		// Example: "CVE-2024-1234", Default: cve := ""

	// Construct the request body
	var requestBody string
	if id == "" && cve == "" {
		requestBody = fmt.Sprintf("recent=%s&details=%s", recent, details)
	} else if cve != "" {
		requestBody = fmt.Sprintf("search=%s&details=%s", cve, details)
	} else {
		requestBody = fmt.Sprintf("id=%s&details=%s", id, details)
	}

	// Create HTTP request
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-VulDB-ApiKey", personalAPIKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Output response
	if resp.StatusCode == 200 {
		fmt.Println("Response:", string(body))
	} else {
		fmt.Printf("Request failed with response code: %d\n", resp.StatusCode)
	}
}
