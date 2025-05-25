package main

import (
	"fmt"
	"htestp/builder"
	"htestp/models"
)

func main() {

	builder := builder.CreateNewBuilder()

	builder.AddStaticNode("https://openlibrary.org/search.json?q=test", models.GET, nil).
		AddFindConstraint("name", models.TypeString)

	println(builder.Run())

	builder.PrintList()

	fmt.Printf("END!")

	//program end
}
