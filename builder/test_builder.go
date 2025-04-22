package test_builder

import (
	httphandler "htestp/http_handler"
	"htestp/models"
	"htestp/nodes"
	"net/http"
)

func CreateNewBuilder() *TestBuilder {
	builder := &TestBuilder{
		head:    nil,
		current: nil,
		client:  http.DefaultClient,
	}

	return builder
}

type TestBuilder struct {
	head    models.Node
	current models.Node
	client  *http.Client
}

func (builder *TestBuilder) AddGetNode(request httphandler.Request) *TestBuilder {

	if builder.head == nil {
		builder.head = &nodes.Get_Node{
			Request: request,
		}
		builder.current = builder.head
	}

	var node models.Node = &nodes.Get_Node{Request: request}
	builder.current.AddNode(node)
	builder.current = node

	return builder
}

func (builder *TestBuilder) AddConstraint(constraint models.Constraint) *TestBuilder {
	builder.current.AddConstraint(constraint)
	return builder
}

func (builder *TestBuilder) Run() bool {

	node := builder.head

	node.Execute(builder.client)

	for _, nextNode := range node.GetNextNodes() {
		if !builder.RunHelper(nextNode) {
			return false
		}
	}

	return true
}

func (builder *TestBuilder) RunHelper(node models.Node) bool {

	if node == nil {
		return true
	}

	node.Execute(builder.client)

	for _, nextNode := range node.GetNextNodes() {
		if !builder.RunHelper(nextNode) {
			return false
		}
	}

	return true
}
