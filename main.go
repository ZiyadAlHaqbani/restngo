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

	builder.AddStaticNode(httphandler.Request{
		Url:    "http://localhost:8080/auth/register?email=test@mail.com&passwordHash=123456",
		Method: string(models.POST),
		// Body:   *bytes.NewBuffer(req_bytes),
	}).AddMatchStoreConstraint(
		[]string{"slideshow"},
		"Yours Truly",
		models.TypeString,
		"authorName",
	).AddDynamicNode("http://google.com/search", models.GET,
		func(m *map[string]models.TypedVariable) map[string]string {
			key := (*m)["authorName"]
			Map := map[string]string{}
			Map["q"] = key.Value.(string)
			return Map
		},
		nil)

	if !builder.Run() {
		fmt.Printf("FAILED!!!\n")
	}

	builder.PrintList()

	fmt.Printf("end!\n")
}
