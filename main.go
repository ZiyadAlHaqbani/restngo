package main

import (
	"fmt"
	"htestp/builder"
	"htestp/models"
)

func main() {

	builder := builder.CreateNewBuilder()

	builder.AddStaticNode("https://openlibrary.org/search.json?q=test", models.GET, nil).
		AddFindStoreConstraint("title", models.TypeString, "test_var")

	println(builder.Run())

	builder.PrintList()

	fmt.Printf("END!")

	//program end
}
