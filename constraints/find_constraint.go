package constraints

import "htestp/models"

type Find_Constraint struct {
	//	this constraint finds the first most shallow match for the field, and not the exact match
	//	for example: an object with field obj.field.field and constraint with field 'field' will *only* match
	//	the first field, since its the most shallow.

	Field string           //	in this context, the field only needs a partial match
	Type  models.MatchType //	defines the expected type of the final field in the 'Field' array

	Status models.MatchStatus
}

func (constraint *Find_Constraint) Constrain(node models.Node) models.MatchStatus {

	obj, err_message := partialTraverse_type(constraint.Field, node.GetResp().Body, constraint.Type)
	if obj == nil {
		return models.MatchStatus{
			Failed:         true,
			Message:        err_message,
			Failed_at_node: &node,
		}
	}

	status := checkType(obj, constraint.Type, constraint.Field)

	if status.Failed {
		status.Failed_at_node = &node
	}

	return status
}

func (constraint *Find_Constraint) ToString() string {
	panic("implement constraints.Find_Constraint.ToString()")
}
