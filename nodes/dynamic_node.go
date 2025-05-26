package nodes

import (
	"bytes"
	"encoding/json"
	"fmt"
	httphandler "htestp/http_handler"
	"htestp/models"
	"htestp/runner/context"
	"net/http"
	"net/url"
)

type DynamicNode struct {
	InnerNode StaticNode
	// TODO: Use url values instead of map[string]string
	QueryBuilderFunc func(storage *map[string]models.TypedVariable) url.Values             `json:"-"`
	BodyBuilderFunc  func(storage *map[string]models.TypedVariable) map[string]interface{} `json:"-"`
	Next             []models.Node
}

// Execute(client *http.Client) (httphandler.HTTPResponse, error)
// Check() bool
// GetResp() httphandler.HTTPResponse
// AddConstraint(Constraint)
// GetNextNodes() []Node
// AddNode(Node)
// ToString() string
// Successful() bool

// DEPRECATED
// func sanitizeQuery(query map[string]string) string {
// 	//	e.g value = "My name is Ahmad" can't be included in url queries.
// 	//	"My name is Ahmad" -> "My+name+is+Ahmad" can be included in url queries.
// 	params := url.Values{}
// 	for key, value := range query {
// 		params.Add(key, value)
// 	}
// 	return params.Encode()
// }

func (node *DynamicNode) Execute(client *http.Client) (httphandler.HTTPResponse, error) {

	if node.QueryBuilderFunc != nil {

		params := node.QueryBuilderFunc(&context.Storage)
		request_params := params.Encode()
		node.InnerNode.Request.Url += "?" + request_params
	}

	if node.BodyBuilderFunc != nil {
		body := node.BodyBuilderFunc(&context.Storage)
		byteArray, err := json.Marshal(body)
		if err != nil {
			return httphandler.HTTPResponse{}, fmt.Errorf("failed to encode the generated request body")
		}

		request_body := bytes.NewBuffer(byteArray)
		node.InnerNode.Request.Body = request_body
	}

	return node.InnerNode.Execute(client)
}

func (node *DynamicNode) Check() bool {

	return node.InnerNode.Check()
}

func (node *DynamicNode) GetResp() httphandler.HTTPResponse {

	return node.InnerNode.GetResp()
}

func (node *DynamicNode) AddConstraint(constraint models.Constraint) {

	node.InnerNode.AddConstraint(constraint)
}

func (node *DynamicNode) AddNode(new models.Node) {

	node.Next = append(node.Next, new)
}

func (node *DynamicNode) ToString() string {
	temp := "Dynamic Node: "
	temp = fmt.Sprintf("%s%s", temp, node.InnerNode.ToString())
	return temp
}

func (node *DynamicNode) Successful() bool {
	// Returns the opposite of flag failed
	return node.InnerNode.Successful()
}

func (node *DynamicNode) GetConstraints() []models.Constraint {
	return node.InnerNode.Constraints
}

func (node *DynamicNode) SetConstraints(constraints []models.Constraint) {
	node.InnerNode.Constraints = constraints
}

func (node *DynamicNode) GetRequest() httphandler.Request {
	return node.InnerNode.Request
}

func (node *DynamicNode) SetRequest(request httphandler.Request) {
	node.InnerNode.Request = request
}

func (node *DynamicNode) GetNextNodes() []models.Node {
	return node.Next
}

func (node *DynamicNode) SetNextNodes(next []models.Node) {
	node.Next = next
}

// GetConstraints() []Constraint
// SetConstraints([]Constraint)

// GetRequest() httphandler.Request
// SetRequest(httphandler.Request)

// GetNextNodes() []Node
// SetNextNodes([]Node)

func (node *DynamicNode) GetID() string {
	return node.InnerNode.GetID()
}
