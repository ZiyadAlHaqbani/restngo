package models

import httphandler "htestp/http_handler"

type Node interface {
	//	url httphandler.Request
	//	response map[string]interface{}
	//	constraints constraint[]

	Execute() map[string]interface{}
	Check() bool
	GetResp() httphandler.HTTPResponse
}
