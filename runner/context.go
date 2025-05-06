package runner

import (
	"fmt"
	"htestp/models"
)

var Storage = map[string]models.TypedVariable{}

// if successful returns a pointer to the stored variable, otherwise returns error
// uses automatic type detection, including TypedVariable itself.
func StoreVariable(varname string, value interface{}) (*models.TypedVariable, error) {
	switch value.(type) {
	case bool:
		Storage[varname] = models.TypedVariable{
			Value: value,
			Type:  models.TypeBool,
		}
		return &Storage[varname], nil
	case float64:
		Storage[varname] = models.TypedVariable{
			Value: value,
			Type:  models.TypeFloat,
		}
		return &Storage[varname], nil
	case string:
		Storage[varname] = models.TypedVariable{
			Value: value,
			Type:  models.TypeString,
		}
		return &Storage[varname], nil
	case []interface{}:
		Storage[varname] = models.TypedVariable{
			Value: value,
			Type:  models.TypeArray,
		}
		return &Storage[varname], nil
	case map[string]interface{}:
		Storage[varname] = models.TypedVariable{
			Value: value,
			Type:  models.TypeObject,
		}
		return &Storage[varname], nil
	case models.TypedVariable:
		Storage[varname] = value.(models.TypedVariable)
		return &Storage[varname], nil
	}
	return nil, fmt.Errorf("type %+T isn't a JSON type, and can't be stored in global context.", value)
}

// if found returns a pointer to the stored variable, else returns nil
func GetVariable(key string) *models.TypedVariable {
	val, exist := Storage[key]
	if exist {
		return &val
	}
	return nil
}
