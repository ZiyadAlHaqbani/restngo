package constraints

import (
	"fmt"
	"htestp/models"
	"log"
	"strconv"
	"strings"
)

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

func traverse(field string, obj interface{}) (interface{}, string) {

	var traversals []models.Traversal

	traversals, found := parseJSONPath(field)
	if !found {
		fmt.Printf("Traversals: %+v\n\n", traversals)
		log.Printf("WARNING: detected a constraint with empty field, are you sure you didn't forget to specify a field for the constraint?")
		log.Printf("WARNING: since no field was specified, following object will be returned: %.100s...", fmt.Sprintf("%+v", obj)) //	solution to golang quirk with '%+v'
		return obj, ""
	}

	temp_obj := obj

	for _, traversal := range traversals {
		if traversal.By_field {
			temp_dict, valid := temp_obj.(map[string]interface{})
			if !valid {
				return nil, fmt.Sprintf("object is type: '%T' and not an object for field: %s", temp_obj, traversal.Field)
			}
			temp_obj, valid = temp_dict[traversal.Field]
			if !valid {
				return nil, fmt.Sprintf("field: %s doesn't exist in object: %.100s", traversal.Field, fmt.Sprintf("%+v", obj))
			}

		} else if traversal.By_index {
			temp_list, valid := temp_obj.([]interface{})
			if !valid {
				return nil, fmt.Sprintf("object is type: '%T' and not an a list for index: %d", temp_obj, traversal.Index)
			}
			if len(temp_list) <= traversal.Index {
				return nil, fmt.Sprintf("index: %d is outside the bounds of list of length: %d", traversal.Index, len(temp_list))
			}
			temp_obj = temp_list[traversal.Index]
		} else {
			log.Panicf("ERROR: malformed traversal: %+v, is neither field nor index.", traversal)
		}
	}

	return temp_obj, ""
}

func parseJSONPath(field string) (traversals []models.Traversal, found bool) {

	if len(field) == 0 {
		found = false
		return
	}

	i := 0

	switch field[0] {
	case '.':
		log.Panicf("ERROR: constraint fields can't start with '.', it must either start with field or index.")
	case '[':
		//	parsing index brackets
		start := 1
		i++
		for ; i < len(field) && field[i] != ']'; i++ {
		}
		if i >= len(field) {
			log.Printf("ERROR: reached end of field without finding a matching ']'.")
			log.Printf("\t%s%c", field[:i-1], field[i-1])
			log.Print("\t" + strings.Repeat(" ", i-1) + "^")
			log.Panic("\t" + strings.Repeat(" ", i-1) + "|")
		}
		i++
		index_str := field[start : i-1]
		index, err := strconv.Atoi(index_str)
		if err != nil {
			log.Panicf("ERROR: %+v, for string: %s", err, index_str)
		}
		traversals = append(traversals, models.IndexTraversal(index))
	default:
		//	parsing field names
		start := i
		for ; i < len(field) && field[i] != '[' && field[i] != '.'; i++ {
		}

		field_str := field[start:i]
		traversals = append(traversals, models.FieldTraversal(field_str))
		// i++
	}

	for i < len(field) {
		switch field[i] {
		case '.':
			//	parsing field names
			i++
			start := i
			for ; i < len(field) && field[i] != '[' && field[i] != '.'; i++ {
			}

			field_str := field[start:i]
			traversals = append(traversals, models.FieldTraversal(field_str))

		case '[':
			//	parsing index brackets
			i++
			start := i
			for ; i < len(field) && field[i] != ']'; i++ {
			}
			if i >= len(field) {
				log.Printf("ERROR: reached end of field without finding a matching ']'.")
				log.Printf("\t%s%c", field[:i-1], field[i-1])
				log.Print("\t" + strings.Repeat(" ", i-1) + "^")
				log.Panic("\t" + strings.Repeat(" ", i-1) + "|")
			}
			i++
			index_str := field[start : i-1]
			index, err := strconv.Atoi(index_str)
			if err != nil {
				log.Panicf("ERROR: %+v, for string: %s", err, index_str)
			}
			traversals = append(traversals, models.IndexTraversal(index))
		default:
			start := i
			log.Printf("ERROR: constraint fields can't start in the middle without a '.' operator before it.")
			if start+1 >= len(field) {
				log.Printf("%s%c", field[:start], field[start])
			} else {
				log.Printf("%s%c%s", field[:start], field[start], field[start+1:])
			}
			log.Printf(strings.Repeat(" ", start) + "^")
			log.Panicf(strings.Repeat(" ", start) + "|")
		}
	}

	found = true
	return
}

