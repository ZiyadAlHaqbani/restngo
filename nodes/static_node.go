package nodes

import (
	"fmt"
	httphandler "htestp/http_handler"
	"htestp/models"
	"net/http"
)

type StaticNode struct {
	ID          string
	Next        []models.Node
	Request     httphandler.Request
	Response    httphandler.HTTPResponse
	Constraints []models.Constraint
	Failed      bool

	INTERNAL bool // flag to define whether or not the current node is only used as an internal node for a wrapper node type
}

func (node *StaticNode) Execute(client *http.Client) (httphandler.HTTPResponse, error) {
	resp, err := httphandler.Handle(client, node.Request)
	if err != nil {
		return httphandler.HTTPResponse{}, err
	}
	node.Response = *resp
	return *resp, nil
}

// TODO: save the first failed constraint inside the node
func (node *StaticNode) Check() bool {
	var status models.MatchStatus
	for _, constraint := range node.Constraints {
		status = constraint.Constrain(node)
		if status.Failed {
			node.Failed = true
		}
	}
	return !node.Failed

}

func (node *StaticNode) GetResp() httphandler.HTTPResponse {
	return node.Response
}

func (node *StaticNode) AddConstraint(constraint models.Constraint) {
	node.Constraints = append(node.Constraints, constraint)
}

func (node *StaticNode) AddNode(new models.Node) {
	node.Next = append(node.Next, new)
}

func (node *StaticNode) ToString() string {

	var temp string
	if !node.INTERNAL {
		temp = "Static Node: "
	}
	temp = fmt.Sprintf("%s(ID: %s), %s_%s", temp, node.ID, node.Request.Method, node.Request.Url)

	if len(node.Constraints) > 0 {
		temp += " {"
		for _, constr := range node.Constraints {
			temp += constr.ToString() + ", "
		}
		temp += "}"
	}

	return temp
}

func (node *StaticNode) Successful() bool {
	return !node.Failed
}

func (node *StaticNode) GetConstraints() []models.Constraint {
	return node.Constraints
}

func (node *StaticNode) SetConstraints(constraints []models.Constraint) {
	node.Constraints = constraints
}

func (node *StaticNode) GetRequest() httphandler.Request {
	return node.Request
}

func (node *StaticNode) SetRequest(request httphandler.Request) {
	node.Request = request
}

func (node *StaticNode) GetNextNodes() []models.Node {
	return node.Next
}

func (node *StaticNode) SetNextNodes(next []models.Node) {
	node.Next = next
}

func (node *StaticNode) GetID() string {
	return node.ID
}
