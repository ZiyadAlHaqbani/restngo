package constraints

import "fmt"

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