func checkType(value interface{}, data_type models.MatchType, field string) models.MatchStatus {

	switch data_type {
	case models.TypeString:
		_, valid := value.(string)
		if !valid {
			return models.MatchStatus{
				Failed:  true,
				Message: fmt.Sprintf("field: %s doesn't match expected type: %s", field, data_type),
			}
		}
	case models.TypeFloat:
		_, valid := value.(float64)
		if !valid {
			return models.MatchStatus{
				Failed:  true,
				Message: fmt.Sprintf("field: %s doesn't match expected type: %s", field, data_type),
			}
		}
	case models.TypeBool:
		_, valid := value.(bool)
		if !valid {
			return models.MatchStatus{
				Failed:  true,
				Message: fmt.Sprintf("field: %s with type: %T doesn't match expected type: %s", field, value, data_type),
			}
		}
	case models.TypeArray:
		_, valid := value.([]interface{})
		if !valid {
			return models.MatchStatus{
				Failed:  true,
				Message: fmt.Sprintf("field: %s with type: %T doesn't match expected type: %s", field, value, data_type),
			}
		}
	case models.TypeObject:
		_, valid := value.(map[string]interface{})
		if !valid {
			return models.MatchStatus{
				Failed:  true,
				Message: fmt.Sprintf("field: %s with type: %T doesn't match expected type: %s", field, value, data_type),
			}
		}

	default:
		panic("ERROR: user assigned type outside of the defined types in 'MatchType'")

	}

	return models.MatchStatus{
		Failed: false,
		// Failed_at_node: &node,
		ValueType:    data_type,
		MatchedValue: value,
	}

}

func checkEquality(value interface{}, expected interface{}, field string) models.MatchStatus {

	switch value.(type) {
	case string:
		_, expected_valid := expected.(string)
		if !(expected_valid && value.(string) == expected.(string)) {
			return models.MatchStatus{
				Failed:  true,
				Message: fmt.Sprintf("field: %s doesn't match expected value: %s", field, expected),
			}
		}

	case float64:
		_, expected_valid := expected.(float64)
		if !(expected_valid && value.(float64) == expected.(float64)) {
			return models.MatchStatus{
				Failed:  true,
				Message: fmt.Sprintf("field: %s doesn't match expected value: %s", field, expected),
			}
		}
	case bool:
		_, expected_valid := expected.(bool)
		if !(expected_valid && value.(bool) == expected.(bool)) {
			return models.MatchStatus{
				Failed:  true,
				Message: fmt.Sprintf("field: %s doesn't match expected value: %s", field, expected),
			}
		}
	case []interface{}:
		_, expected_valid := expected.([]interface{})
		valid, msg := matchLists(value.([]interface{}), expected.([]interface{}))
		if expected_valid {
			return models.MatchStatus{
				Failed:  !valid,
				Message: msg,
			}
		}
		return models.MatchStatus{
			Failed:  true,
			Message: fmt.Sprintf("field: %s doesn't match expected value: %s", field, expected),
		}
	case map[string]interface{}:
		_, expected_valid := expected.(map[string]interface{})
		valid, msg := matchMaps(value.(map[string]interface{}), expected.(map[string]interface{}))
		if expected_valid {
			return models.MatchStatus{
				Failed:  !valid,
				Message: msg,
			}
		}
		return models.MatchStatus{
			Failed:  true,
			Message: fmt.Sprintf("field: %s doesn't match expected value: %s", field, expected),
		}

	default:
		panic("ERROR: user assigned type outside of the defined types in 'MatchType'")

	}

	return models.MatchStatus{
		Failed: false,
		// Don't set MatchedType because it's not known
		MatchedValue: value,
	}

}
