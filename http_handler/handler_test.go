package httphandler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleWithRequestWithNoBody(t *testing.T) {

	_, err := Handle(http.DefaultClient, Request{
		Url:    "https://httpbin.org",
		Method: "GET",
	})

	assert.Nil(t, err, "expected bodyless request to not return error")

}

func TestHandleWithRequestWithBody(t *testing.T) {

	test_body := map[string]interface{}{
		"q": 12,
	}

	json, err := json.Marshal(test_body)
	assert.Nil(t, err, "marshalling test_body failed")

	_, err = Handle(http.DefaultClient, Request{
		Url:    "https://openlibrary.com/search.json",
		Method: "POST",
		Body:   bytes.NewBuffer(json),
	})
	assert.Nil(t, err, "expected request with body to not return error")

}
