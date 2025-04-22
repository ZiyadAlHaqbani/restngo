package nodes

import (
	httphandler "htestp/http_handler"
	"htestp/models"
	"net/http"
)

type Get_Node struct {
	Next        models.Node
	Request     httphandler.Request
	Response    httphandler.HTTPResponse
	Constraints []models.Constraint
}

func (node *Get_Node) Execute(client *http.Client) (httphandler.HTTPResponse, error) {
	resp, err := httphandler.Handle(client, node.Request)
	if err != nil {
		return httphandler.HTTPResponse{}, err
	}

	return *resp, nil
}

func (node *Get_Node) Check() bool {
	for _, constraint := range node.Constraints {
		constraint.Constrain(node)
	}
}

// Execute() map[string]interface{}
// Check() bool
// GetResp() httphandler.HTTPResponse

type Post_Node struct {
	next        models.Node
	request     httphandler.Request
	response    httphandler.HTTPResponse
	constraints models.Constraint
}
