package constraints

import (
	"htestp/models"
	profilers "htestp/profiler"
)

type Match_Store_Constraint struct {
	InnerConstraint Match_Constraint
	Storage         *map[string]models.TypedVariable // Reference to global storage
	Varname         string                           // varname represents the name of the variable used for storage mapping in global storage.
}

func (match *Match_Store_Constraint) Constrain(node models.Node) models.MatchStatus {
	defer profilers.ProfileScope("Match_Store_Constraint::Constrain")()
	status := match.InnerConstraint.Constrain(node)
	if !status.Failed {
		(*match.Storage)[match.Varname] = models.TypedVariable{
			Value: status.MatchedValue,
			Type:  match.InnerConstraint.Type,
		}
	}
	return status
}

func (match *Match_Store_Constraint) ToString() string {
	defer profilers.ProfileScope("Match_Store_Constraint::ToString")()
	return "str_" + match.InnerConstraint.ToString()
}
