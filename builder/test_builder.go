package test_builder

import (
	"fmt"
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

	new := &nodes.Get_Node{Request: request}
	if builder.head == nil {
		builder.head = new
		builder.current = builder.head
		return builder
	}

	builder.current.AddNode(new)
	builder.current = new

	return builder
}

func (builder *TestBuilder) AddConstraint(constraint models.Constraint) *TestBuilder {
	builder.current.AddConstraint(constraint)
	return builder
}

func (builder *TestBuilder) PrintList() {
	builder.printListHelper(builder.head)
}
func (builder *TestBuilder) printListHelper(node models.Node) {

	fmt.Printf("\n\t%+v\n", node.ToString())
	fmt.Printf("\t")

	if node == nil || !node.Successful() {
		return
	}
	for _, next := range node.GetNextNodes() {
		builder.printListHelper(next)
	}
	fmt.Printf("\n")
}

func (builder *TestBuilder) Run() bool {

	node := builder.head

	node.Execute(builder.client)
	if !node.Check() {
		return false
	}

	for _, nextNode := range node.GetNextNodes() {
		if !builder.runHelper(nextNode) {
			return false
		}
	}

	return true
}
func (builder *TestBuilder) runHelper(node models.Node) bool {

	if node == nil {
		return true
	}

	node.Execute(builder.client)
	if !node.Check() {
		return false
	}

	for _, nextNode := range node.GetNextNodes() {
		if !builder.runHelper(nextNode) {
			return false
		}
	}

	return true
}
