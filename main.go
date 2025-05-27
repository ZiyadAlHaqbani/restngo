package main

import (
	"fmt"
	"htestp/builder"
	"htestp/models"
	"net/url"
)

func main() {

	builder := builder.CreateNewBuilder()

	builder.AddStaticNode("https://openlibrary.org/search.json?q=test", models.GET, nil).
		AddExistStoreConstraint("title", models.TypeString, "test_var").
		AddDynamicNode("https://openlibrary.org/search.json", models.GET,
			func(ctx *map[string]models.TypedVariable) url.Values {
				params := url.Values{}
				if variable := (*ctx)["test_var"]; variable.Type == models.TypeString {
					params.Add("q", variable.Value.(string))
				}
				return params
			}, nil)

	builder.Run()

	builder.PrintList()

	// json, err := json.MarshalIndent(builder, "", " ")
	// if err != nil {
	// 	println(err)
	// }

	// fmt.Printf("%s", json)

	fmt.Printf("END!")

	//program end
}
