# iUP Http Client - GoLang Version

Simple library that implements http(s) client

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Example

Create a file example.go

```
package main

import (
	"fmt"

	httpclient "github.com/paulofilipem/iup-http-client-util"
)

func main() {

	fmt.Println("\nSending a Simple HTTP Request")
	httpclient.Url = "http://google.com"
	httpclient.Method = "GET"
	httpclient.HeadersValue = &[]string{"Content-Type: text/html", "Accept: text/html"}
	httpclient.Send()
	if httpclient.ResponseCode == 200 {
		fmt.Println("OK, StatusCode OK - 200")
	} else {
		fmt.Println("Error, StatusCode:" + string(httpclient.ResponseCode))
	}

	for _, value := range *httpclient.ResponseHeaders {
		fmt.Println("Reponse Header: ", value)
	}

	fmt.Println("\nSending a Simple HTTPS Request")
	httpclient.Url = "https://google.com"
	httpclient.Method = "GET"
	httpclient.HeadersValue = &[]string{"Content-Type: text/html", "Accept: text/html"}
	httpclient.Send()
	if httpclient.ResponseCode == 200 {
		fmt.Println("OK, StatusCode OK - 200")
	} else {
		fmt.Println("Error, StatusCode:" + string(httpclient.ResponseCode))
	}

	fmt.Println("\nSending a Request using SSL Client Auth")
	httpclient.Url = "https://google.com"
	httpclient.Method = "GET"
	httpclient.HeadersValue = &[]string{"Content-Type: text/html", "Accept: text/html"}
	httpclient.Send()
	if httpclient.ResponseCode == 200 {
		fmt.Println("OK, StatusCode OK - 200")
	} else {
		fmt.Println("Error, StatusCode:" + string(httpclient.ResponseCode))
	}

	fmt.Println("\nSending a multiple request, using cookie storage(on memory)")
	var baseUrl = "https://google.com"
	httpclient.Url = baseUrl + "/authenticate"
	httpclient.Method = "POST"
	httpclient.HeadersValue = &[]string{"Content-Type: text/html", "Accept: text/html"}
	httpclient.Send()
	if httpclient.ResponseCode == 200 {
		fmt.Println("OK, StatusCode OK - 200")
	} else {
		fmt.Println("Error, StatusCode:" + string(httpclient.ResponseCode))
	}

	httpclient.Test()

}
```

And run

```
go run example.go
```

## Contributing

Feel free :D

## Authors

* **Paulo Filipe Macedo @paulofilipem** 

## License

This project is licensed under the MIT License - see the [LICENSE.md](http://en.wikipedia.org/wiki/MIT_License) file for details
