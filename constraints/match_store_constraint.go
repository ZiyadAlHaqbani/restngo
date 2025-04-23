package constraints

import "htestp/models"

type Match_Store_Constraint struct {
	InnerConstraint Match_Constraint
	Storage         *map[string]models.TypedVariable // Reference to global storage
	Varname         string                           // varname represents the name of the variable used for storage mapping in global storage.
}

func (match *Match_Store_Constraint) Constrain(node models.Node) models.MatchStatus {
	status := match.InnerConstraint.Constrain(node)
	if status.Success {
		(*match.Storage)[match.Varname] = models.TypedVariable{
			Value: match.InnerConstraint.Expected,
			Type:  match.InnerConstraint.Type,
		}
	}
	return status
}
