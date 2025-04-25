package main

import (
	"fmt"
	test_builder "htestp/builder"
	httphandler "htestp/http_handler"
	"htestp/models"
)

func main() {

	fmt.Printf("start!\n")

	builder := test_builder.CreateNewBuilder()

	// req_body := map[string]string{
	// 	"email":        "test@mail.com",
	// 	"passwordHash": "123456",
	// }

	// req_bytes, err := json.Marshal(req_body)
	// if err != nil {
	// 	panic("json ERROR")
	// }

	builder.
		AddStaticNode(httphandler.Request{
			Url:    "http://httpbin.org/json",
			Method: string(models.GET),
			// Body:   *bytes.NewBuffer(req_bytes),
		}).
		AddMatchStoreConstraint(
			[]string{"slideshow", "author"},
			"Yours Truly",
			models.TypeString,
			"authorName",
		).
		AddExistStoreConstraint(
			[]string{"slideshow", "author"}, models.TypeString, "authorExists",
		).
		AddDynamicNode("https://openlibrary.org/search.json", models.GET,
			func(m *map[string]models.TypedVariable) map[string]string {
				key := (*m)["authorName"]
				Map := map[string]string{}
				Map["q"] = key.Value.(string)
				return Map
			},
			nil).AddExistConstraint([]string{"q"}, models.TypeString).AddExistConstraint_("[12].q[12].q.ahmad[12][432432]", models.TypeArray)

	if !builder.Run() {
		fmt.Printf("FAILED!!!\n")
	} else {
		fmt.Printf("Success!!!\n")
	}

	builder.PrintList()

	fmt.Printf("end!\n")
}
