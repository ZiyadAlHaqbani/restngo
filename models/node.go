package models

import (
	httphandler "htestp/http_handler"
	"net/http"
)

type Node interface {
	//	url httphandler.Request
	//	response map[string]interface{}
	//	constraints Constraint[]
	//	next Node

	Execute(client *http.Client) (httphandler.HTTPResponse, error)
	Check() bool
	GetResp() httphandler.HTTPResponse
}
