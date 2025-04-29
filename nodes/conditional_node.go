package nodes

import (
	"fmt"
	httphandler "htestp/http_handler"
	"htestp/models"
	"log"
	"net/http"
)

type ConditionalNode struct {
	FirstPathSucceeded  bool
	SecondPathSucceeded bool

	FirstBranchNode  models.Node //	if the first path succeeds, then GetNextNodes will return this node.
	SecondBranchNode models.Node //	if the second path succeeds, then GetNextNodes will return this node.

	// Next        []models.Node
	// Request     httphandler.Request
	// Response    httphandler.HTTPResponse
	// Constraints []models.Constraint
	// Failed      bool

}

func (node *ConditionalNode) Execute(client *http.Client) (httphandler.HTTPResponse, error) {

	//	execute the first path
	resp_1, err_1 := node.FirstBranchNode.Execute()
	if err_1 == nil {
		return resp_1, nil
	}

	log.Printf("DEBUG: first branch of conditional node failed: %+v", err_1)

	//	execute the second path
	resp_2, err_2 := node.FirstBranchNode.Execute()
	if err_2 == nil {
		return resp_2, nil
	}

	return httphandler.HTTPResponse{}, err_2
}

func (node *ConditionalNode) Check() bool {
	var status models.MatchStatus
	for _, constraint := range node.Constraints {
		status = constraint.Constrain(node)
		if status.Failed {
			node.Failed = true
		}
	}
	return !node.Failed

}

func (node *ConditionalNode) GetResp() httphandler.HTTPResponse {
	return node.Response
}

func (node *ConditionalNode) AddConstraint(constraint models.Constraint) {
	node.Constraints = append(node.Constraints, constraint)
}

func (node *ConditionalNode) AddNode(new models.Node) {
	node.Next = append(node.Next, new)
}

func (node *ConditionalNode) GetNextNodes() []models.Node {
	if node.FirstPathSucceeded {
		return []models.Node{node.FirstBranchNode}
	} else if node.SecondPathSucceeded {
		return []models.Node{node.SecondBranchNode}
	}

	return []models.Node{node.FirstBranchNode}
}

func (node *ConditionalNode) ToString() string {
	temp := fmt.Sprintf("%s_%s", node.Request.Method, node.Request.Url)

	if len(node.Constraints) > 0 {
		temp += " {"
		for _, constr := range node.Constraints {
			temp += constr.ToString() + ", "
		}
		temp += "}"
	}
	// resp := node.GetResp()
	// temp = fmt.Sprintf("%s\n%s", temp, resp.ToString())

	return temp
}

func (node *ConditionalNode) Successful() bool {
	return (node.FirstBranchNode.Successful() && node.SecondBranchNode.Successful())
}

// Execute(client *http.Client) (httphandler.HTTPResponse, error)
// Check() bool
// GetResp() httphandler.HTTPResponse
// AddConstraint(Constraint)
// AddNode(Node)
// GetNextNodes() []Node
// ToString() string
// Successful() bool
