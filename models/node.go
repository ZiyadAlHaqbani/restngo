package models

import (
	httphandler "htestp/http_handler"
	"net/http"
)

type Node interface {
	Execute(client *http.Client) (httphandler.HTTPResponse, error)
	Check() bool
	GetResp() httphandler.HTTPResponse
	AddConstraint(Constraint)
	AddNode(Node)
	GetNextNodes() []Node
	ToString() string
	Successful() bool
}
