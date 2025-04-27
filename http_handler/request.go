package httphandler

import (
	"bytes"
	"net/http"
)

type Request struct {
	Url           string
	Method        string
	Header        *http.Header
	Body          *bytes.Buffer
	ContentLength int64
	Retries       int
}
