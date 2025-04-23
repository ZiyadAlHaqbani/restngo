package nodes

import (
	"bytes"
	"encoding/json"
	"fmt"
	httphandler "htestp/http_handler"
	"htestp/models"
	"net/http"
)

type DynamicNode struct {
	InnerNode        StaticNode
	QueryBuilderFunc func(storage map[string]models.TypedVariable) map[string]string
	BodyBuilderFunc  func(storage map[string]models.TypedVariable) map[string]interface{}
	storage          map[string]models.TypedVariable
}

// Execute(client *http.Client) (httphandler.HTTPResponse, error)
// Check() bool
// GetResp() httphandler.HTTPResponse
// AddConstraint(Constraint)
// GetNextNodes() []Node
// AddNode(Node)
// ToString() string
// Successful() bool

func (node *DynamicNode) Execute(client *http.Client) (httphandler.HTTPResponse, error) {

	request_params := ""

	var request_body *bytes.Buffer

	if node.QueryBuilderFunc != nil {
		firstParam := true
		params := node.QueryBuilderFunc(node.storage)
		for key, value := range params {
			if firstParam {
				firstParam = false
				request_params = fmt.Sprintf("%s?%s=%s", request_params, key, value)
			}
			request_params = fmt.Sprintf("%s&%s=%s", request_params, key, value)
		}
	}

	if node.BodyBuilderFunc != nil {
		body := node.BodyBuilderFunc(node.storage)
		byteArray, err := json.Marshal(body)
		if err != nil {
			return httphandler.HTTPResponse{}, fmt.Errorf("failed to encode the generated request body")
		}

		request_body = bytes.NewBuffer(byteArray)

	}

	node.InnerNode.Request.Body = *request_body
	node.InnerNode.Request.Url += request_params

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

func (node *DynamicNode) GetNextNodes() []models.Node {
	return node.InnerNode.GetNextNodes()
}

func (node *DynamicNode) AddNode(new models.Node) {
	node.InnerNode.AddNode(new)
}

func (node *DynamicNode) ToString() string {
	temp := "Dynamic Node: "
	temp = fmt.Sprintf("%s%s", temp, node.InnerNode.ToString())
	return temp
}

func (node *DynamicNode) Successful() bool {
	return node.InnerNode.Successful()
}
