package main

import (
	"fmt"
	"htestp/dsl/parser"
	"htestp/dsl/scanner"
	"htestp/runner/runner"
	"os"
)

func main() {

	res, err := os.ReadFile("example.restngo")
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	scanner := scanner.CreateScanner(string(res))

	parser := parser.CreateParser(scanner.Scan())
	parser.Parse()

	// runner.RunHelper(http.DefaultClient, parser.Head)

	branch_points := runner.GetBranches_(parser.Head)

	fmt.Printf("len: %d\n", len(branch_points))
	for _, point := range branch_points {
		fmt.Printf("%+v\n", point)
	}

	// scanner.CreateScanner()
	// parser.CreateParser()

	// builder := builder.CreateNewBuilder()

	// builder.AddStaticNode("https://openlibrary.org/search.json?q=test", models.GET, nil).
	// 	AddFindStoreConstraint("title", models.TypeString, "test_var").
	// 	AddDynamicNode("https://openlibrary.org/search.json", models.GET,
	// 		func(ctx *map[string]models.TypedVariable) url.Values {
	// 			params := url.Values{}
	// 			if variable := (*ctx)["test_var"]; variable.Type == models.TypeString {
	// 				params.Add("q", variable.Value.(string))
	// 			}
	// 			return params
	// 		}, nil)

	// builder.Run()

	// builder.PrintList()

	// fmt.Printf("END!")

}
