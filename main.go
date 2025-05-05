package main

import (
	"fmt"
	"htestp/dsl/scanner"
)

func main() {

	builder1 := test_builder.CreateNewBuilder()
	builder1.
		AddStaticNode(
			"https://httpbin.org/json",
			models.GET,
			nil,
		).
		AddMatchStoreConstraint(
			"slideshow.author",
			"Yours Truly",
			models.TypeString,
			"authorName",
		).
		AddDynamicNode("https://openlibrary.org/search.json", models.GET,
			func(m *map[string]models.TypedVariable) url.Values {
				key := (*m)["authorName"]
				params := url.Values{}
				params.Set("q", key.Value.(string))
				return params
			}, nil).
		AddExistConstraint("docs[2].author_key[0]", models.TypeString).
		AddMatchStoreConstraint(
			"docs[2].author_key[0]",
			"OL3783157A",
			models.TypeString,
			"authorKey")

	builder1.AddDynamicNodeBranch(
		"https://httpbin.org/get",
		models.GET,
		func(m *map[string]models.TypedVariable) url.Values {
			authorKey := (*m)["authorKey"]
			params := url.Values{}
			params.Set("author_key", authorKey.Value.(string))
			return params
		},
		nil,
	).AddExistConstraint("args.author_key", models.TypeString)


	s := scanner.CreateScanner(source)
	s.Scan()

	testScanner := scanner.CreateScanner(s.ToString())
	testScanner.Scan()
	fmt.Printf("%s\n", s.ToString())
	fmt.Printf("%s\n", testScanner.ToString())
	fmt.Printf("%+v", testScanner.ToString() == s.ToString())
	//program end
}
