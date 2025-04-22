package nodes

import (
	httphandler "htestp/http_handler"
	"htestp/models"
	"net/http"
)

type Get_Node struct {
	Next         []models.Node
	Request      httphandler.Request
	Response     httphandler.HTTPResponse
	Constraints  []models.Constraint
	match_status models.MatchStatus
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
	var status models.MatchStatus
	for _, constraint := range node.Constraints {
		status = constraint.Constrain(node)
		if !status.Success {
			node.match_status = status
			return false
		}
	}
	node.match_status = status
	return true

}

func (node *Get_Node) GetResp() httphandler.HTTPResponse {
	return node.Response
}

func (node *Get_Node) AddConstraint(constraint models.Constraint) {
	node.Constraints = append(node.Constraints, constraint)
}

func (node *Get_Node) AddNode(new models.Node) {
	node.Next = append(node.Next, new)
}

func (node *Get_Node) GetNextNodes() []models.Node {
	return node.Next
}
