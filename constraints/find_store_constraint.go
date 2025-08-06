package constraints

import (
	"htestp/models"
	"htestp/runner/context"
)

type Find_Store_Constraint struct {
	InnerConstraint Find_Constraint
	Varname         string
}

func (match *Find_Store_Constraint) Constrain(node models.Node) models.MatchStatus {

	status := match.InnerConstraint.Constrain(node)
	if node.Successful() {
		newVar := models.TypedVariable{
			Value: status.MatchedValue,
			Type:  match.InnerConstraint.Type,
		}

		context.StoreVariable(match.Varname, newVar)
	}
	return status
}

func (match *Find_Store_Constraint) ToString() string {

	return "str_" + match.InnerConstraint.ToString()
}
