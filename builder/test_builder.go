package test_builder

import (
	"fmt"
	"htestp/constraints"
	httphandler "htestp/http_handler"
	"htestp/models"
	"htestp/nodes"
	"log"
	"net/http"
)

func CreateNewBuilder() *TestBuilder {
	builder := &TestBuilder{
		head:    nil,
		current: nil,
		client:  http.DefaultClient,
		Storage: map[string]models.TypedVariable{},
	}

	return builder
}

type TestBuilder struct {
	head    models.Node
	current models.Node
	client  *http.Client
	Storage map[string]models.TypedVariable
}

// TODO: change to take url and method
// TODO: add AddStaticNodeRaw
func (builder *TestBuilder) AddStaticNode(request httphandler.Request) *TestBuilder {

	new := &nodes.StaticNode{Request: request}
	if builder.head == nil {
		builder.head = new
		builder.current = builder.head
		return builder
	}

	builder.current.AddNode(new)
	builder.current = new

	return builder
}

// TODO: add AddDynamicNodeRaw
func (builder *TestBuilder) AddDynamicNode(url string, method models.HTTPMethod, queryBuilder func(*map[string]models.TypedVariable) map[string]string, bodyBuilder func(*map[string]models.TypedVariable) map[string]interface{}) *TestBuilder {

	new := &nodes.DynamicNode{
		InnerNode: nodes.StaticNode{Request: httphandler.Request{
			Url:    url,
			Method: string(method),
		}},
		QueryBuilderFunc: queryBuilder,
		BodyBuilderFunc:  bodyBuilder,
		Storage:          &builder.Storage,
	}

	if builder.head == nil {
		builder.head = new
		builder.current = builder.head
		return builder
	}
	builder.current.AddNode(new)
	builder.current = new

	return builder
}

func (builder *TestBuilder) AddMatchConstraint(fields []string, expectedValue interface{}, expectedType models.MatchType) *TestBuilder {
	constraint := constraints.Match_Constraint{
		Field:    fields,
		Type:     expectedType,
		Expected: expectedValue,
	}
	builder.current.AddConstraint(&constraint)
	return builder
}

func (builder *TestBuilder) AddMatchStoreConstraint(fields []string, expectedValue interface{}, expectedType models.MatchType, varname string) *TestBuilder {
	constraint := constraints.Match_Store_Constraint{
		InnerConstraint: constraints.Match_Constraint{
			Field:    fields,
			Type:     expectedType,
			Expected: expectedValue,
		},
		Storage: &builder.Storage,
		Varname: varname,
	}
	builder.current.AddConstraint(&constraint)
	return builder
}

func (builder *TestBuilder) AddExistsConstraint(fields []string, expectedType models.MatchType) *TestBuilder {
	constraint := constraints.Exist_Constraint{
		Field: fields,
		Type:  expectedType,
	}
	builder.current.AddConstraint(&constraint)
	return builder
}

// func (builder *TestBuilder) AddMatchStoreConstraint(fields []string, expectedValue interface{}, expectedType models.MatchType, varname string) *TestBuilder {
// 	constraint := constraints.Match_Store_Constraint{
// 		InnerConstraint: constraints.Match_Constraint{
// 			Field:    fields,
// 			Type:     expectedType,
// 			Expected: expectedValue,
// 		},
// 		Storage: &builder.Storage,
// 		Varname: varname,
// 	}
// 	builder.current.AddConstraint(&constraint)
// 	return builder
// }

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
	return builder.runHelper(node)

}
func (builder *TestBuilder) runHelper(node models.Node) bool {

	if node == nil {
		return true
	}

	_, err := node.Execute(builder.client)
	if err != nil {
		log.Print(err)
		return false
	}
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
