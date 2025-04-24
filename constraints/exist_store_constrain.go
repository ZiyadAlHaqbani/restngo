package constraints

import "htestp/models"

type Exist_Store_Constraint struct {
	InnerConstraint Exist_Constraint
	Storage         *map[string]models.TypedVariable
	Varname         string
}

func (match *Exist_Store_Constraint) Constrain(node models.Node) models.MatchStatus {
	status := match.InnerConstraint.Constrain(node)
	if node.Successful() {
		(*match.Storage)[match.Varname] = models.TypedVariable{
			Value: status.MatchedValue,
			Type:  match.InnerConstraint.Type,
		}
	}
	return status
}
