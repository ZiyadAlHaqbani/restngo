package constraints

import (
	"htestp/models"
	profilers "htestp/profiler"
)

type Exist_Store_Constraint struct {
	InnerConstraint Exist_Constraint
	Storage         *map[string]models.TypedVariable
	Varname         string
}

func (match *Exist_Store_Constraint) Constrain(node models.Node) models.MatchStatus {
	defer profilers.ProfileScope("Exist_Store_Constraint::Constrain")()
	status := match.InnerConstraint.Constrain(node)
	if node.Successful() {
		(*match.Storage)[match.Varname] = models.TypedVariable{
			Value: status.MatchedValue,
			Type:  match.InnerConstraint.Type,
		}
	}
	return status
}

func (match *Exist_Store_Constraint) ToString() string {
	defer profilers.ProfileScope("Exist_Store_Constraint::ToString")()
	return "str_" + match.InnerConstraint.ToString()
}
