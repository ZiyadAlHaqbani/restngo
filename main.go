package main

import (
	"fmt"
	test_builder "htestp/builder"
	"htestp/models"
)

func main() {

	//start of the program
	builder1 := test_builder.CreateNewBuilder()
	builder1.
		AddStaticNode(
			"https://httpbin.org/json",
			models.GET,
			nil,
		)

	builder1.AddStaticNodeBranch("https://openlibrary.org/search.json", models.GET, nil)

	builder1.AddMatchStoreConstraint(
		"slideshow.author",
		"Yours Truly",
		models.TypeString,
		"authorName",
	).
		AddDynamicNode("https://openlibrary.org/search.json", models.GET,
			func(m *map[string]models.TypedVariable) map[string]string {
				key := (*m)["authorName"]
				Map := map[string]string{}
				Map["q"] = key.Value.(string)
				return Map
			}, nil).
		AddExistConstraint("docs[2].author_key[0]", models.TypeString).
		AddMatchStoreConstraint(
			"docs[2].author_key[0]",
			"OL3783157A",
			models.TypeString,
			"authorKey")

	//	each operation builds a new branch, with the parent being the previous builder's current branch
	//	the builder doesn't proceed to any branch and stays at current

	//	WARNING: when running the test, you must always start from the root builder
	status := builder1.Run()
	fmt.Printf("Test Passed: %v", status)
	builder1.PrintList()
	//program end
}
