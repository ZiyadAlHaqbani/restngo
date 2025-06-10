package constraints

import (
	"encoding/json"
	"htestp/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckEquality(t *testing.T) {

}

func TestCheckType(t *testing.T) {

}

func TestMatchLists(t *testing.T) {

}

func TestMatchMaps(t *testing.T) {

}

func TestTraverse(t *testing.T) {

}

func TestPartialTraverse(t *testing.T) {

	JSON, err := loadJSONFromFile("../assets/test_load.json")
	if err != nil {
		assert.Error(t, err, "expected to load json from file")
	}

	assert.Panics(t, func() {
		partialTraverse("test_field", JSON, models.TypeArray, &models.TypedVariable{
			Value: "test_value",
			Type:  "",
		})
	}, "expected call to panic")

	res, _ := partialTraverse_type("name", JSON, models.TypeString)
	assert.NotNil(t, res, "expected call to match with type 'string'")

	res, msg := partialTraverse_value("name", JSON, &models.TypedVariable{
		Value: "Zamokuhle Zulu",
		Type:  models.TypeString,
	})
	assert.NotNil(t, res, "expected call to match with value 'Zamokuhle Zulu', %q", msg)

	JSON, err = loadJSONFromFile("../assets/test_load_2.json")
	if err != nil {
		assert.Error(t, err)
	}
	JSON_part, err := loadJSONFromFile("../assets/test_load_2_users.json")
	if err != nil {
		assert.Error(t, err)
	}

	res, msg = partialTraverse_value("users", JSON, &models.TypedVariable{
		Value: JSON_part,
		Type:  models.TypeArray,
	})
	assert.NotNil(t, res, msg)

}

func TestParseJSONPath(t *testing.T) {

}

func loadJSONFromFile(filepath string) (interface{}, error) {

	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var JSON interface{}
	err = json.Unmarshal(file, &JSON)
	if err != nil {
		return nil, err
	}

	return JSON, nil
}
