package nodes

import (
	httphandler "htestp/http_handler"
	"htestp/models"
	"net/http"
)

// Execute(client *http.Client) (httphandler.HTTPResponse, error)
// Check() bool
// GetResp() httphandler.HTTPResponse
// AddConstraint(Constraint)
// AddNode(Node)
// GetNextNodes() []Node
// ToString() string
// Successful() bool

type MockNode struct {
	ShouldSucceed bool
	ID            string
}

func (node *MockNode) Execute(client *http.Client) (httphandler.HTTPResponse, error) {

	return httphandler.HTTPResponse{}, nil

}

func (node *MockNode) Check() bool {
	return node.ShouldSucceed
}

func (node *MockNode) GetResp() httphandler.HTTPResponse {
	return httphandler.HTTPResponse{}
}

func (node *MockNode) AddConstraint(constraint models.Constraint) {

}

func (node *MockNode) AddNode(n models.Node) {

}

func (node *MockNode) ToString() string {
	return "mockNode"
}

func (node *MockNode) Successful() bool {
	return node.ShouldSucceed
}

func (node *MockNode) GetConstraints() []models.Constraint {
	return nil
}

func (node *MockNode) SetConstraints(constraints []models.Constraint) {
}

func (node *MockNode) GetRequest() httphandler.Request {
	return httphandler.Request{}
}

func (node *MockNode) SetRequest(request httphandler.Request) {
}

func (node *MockNode) GetNextNodes() []models.Node {
	return []models.Node{}
}

func (node *MockNode) SetNextNodes(next []models.Node) {
}

func (node *MockNode) GetID() string {
	return node.ID
}
