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

	test_json, err := os.ReadFile("../dsl/cache/test_load.json")
	if err != nil {
		assert.Error(t, err, "error reading test_load.json")
	}

	var JSON interface{}
	err = json.Unmarshal(test_json, &JSON)
	if err != nil {
		assert.Error(t, err, "error unmarhsalling test_load.json")
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

	match_string := `
	{
    "name": "Bhupesh Menon","language": "Hindi","id": "0CEPNRDV98KT3ORP",
    "bio": "Maecenas tempus neque ut porttitor malesuada. Phasellus massa ligula, hendrerit eget efficitur eget, tincidunt in ligula. Quisque mauris ligula, efficitur porttitor sodales ac, lacinia non ex. Maecenas quis nisi nunc.",
    "version": 2.69}`
	var match_json map[string]interface{}

	err = json.Unmarshal([]byte(match_string), &match_json)
	if err != nil {
		assert.Error(t, err, "error unmarhsalling test_load.json")
	}

}

func TestParseJSONPath(t *testing.T) {

}
