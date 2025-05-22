package main

import (
	"fmt"
	"htestp/builder"
	"htestp/models"
)

func main() {

	builder := builder.CreateNewBuilder()

	builder.AddStaticNode("https://openlibrary.org/search.json?q=test", models.GET, nil).
		AddFindConstraint("[0]", models.TypeString)

	print(builder.Run())

	fmt.Printf("END!")

	//program end
}
