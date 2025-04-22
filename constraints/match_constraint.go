package constraints

import (
	"fmt"
	"htestp/models"
	"strconv"
)

type MatchType string

const (
	TypeString MatchType = "string"
	TypeFloat  MatchType = "float64"
	TypeBool   MatchType = "bool"
	TypeObject MatchType = "map[string]interface{}"
	TypeArray  MatchType = "[]interface{}"
)

type Match_Constraint struct {
	Field    []string    //	e.g ["user", "age"] would access user.age
	Type     MatchType   //	defines the type of the final field in the 'Field' array
	Expected interface{} //	the expected value in the final field
}

func matchLists(a []interface{}, b []interface{}) (bool, string) {
	
	if len(a) != len(b){
		return false
	}

	for i, value := range a {
		switch v := value.(type) {
		case string:
			expected, valid := b[i].(string)
			if !valid
		case float64:
			temp = fmt.Sprintf("%s\n%s%d: %f", temp, tabs(depth), key, v)
		case bool:
			temp = fmt.Sprintf("%s\n%s%d: %b", temp, tabs(depth), key, v)
		case []interface{}:
			temp = resp.resolveInterfaceList(temp, strconv.Itoa(key), v, depth+1)
		case map[string]interface{}:
			temp = resp.resolveInterfaceMap(temp, strconv.Itoa(key), v, depth+1)
		default:
			temp = fmt.Sprintf("%s\n%s%d: unknown type", tabs(depth), temp, key)
		}
	}

	return true
}

func (match *Match_Constraint) Constrain(node models.Node) models.MatchStatus {
	resp := node.GetResp()

	var obj map[string]interface{}
	for _, key := range match.Field[:len(match.Field)-1] {
		var valid bool
		obj, valid = resp.Body[key].(map[string]interface{})
		if !valid {
			return models.MatchStatus{
				Success:        false,
				Message:        fmt.Sprintf("object: %+v doesn't include field: %s", obj, key),
				Failed_at_node: &node,
			}
		}
	}

	key := match.Field[len(match.Field)-1]

	actual := obj[key]

	switch match.Type {
	case TypeString:
		value := match.Expected.(string)
		actual_value, valid := actual.(string)
		if !valid {
			return models.MatchStatus{
				Success:        false,
				Message:        fmt.Sprintf("field: %s doesn't match expected type: %s", key, match.Type),
				Failed_at_node: &node,
			}
		}
		if value != actual_value {
			return models.MatchStatus{
				Success:        false,
				Message:        fmt.Sprintf("field: %s doesn't match expected value: %s", key, value),
				Failed_at_node: &node,
			}
		}
	case TypeFloat:
		value := match.Expected.(float64)
		actual_value, valid := actual.(float64)
		if !valid {
			return models.MatchStatus{
				Success:        false,
				Message:        fmt.Sprintf("field: %s doesn't match expected type: %s", key, match.Type),
				Failed_at_node: &node,
			}
		}
		if value != actual_value {
			return models.MatchStatus{
				Success:        false,
				Message:        fmt.Sprintf("field: %s doesn't match expected value: %f", key, value),
				Failed_at_node: &node,
			}
		}
	case TypeBool:
		value := match.Expected.(bool)
		actual_value, valid := actual.(bool)
		if !valid {
			return models.MatchStatus{
				Success:        false,
				Message:        fmt.Sprintf("field: %s doesn't match expected type: %s", key, match.Type),
				Failed_at_node: &node,
			}
		}
		if value != actual_value {
			return models.MatchStatus{
				Success:        false,
				Message:        fmt.Sprintf("field: %s doesn't match expected value: %b", key, value),
				Failed_at_node: &node,
			}
		}
	case TypeArray:
		value := match.Expected.([]interface{})
		actual_value, valid := actual.([]interface{})
		if !valid {
			return models.MatchStatus{
				Success:        false,
				Message:        fmt.Sprintf("field: %s doesn't match expected type: %s", key, match.Type),
				Failed_at_node: &node,
			}
		}
		if matchLists(actual_value, value) {
			return models.MatchStatus{
				Success:        false,
				Message:        fmt.Sprintf("field: %s doesn't match expected value: %f", key, value),
				Failed_at_node: &node,
			}
		}

	}
	return models.MatchStatus{
		// Message:        fmt.Sprintf("matched field: %v with value: %s", match.Field, match.Type),
		Success:        true,
		Failed_at_node: &node,
	}
}
