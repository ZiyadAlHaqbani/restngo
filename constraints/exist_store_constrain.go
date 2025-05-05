package constraints

import (
	"htestp/models"
	"htestp/runner"
)

type Exist_Store_Constraint struct {
	InnerConstraint Exist_Constraint
	Varname         string
}

func (match *Exist_Store_Constraint) Constrain(node models.Node) models.MatchStatus {

	status := match.InnerConstraint.Constrain(node)
	if node.Successful() {
		newVar := models.TypedVariable{
			Value: status.MatchedValue,
			Type:  match.InnerConstraint.Type,
		}

		runner.StoreVariable(match.Varname, newVar)
	}
	return status
}

func (match *Exist_Store_Constraint) ToString() string {

	return "str_" + match.InnerConstraint.ToString()
}
