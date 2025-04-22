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
	node.Response = *resp
	return *resp, nil
}

// TODO: save the first failed constraint inside the node
func (node *Get_Node) Check() bool {
	for _, constraint := range node.Constraints {
		if !constraint.Constrain(node).Success {
			return false
		}
	}
	return true

}

// Execute() map[string]interface{}
// Check() bool
func (node *Get_Node) GetResp() httphandler.HTTPResponse {
	return node.Response
}

type Post_Node struct {
	next        models.Node
	request     httphandler.Request
	response    httphandler.HTTPResponse
	constraints models.Constraint
}
