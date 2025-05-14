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
	ToString() string
	Successful() bool
  GetID() string
  
	// getters and setters(makes behavior too OOP like, but essential to fix some issues)
	GetConstraints() []Constraint
	SetConstraints([]Constraint)

	GetRequest() httphandler.Request
	SetRequest(httphandler.Request)

	GetNextNodes() []Node
	SetNextNodes([]Node)

	

}
