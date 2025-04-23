package main

import (
	"fmt"
	test_builder "htestp/builder"
	httphandler "htestp/http_handler"
	"htestp/models"
)

func main() {

	fmt.Printf("start!\n")

	builder := test_builder.CreateNewBuilder()

	builder.AddStaticNode(httphandler.Request{
		Url:    "https://httpbin.org/json",
		Method: "GET",
	}).AddMatchStoreConstraint(
		[]string{"slideshow", "author"},
		"Yours Truly",
		models.TypeString,
		"authorName",
	)
	if !builder.Run() {
		fmt.Printf("FAILED!!!\n")
	}

	builder.PrintList()

	fmt.Printf("end!\n")
}
