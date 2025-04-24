package constraints

import "htestp/models"

type Exist_Constraint struct {
	Field []string         //	e.g ["user", "age"] would access user.age
	Type  models.MatchType //	defines the expected type of the final field in the 'Field' array
}

func (match *Exist_Constraint) Constrain(node models.Node) models.MatchStatus {
	panic("TODO: implement Exist_Constraint.Constrain()")
}
