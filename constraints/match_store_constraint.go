package constraints

import (
	"htestp/models"
	"htestp/runner"
)

type Match_Store_Constraint struct {
	InnerConstraint Match_Constraint
	Varname         string // varname represents the name of the variable used for storage mapping in global storage.
}

func (match *Match_Store_Constraint) Constrain(node models.Node) models.MatchStatus {

	status := match.InnerConstraint.Constrain(node)
	if !status.Failed {
		newVar := models.TypedVariable{
			Value: status.MatchedValue,
			Type:  match.InnerConstraint.Type,
		}

		runner.StoreVariable(match.Varname, newVar)

	}
	return status
}

func (match *Match_Store_Constraint) ToString() string {

	return "str_" + match.InnerConstraint.ToString()
}
