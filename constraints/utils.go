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

func partialTraverse_type(subfield string, obj interface{}, expected_type models.MatchType) (interface{}, string) {
	return partialTraverse(subfield, obj, expected_type, nil)
}

func partialTraverse_value(subfield string, obj interface{}, expected_value *models.TypedVariable) (interface{}, string) {
	return partialTraverse(subfield, obj, models.NOTYPE, expected_value)
}

// this function is the general case for the partial traverse, this is why it contains two extra paramaters, expected_type and expected_value
func partialTraverse(subfield string, obj interface{}, expected_type models.MatchType, expected_value *models.TypedVariable) (interface{}, string) {
	//TODO: refactor to fix ugly code
	if expected_type != models.NOTYPE && expected_value != nil {
		log.Panicf("ERROR: can't traverse for both type and value in partialTraverse")
	}

	traversals, _ := parseJSONPath(subfield)

	temp_obj := obj

	// check for initial shallow match, if found then check for validity of match.
	temp_obj, err_message := _traverse(traversals, temp_obj)
	if temp_obj != nil {
		if _partialTraverseCheck(obj, expected_type, expected_value) {
			return temp_obj, err_message
		}
	}

	//	if not found in the inital match, then perform a BFS match to find the required type/value
	//	first fill the traversal queue
	if temp_obj == nil {
		objects_queue := models.NewQueue[interface{}]()

		if unwrapped_obj, valid := obj.(map[string]interface{}); valid {
			for _, value := range unwrapped_obj {
				if v, valid := value.(map[string]interface{}); valid {
					objects_queue.Enqueue(v)
				} else if v, valid := value.([]interface{}); valid {
					objects_queue.Enqueue(v)
				}
			}
		} else if unwrapped_array, valid := obj.([]interface{}); valid {
			for _, value := range unwrapped_array {
				if v, valid := value.(map[string]interface{}); valid {
					objects_queue.Enqueue(v)
				} else if v, valid := value.([]interface{}); valid {
					objects_queue.Enqueue(v)
				}
			}
		} else {
			return nil, fmt.Sprintf("couldn't traverse object: %v as its not a JSON object or an array", obj)
		}

		for objects_queue.Len() > 0 {
			temp_obj = objects_queue.Dequeue()
			if _partialTraverseCheck(temp_obj, expected_type, expected_value) {
				return temp_obj, err_message
			}

			if obj, err_message := _traverse(traversals, temp_obj); obj != nil {
				if _partialTraverseCheck(obj, expected_type, expected_value) {
					return obj, err_message
				}
			}

			if unwrapped_obj, valid := temp_obj.(map[string]interface{}); valid {
				for _, value := range unwrapped_obj {
					if v, valid := value.(map[string]interface{}); valid {
						objects_queue.Enqueue(v)
					} else if v, valid := value.([]interface{}); valid {
						objects_queue.Enqueue(v)
					}
				}
			} else if unwrapped_array, valid := temp_obj.([]interface{}); valid {
				for _, value := range unwrapped_array {
					if v, valid := value.(map[string]interface{}); valid {
						objects_queue.Enqueue(v)
					} else if v, valid := value.([]interface{}); valid {
						objects_queue.Enqueue(v)
					}
				}
			}

		}

		return nil, "TODO: error message when no partial field match found"
	}

	return temp_obj, err_message
}

// check if the root json object matches the given criteria, return true if yes and false if not.
func _partialTraverseCheck(obj interface{}, expected_type models.MatchType, expected_value *models.TypedVariable) bool {

	if expected_type != models.NOTYPE {
		status := checkType(obj, expected_type, "")
		return !status.Failed
	} else if expected_value != nil {
		return !checkEquality_typed(obj, expected_value, "").Failed
	}

	return true
}

func traverse(field string, obj interface{}) (interface{}, string) {

	var traversals []models.Traversal

	traversals, _ = parseJSONPath(field)

	return _traverse(traversals, obj)

}

func _traverse(traversals []models.Traversal, obj interface{}) (interface{}, string) {

	if len(traversals) == 0 {
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

func checkEquality_typed(value interface{}, expected *models.TypedVariable, field string) models.MatchStatus {
	return checkEquality(value, expected.Value, field)
}

func checkEquality(value interface{}, expected interface{}, field string) models.MatchStatus {

	if expected == nil {
		return models.MatchStatus{
			Failed:  true,
			Message: fmt.Sprintf("field: %s expected value is nil", field),
		}
	}

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
		if !expected_valid {
			return models.MatchStatus{
				Failed:  true,
				Message: "",
			}
		}

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
		if !expected_valid {
			return models.MatchStatus{
				Failed:  true,
				Message: "",
			}
		}

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
