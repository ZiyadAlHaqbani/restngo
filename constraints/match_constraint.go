package constraints

import (
	"fmt"
	"htestp/models"
)

type Match_Constraint struct {
	Field    []string         //	e.g ["user", "age"] would access user.age
	Type     models.MatchType //	defines the type of the final field in the 'Field' array
	Expected interface{}      //	the expected value in the final field
}

func matchLists(a []interface{}, b []interface{}) (bool, string) {

	if len(a) != len(b) {
		return false, fmt.Sprintf("expected list of length: %d but found list of length: %d", len(b), len(a))
	}

	for i, value := range a {
		switch v := value.(type) {
		case string:
			expected, valid := b[i].(string)
			if !valid {
				return false, fmt.Sprintf("type mismatch")
			}
			if expected != v {
				return false, fmt.Sprintf("expected value: %s but found: %s", expected, v)
			}
		case float64:
			expected, valid := b[i].(float64)
			if !valid {
				return false, fmt.Sprintf("type mismatch")
			}
			if expected != v {
				return false, fmt.Sprintf("expected value: %s but found: %s", expected, v)
			}
		case bool:
			expected, valid := b[i].(bool)
			if !valid {
				return false, fmt.Sprintf("type mismatch")
			}
			if expected != v {
				return false, fmt.Sprintf("expected value: %s but found: %s", expected, v)
			}
		case []interface{}:
			expected, valid := b[i].([]interface{})
			if !valid {
				return false, fmt.Sprintf("type mismatch")
			}
			valid, msg := matchLists(v, expected)
			if !valid {
				return false, msg
			}
		case map[string]interface{}:
			expected, valid := b[i].(map[string]interface{})
			if !valid {
				return false, fmt.Sprintf("type mismatch")
			}
			valid, msg := matchMaps(v, expected)
			if !valid {
				return false, msg
			}
		default:
			return false, "unexpected type found!!!"
		}
	}

	return true, ""
}
func matchMaps(a map[string]interface{}, b map[string]interface{}) (bool, string) {
	if len(a) != len(b) {
		return false, fmt.Sprintf("expected list of length: %d but found list of length: %d", len(b), len(a))
	}

	for key, value := range a {
		switch v := value.(type) {
		case string:
			expected, valid := b[key].(string)
			if !valid {
				return false, fmt.Sprintf("type mismatch")
			}
			if expected != v {
				return false, fmt.Sprintf("expected value: %s but found: %s", expected, v)
			}
		case float64:
			expected, valid := b[key].(float64)
			if !valid {
				return false, fmt.Sprintf("type mismatch")
			}
			if expected != v {
				return false, fmt.Sprintf("expected value: %s but found: %s", expected, v)
			}
		case bool:
			expected, valid := b[key].(bool)
			if !valid {
				return false, fmt.Sprintf("type mismatch")
			}
			if expected != v {
				return false, fmt.Sprintf("expected value: %s but found: %s", expected, v)
			}
		case []interface{}:
			expected, valid := b[key].([]interface{})
			if !valid {
				return false, fmt.Sprintf("type mismatch")
			}
			valid, msg := matchLists(v, expected)
			if !valid {
				return false, msg
			}
		case map[string]interface{}:
			expected, valid := b[key].(map[string]interface{})
			if !valid {
				return false, fmt.Sprintf("type mismatch")
			}
			valid, msg := matchMaps(v, expected)
			if !valid {
				return false, msg
			}
		default:
			return false, "unexpected type found!!!"
		}
	}

	return true, ""
}

func (match *Match_Constraint) Constrain(node models.Node) models.MatchStatus {
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
		// Message:        fmt.Sprintf("matched field: %v with value: %s", match.Field, match.Type),
		Failed:         false,
		Failed_at_node: &node,
	}
}
