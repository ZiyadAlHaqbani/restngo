package nodes

import (
	"fmt"
	httphandler "htestp/http_handler"
	"htestp/models"
	"net/http"
)

type StaticNode struct {
	Next         []models.Node
	Request      httphandler.Request
	Response     httphandler.HTTPResponse
	Constraints  []models.Constraint
	match_status models.MatchStatus
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
			node.match_status = status
			return false
		}
	}
	node.match_status = status
	return true

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

func (node *StaticNode) GetNextNodes() []models.Node {
	return node.Next
}

func (node *StaticNode) ToString() string {
	temp := fmt.Sprintf("%s_%s", node.Request.Method, node.Request.Url)
	temp = fmt.Sprintf("%s %s", temp, node.match_status.ToString())
	// resp := node.GetResp()
	// temp = fmt.Sprintf("%s\n%s", temp, resp.ToString())

	return temp
}

func (node *StaticNode) Successful() bool {
	return !node.match_status.Failed
}
