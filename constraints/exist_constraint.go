package constraints

import (
	"htestp/models"
)

type Exist_Constraint struct {
	Field string
	Type  models.MatchType //	defines the expected type of the final field in the 'Field' array

	Status models.MatchStatus
}

//	func (match *Exist_Constraint) Constrain(node models.Node) models.MatchStatus {
//		status := match.constrain(node)
//		match.Status = status
//		return status
//	}
func (exist *Exist_Constraint) Constrain(node models.Node) models.MatchStatus {

	status := exist.constrain(node)
	exist.Status = status
	return status
}

func (exist *Exist_Constraint) constrain(node models.Node) models.MatchStatus {

	resp := node.GetResp()

	result, msg := traverse(exist.Field, resp.Body)
	if result == nil {
		return models.MatchStatus{
			Failed:  true,
			Message: msg}
	}

	status := checkType(result, exist.Type, exist.Field)
	status.Failed_at_node = &node
	return status
}

func (match *Exist_Constraint) ToString() string {

	return "exst_cnstr " + match.Status.ToString()
}
