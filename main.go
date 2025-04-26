package main

import (
	"fmt"
	test_builder "htestp/builder"
	"htestp/models"
)

func main() {

	//start of the program
	builder := test_builder.CreateNewBuilder()
	builder.
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
			func(m *map[string]models.TypedVariable) map[string]string {
				key := (*m)["authorName"]
				Map := map[string]string{}
				Map["q"] = key.Value.(string)
				return Map
			}, nil).
		AddExistConstraint("docs[2].author_key[0]", models.TypeString)

	//	each operation builds a new branch, with the parent being the previous builder's current branch
	//	the builder doesn't proceed to any branch and stays at current
	branch1 := builder.AddStaticNodeBranch("http://httpbin.org/b1", models.GET, nil)
	branch2 := builder.AddStaticNodeBranch("http://httpbin.org/b2", models.GET, nil)
	branch3 := builder.AddStaticNodeBranch("http://httpbin.org/b3", models.GET, nil)

	//	branch objects are also test builders, and can be used in the same manner
	branch1.AddStaticNode("http://httpbin.org/b11", models.GET, nil).AddExistConstraint("num[12]", models.TypeFloat)
	branch2.AddStaticNode("http://httpbin.org/b22", models.GET, nil).AddExistConstraint("num[12]", models.TypeFloat)
	branch3.AddStaticNode("http://httpbin.org/b33", models.GET, nil).AddExistConstraint("num[12]", models.TypeFloat)

	//	WARNING: when running the test, you must always start from the root builder
	status := builder.Run()
	fmt.Printf("Test Passed: %v", status)

	//	option to print the contents of the test
	//	each branch will print only the contents of its branch

	branch1.PrintList()
	branch2.PrintList()
	branch3.PrintList()

	//	example of the unsafe function SetBranchTo(), in most cases it shouldn't be used, as it can lead to unintentional unallocation of nodes
	builder.SetBranchTo(branch3)

	//program end
}
