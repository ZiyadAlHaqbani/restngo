package constraints

import (
	"fmt"
	"htestp/models"
)

type Exist_Constraint struct {
	Field []string         //	e.g ["user", "age"] would access user.age
	Type  models.MatchType //	defines the expected type of the final field in the 'Field' array

	Status models.MatchStatus
}

func (match *Exist_Constraint) Constrain(node models.Node) models.MatchStatus {
	status := match.constrain(node)
	match.Status = status
	return status
}
func (match *Exist_Constraint) constrain(node models.Node) models.MatchStatus {
	resp := node.GetResp()

	var obj map[string]interface{}
	for _, key := range match.Field[:len(match.Field)-1] {
		var valid bool
		obj, valid = resp.Body[key].(map[string]interface{})
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("object: %+v doesn't include field: %s", obj, key),
				Failed_at_node: &node,
			}
		}
	}

	//	important to seperate the loop into two parts to acquire the final desired field name
	key := match.Field[len(match.Field)-1]
	actual, exists := obj[key] //	check if the field exists in the object or not
	if !exists {
		msg := ""
		for _, subField := range match.Field[:len(match.Field)-1] {
			msg += subField + "."
		}

		finalSubField := match.Field[len(match.Field)-1]
		msg += finalSubField

		return models.MatchStatus{
			Failed:         true,
			Message:        fmt.Sprintf("field: '%s' doesn't exist", msg),
			Failed_at_node: &node,
		}
	}

	switch match.Type {
	case models.TypeString:
		_, valid := actual.(string)
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s doesn't match expected type: %s", key, match.Type),
				Failed_at_node: &node,
			}
		}
	case models.TypeFloat:
		_, valid := actual.(float64)
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s doesn't match expected type: %s", key, match.Type),
				Failed_at_node: &node,
			}
		}
	case models.TypeBool:
		_, valid := actual.(bool)
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s with type: %T doesn't match expected type: %s", key, actual, match.Type),
				Failed_at_node: &node,
			}
		}
	case models.TypeArray:
		_, valid := actual.([]interface{})
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s with type: %T doesn't match expected type: %s", key, actual, match.Type),
				Failed_at_node: &node,
			}
		}
	case models.TypeObject:
		_, valid := actual.(map[string]interface{})
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s with type: %T doesn't match expected type: %s", key, actual, match.Type),
				Failed_at_node: &node,
			}
		}

	}

	return models.MatchStatus{
		Failed:         false,
		Failed_at_node: &node,
		MatchedValue:   actual,
		ValueType:      match.Type,
	}
}

func (match *Exist_Constraint) ToString() string {
	return "exst_cnstr_" + match.Status.ToString()
}
