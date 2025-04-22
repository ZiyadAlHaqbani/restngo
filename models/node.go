package models

import (
	httphandler "htestp/http_handler"
	"net/http"
)

type Node interface {
	// Next         []models.Node
	// Request      httphandler.Request
	// Response     httphandler.HTTPResponse
	// Constraints  []models.Constraint
	// match_status models.MatchStatus

	Execute(client *http.Client) (httphandler.HTTPResponse, error)
	Check() bool
	GetResp() httphandler.HTTPResponse
	AddConstraint(Constraint)
	AddNode(Node)
	GetNextNodes() []Node
}
