# XAT

This is a simple Go program that can be used to test for XML External Entity (XXE) injection vulnerabilities.

## Prerequisites

- Go (1.13 or later)

## Usage 

To use `XAT`, you need to specify the URL of the vulnerable server as a command line argument using the `-u` flag.

```
    $ go run main.go -u http://vulnerable-server.com/xxe-endpoint
```

The program will send an XML payload containing an XXE injection to the specified URL using an HTTP POST request. If the server is vulnerable to XXE injection, the response from the server may include sensitive information.

## Warning

Please use this tool responsibly and only on systems that you have permission to test. Do not use this tool to perform unauthorized attacks on systems that you do not own or have permission to test.

## License 

- Licensed under `MIT`
