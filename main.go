package main

import (
	"fmt"
	test_builder "htestp/builder"
	"htestp/constraints"
	httphandler "htestp/http_handler"
	"htestp/models"
)

func main() {

	fmt.Printf("start!\n")

	builder := test_builder.CreateNewBuilder()

	builder.AddStaticNode(httphandler.Request{
		Url:    "https://httpbin.org/json",
		Method: "GET",
	}).AddConstraint(&constraints.Match_Constraint{
		Field:    []string{"slideshow", "author"},
		Type:     models.TypeString,
		Expected: "Yours Truly",
	}).AddStaticNode(httphandler.Request{
		Url:    "https://pokeapi.co/api/v2/pokemon/ditto",
		Method: "GET",
	}).AddConstraint(&constraints.Match_Constraint{
		Field:    []string{"height"},
		Type:     models.TypeFloat,
		Expected: 3.0,
	})

	if !builder.Run() {
		fmt.Printf("FAILED!!!\n")
	}

	builder.PrintList()

	fmt.Printf("end!\n")
}
