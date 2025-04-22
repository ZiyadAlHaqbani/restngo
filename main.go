package main

import (
	"fmt"
	httphandler "htestp/http_handler"
	"htestp/nodes"
	"net/http"
)

func main() {

	fmt.Printf("start!\n")

	var client *http.Client = http.DefaultClient

	var node = nodes.Get_Node{
		Next: nil,
		Request: httphandler.Request{
			Url:    "https://httpbin.org/bearer",
			Method: "GET",
		},
		Constraints: nil,
	}
	resp, err := node.Execute(client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %+s", resp.ToString())
	fmt.Printf("end!\n")

}
