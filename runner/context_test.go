package runner

import (
	"htestp/models"
	"testing"
)

func TestStoreVariableAndRetrieve(t *testing.T) {

	//	Storing test
	correctTypedVariable := models.TypedVariable{
		Value: "test_string",
		Type:  models.TypeString,
	}
	_, typed_err := StoreVariable("typedvariable", correctTypedVariable)
	if typed_err != nil {
		t.Errorf("Expected nil, got %+v", typed_err)
	}

	correctStringVariable := "test_string"
	_, string_err := StoreVariable("string", correctStringVariable)
	if string_err != nil {
		t.Errorf("Expected nil, got %+v", string_err)
	}

	var correctFloatVariable float64 = 10.5
	_, float_err := StoreVariable("float", correctFloatVariable)
	if float_err != nil {
		t.Errorf("Expected nil, got %+v", float_err)
	}

	correctBoolVariable := false
	_, bool_err := StoreVariable("bool", correctBoolVariable)
	if bool_err != nil {
		t.Errorf("Expected nil, got %+v", bool_err)
	}

	correctArrayVariable := []interface{}{"a", 10, true}
	_, array_err := StoreVariable("array", correctArrayVariable)
	if array_err != nil {
		t.Errorf("Expected nil, got %+v", array_err)
	}

	correctObjectVariable := map[string]interface{}{"string": "test", "float": 10.5, "bool": true}
	_, object_err := StoreVariable("object", correctObjectVariable)
	if object_err != nil {
		t.Errorf("Expected nil, got %+v", object_err)
	}

	//	Retrieval test
	if GetVariable("typedvariable") == nil {
		t.Errorf("Expected value, got nil")
	}

	if GetVariable("string") == nil {
		t.Errorf("Expected value, got nil")
	}

	if GetVariable("float") == nil {
		t.Errorf("Expected value, got nil")
	}

	if GetVariable("bool") == nil {
		t.Errorf("Expected value, got nil")
	}

	if GetVariable("array") == nil {
		t.Errorf("Expected value, got nil")
	}

	if GetVariable("object") == nil {
		t.Errorf("Expected value, got nil")
	}

}
