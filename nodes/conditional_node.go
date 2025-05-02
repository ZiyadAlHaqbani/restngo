package nodes

import (
	httphandler "htestp/http_handler"
	"htestp/models"
	"net/http"
)

// TODO:	add a way to create node objects that don't contain constraints
// TODO:	find a way to integrate this node into test_builder, most
// TODO:	likely through a seperate conditionalNodeBuilder?
type ConditionalNode struct {
	TrueNode  models.Node //	if the callback return true, this node will run.
	FalseNode models.Node //	if the callback return false, this node will run.

	ConditionFunc func(map[string]models.TypedVariable) bool //	callback that determines which node will be run,
	branch        bool                                       //	whether the callback returns true or false
	Storage       *map[string]models.TypedVariable           //	global context

	Next        []models.Node
	Constraints []models.Constraint
	Failed      bool
}

func (node *ConditionalNode) Execute(client *http.Client) (httphandler.HTTPResponse, error) {
	panic("TODO: fully implement the conditional node, including adding constraints to children at runtime")
	node.branch = node.ConditionFunc(*node.Storage)

	if node.branch {
		return node.TrueNode.Execute(client)
	} else {
		return node.FalseNode.Execute(client)
	}
}

func (node *ConditionalNode) Check() bool {

	var desiredNode models.Node
	if node.branch {
		desiredNode = node.TrueNode
	} else {
		desiredNode = node.FalseNode
	}

	//	to avoid duplicate assignment of constraints, we assign constraint after checking condition

	var status models.MatchStatus
	for _, constraint := range node.Constraints {
		status = constraint.Constrain(desiredNode)
		if status.Failed {
			node.Failed = true
		}

	}
	return !node.Failed

}

func (node *ConditionalNode) GetResp() httphandler.HTTPResponse {
	if node.branch {
		return node.TrueNode.GetResp()
	} else {
		return node.FalseNode.GetResp()
	}
}

func (node *ConditionalNode) AddConstraint(constraint models.Constraint) {
	node.Constraints = append(node.Constraints, constraint)
}

func (node *ConditionalNode) AddNode(new models.Node) {
	node.Next = append(node.Next, new)
}

func (node *ConditionalNode) GetNextNodes() []models.Node {
	return node.Next
}

func (node *ConditionalNode) ToString() string {

	temp := "ConditionalNode:"

	if node.branch {
		temp = "\n\t" + node.TrueNode.ToString()
	} else {
		temp = "\n\t" + node.FalseNode.ToString()
	}

	return temp
}

func (node *ConditionalNode) Successful() bool {
	return node.Failed
}
