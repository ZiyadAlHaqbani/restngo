package test_builder

import (
	"bytes"
	"fmt"
	"htestp/constraints"
	httphandler "htestp/http_handler"
	"htestp/models"
	"htestp/nodes"
	"log"
	"net/http"
	"sync"
)

func CreateNewBuilder() *TestBuilder {
	// storage := new(map[string]models.TypedVariable)
	builder := &TestBuilder{
		head:    nil,
		current: nil,
		client:  http.DefaultClient,
		Storage: &map[string]models.TypedVariable{},
	}

	return builder
}

type TestBuilder struct {
	head    models.Node
	current models.Node
	client  *http.Client
	Storage *map[string]models.TypedVariable
}

// TODO: add AddStaticNodeRaw
func (builder *TestBuilder) AddStaticNode(url string, method models.HTTPMethod, body *bytes.Buffer) *TestBuilder {

	request := httphandler.Request{
		Url:    url,
		Method: string(method),
	}

	if body != nil {
		request.Body = *body
	}

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
//
//	dynamic nodes use data from the global context to build their own queries and bodies, must be noted that the query builder
//	assumes your url does not include a query
func (builder *TestBuilder) AddDynamicNode(url string, method models.HTTPMethod, queryBuilder func(*map[string]models.TypedVariable) map[string]string, bodyBuilder func(*map[string]models.TypedVariable) map[string]interface{}) *TestBuilder {

	new := &nodes.DynamicNode{
		InnerNode: nodes.StaticNode{Request: httphandler.Request{
			Url:    url,
			Method: string(method),
		}},
		QueryBuilderFunc: queryBuilder,
		BodyBuilderFunc:  bodyBuilder,
		Storage:          builder.Storage,
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

// WARNING: this is a dangerous function that shouldn't be used in most cases, it sets the current node of the caller
// builder to the callee builder, this will terminate the old branch if there is no builder refrencing it
func (builder *TestBuilder) SetBranchTo(callee *TestBuilder) {
	log.Print("WARNING: you called SetBranchTo() which can be unsafe, are you sure you need to use it?")
	builder.current = callee.current
}

func (builder *TestBuilder) AddStaticNodeBranch(url string, method models.HTTPMethod, body *bytes.Buffer) *TestBuilder {

	//	construct the new static node
	request := httphandler.Request{
		Url:    url,
		Method: string(method),
	}

	if body != nil {
		request.Body = *body
	}

	new := nodes.StaticNode{
		Request: request,
	}

	//	add the new node to the existing node tree
	//	but don't proceed to the next node
	builder.current.AddNode(&new)

	//	TODO: look into making the head the same , or add a new flag to determine if the builder is a branch
	//	TODO: builder to fix issues with using .Run() on a branch builder
	//	construct the new branch builder, this builder build on the existing tree but it starts from the new branch
	branchBuilder := TestBuilder{
		head:    &new,
		current: &new,
		client:  builder.client,
		Storage: builder.Storage,
	}

	return &branchBuilder
}

func (builder *TestBuilder) AddMatchConstraint(field string, expectedValue interface{}, expectedType models.MatchType) *TestBuilder {
	constraint := constraints.Match_Constraint{
		Field:    field,
		Type:     expectedType,
		Expected: expectedValue,
	}
	builder.current.AddConstraint(&constraint)
	return builder
}

func (builder *TestBuilder) AddMatchStoreConstraint(field string, expectedValue interface{}, expectedType models.MatchType, varname string) *TestBuilder {
	constraint := constraints.Match_Store_Constraint{
		InnerConstraint: constraints.Match_Constraint{
			Field:    field,
			Type:     expectedType,
			Expected: expectedValue,
		},
		Storage: builder.Storage,
		Varname: varname,
	}
	builder.current.AddConstraint(&constraint)
	return builder
}

func (builder *TestBuilder) AddExistConstraint(field string, expectedType models.MatchType) *TestBuilder {
	constraint := constraints.Exist_Constraint{
		Field: field,
		Type:  expectedType,
	}
	builder.current.AddConstraint(&constraint)
	return builder
}

func (builder *TestBuilder) AddExistStoreConstraint(field string, expectedType models.MatchType, varname string) *TestBuilder {
	constraint := constraints.Exist_Store_Constraint{
		InnerConstraint: constraints.Exist_Constraint{
			Field: field,
			Type:  expectedType,
		},
		Storage: builder.Storage,
		Varname: varname,
	}
	builder.current.AddConstraint(&constraint)
	return builder
}

func (builder *TestBuilder) PrintList() {
	builder.printListHelper(builder.head)
}
func (builder *TestBuilder) printListHelper(node models.Node) {

	type_str := ""
	switch node.(type) {
	case *nodes.StaticNode:
		type_str = "Static_Node"
	case *nodes.DynamicNode:
		type_str = "Dynamic_Node"
	}

	fmt.Printf("\n\t%s: %+v\n", type_str, node.ToString())
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

	if len(node.GetNextNodes()) == 1 {
		return builder.runHelper(node.GetNextNodes()[0])
	}

	// branches will still run even if a node in the level fails.
	success := true
	var success_mux sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(node.GetNextNodes()))
	for _, nextNode := range node.GetNextNodes() {
		go func() {
			defer wg.Done()
			Successful := builder.runHelper(nextNode)
			if !Successful {
				success_mux.Lock()
				defer success_mux.Unlock()
				success = false
			}
		}()
	}

	return success
}
