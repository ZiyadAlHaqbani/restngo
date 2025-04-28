package constraints

import (
	"htestp/models"
)

type Match_Constraint struct {
	Field    string           //	e.g ["user", "age"] would access user.age
	Type     models.MatchType //	defines the type of the final field in the 'Field' array
	Expected interface{}      //	the expected value in the final field

	Status models.MatchStatus
}

func (match *Match_Constraint) Constrain(node models.Node) models.MatchStatus {

	status := match.constrain(node)
	match.Status = status
	return status
}
func (match *Match_Constraint) constrain(node models.Node) models.MatchStatus {

	resp := node.GetResp()

	result, msg := traverse(match.Field, resp.Body)
	if result == nil {
		return models.MatchStatus{
			Failed:  true,
			Message: msg}
	}

	status := checkType(result, match.Type, match.Field)
	status.Failed_at_node = &node
	return status
}

func (match *Match_Constraint) ToString() string {

	return "mtch_cnstr_" + match.Status.ToString()
}
