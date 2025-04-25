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

func traverse(field string, obj map[string]interface{}) interface{} {

	var traversals []models.Traversal

	traversals, found := parseJSONPath(field)
	if found {
		return nil
	}

	fmt.Printf("Traversals: %+v\n\n", traversals)
	log.Printf("WARNING: detected a match with empty field, are you sure you didn't forget to specify a field for a match operation?")
	log.Printf("WARNING: since no field was specified, following object will be returned: %.100s...", fmt.Sprintf("%+v", obj)) //	solution to golang quirk with '%+v'
	return obj
}

func parseJSONPath(field string) (traversals []models.Traversal, found bool) {

	if len(field) == 0 {
		found = false
		return
	}

	i := 0

	switch field[0] {
	case '.':
		log.Fatalf("ERROR: constraint fields can't start with '.', it must either start with field or index.")
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
			log.Fatal("\t" + strings.Repeat(" ", i-1) + "|")
		}
		i++
		index_str := field[start : i-1]
		index, err := strconv.Atoi(index_str)
		if err != nil {
			log.Fatalf("ERROR: %+v, for string: %s", err, index_str)
		}
		traversals = append(traversals, models.IndexTraversal(index))
	default:
		//	parsing field names
		start := i
		for ; i < len(field) && field[i] != '[' && field[i] != '.'; i++ {
		}

		field_str := field[start:i]
		traversals = append(traversals, models.FieldTraversal(field_str))
		i++
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
			// i++

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
				log.Fatal("\t" + strings.Repeat(" ", i-1) + "|")
			}
			i++
			index_str := field[start : i-1]
			index, err := strconv.Atoi(index_str)
			if err != nil {
				log.Fatalf("ERROR: %+v, for string: %s", err, index_str)
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
			log.Fatalf(strings.Repeat(" ", start) + "|")
		}
	}

	found = true
	return
}
