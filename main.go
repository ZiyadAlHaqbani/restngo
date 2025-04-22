package main

import (
	"fmt"
	test_builder "htestp/builder"
	"htestp/constraints"
	httphandler "htestp/http_handler"
)

func main() {

	fmt.Printf("start!\n")

	builder := test_builder.CreateNewBuilder()

	builder.AddGetNode(httphandler.Request{
		Url:     "https://httpbin.com/json",
		Method:  "GET",
		Retries: 5,
	}).
		AddConstraint(&constraints.Match_Constraint{
			Field:    []string{"slideshow", "author"},
			Type:     constraints.TypeString,
			Expected: "Yours Truly",
		}).AddConstraint(&constraints.Match_Constraint{
		Field:    []string{"slideshow", "me"},
		Type:     constraints.TypeString,
		Expected: "Yours Truly",
	}).Run()

	fmt.Printf("end!\n")
}
