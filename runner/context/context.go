package context

import (
	"fmt"
	"htestp/models"
	"log"
)

var Storage = map[string]models.TypedVariable{}

// if successful returns a pointer to the stored variable, otherwise returns error
// uses automatic type detection, including TypedVariable itself.
func StoreVariable(varname string, value interface{}) (*models.TypedVariable, error) {
	switch value.(type) {
	case bool:
		v := models.TypedVariable{
			Value: value,
			Type:  models.TypeBool,
		}
		Storage[varname] = v
		return &v, nil
	case float64:
		v := models.TypedVariable{
			Value: value,
			Type:  models.TypeFloat,
		}
		Storage[varname] = v
		return &v, nil
	case string:
		v := models.TypedVariable{
			Value: value,
			Type:  models.TypeString,
		}
		Storage[varname] = v
		return &v, nil
	case []interface{}:
		v := models.TypedVariable{
			Value: value,
			Type:  models.TypeArray,
		}
		Storage[varname] = v
		return &v, nil
	case map[string]interface{}:
		v := models.TypedVariable{
			Value: value,
			Type:  models.TypeObject,
		}
		Storage[varname] = v
		return &v, nil
	case models.TypedVariable:
		v := value.(models.TypedVariable)
		Storage[varname] = v
		return &v, nil
	}
	return nil, fmt.Errorf("type %+T isn't a JSON type, and can't be stored in global context", value)
}

// if found returns a pointer to the stored variable, else returns nil
func GetVariable(key string) *models.TypedVariable {
	val, exist := Storage[key]
	if exist {
		return &val
	}
	return nil
}

// removes every entry from the global context
func PurgeContext() {
	log.Print("WARNINIG: you are removing every entry from the global context!")

	length := len(Storage)
	for key, _ := range Storage {
		delete(Storage, key)
	}

	log.Printf("WARNING: removed: %d from the global context, now the context is empty.", length)
}
