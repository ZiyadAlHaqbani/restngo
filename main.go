package main

import (
	"fmt"
	"htestp/nodes"
	"net/http"
)

func main() {

	fmt.Printf("start!\n")

	var client *http.Client = http.DefaultClient

	var node = nodes.Get_Node{}

	fmt.Printf("end!\n")

}
