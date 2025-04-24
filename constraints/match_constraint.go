package constraints

import (
	"fmt"
	"htestp/models"
)

type Match_Constraint struct {
	Field    []string         //	e.g ["user", "age"] would access user.age
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
	actual := obj[key]

	switch match.Type {
	case models.TypeString:
		value := match.Expected.(string)
		actual_value, valid := actual.(string)
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s doesn't match expected type: %s", key, match.Type),
				Failed_at_node: &node,
			}
		}
		if value != actual_value {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s doesn't match expected value: %s", key, value),
				Failed_at_node: &node,
			}
		}
	case models.TypeFloat:
		value := match.Expected.(float64)
		actual_value, valid := actual.(float64)
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s doesn't match expected type: %s", key, match.Type),
				Failed_at_node: &node,
			}
		}
		if value != actual_value {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s doesn't match expected value: %f", key, value),
				Failed_at_node: &node,
			}
		}
	case models.TypeBool:
		value := match.Expected.(bool)
		actual_value, valid := actual.(bool)
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s doesn't match expected type: %s", key, match.Type),
				Failed_at_node: &node,
			}
		}
		if value != actual_value {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s doesn't match expected value: %b", key, value),
				Failed_at_node: &node,
			}
		}
	case models.TypeArray:
		value := match.Expected.([]interface{})
		actual_value, valid := actual.([]interface{})
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s doesn't match expected type: %s", key, match.Type),
				Failed_at_node: &node,
			}
		}
		valid, msg := matchLists(actual_value, value)
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        msg,
				Failed_at_node: &node,
			}
		}
	case models.TypeObject:
		value := match.Expected.(map[string]interface{})
		actual_value, valid := actual.(map[string]interface{})
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        fmt.Sprintf("field: %s doesn't match expected type: %s", key, match.Type),
				Failed_at_node: &node,
			}
		}
		valid, msg := matchMaps(actual_value, value)
		if !valid {
			return models.MatchStatus{
				Failed:         true,
				Message:        msg,
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

func (match *Match_Constraint) ToString() string {
	return "mtch_cnstr_" + match.Status.ToString()
}
