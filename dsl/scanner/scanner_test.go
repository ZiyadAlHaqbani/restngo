package scanner

import (
	"testing"
)

func TestCorrectScanner(t *testing.T) {

	source :=
		`
StaticNode("string", 45435, POST,
StaticNode("string", 45435, POST,
ExistConstraint()))
	`

	s := CreateScanner(source)
	s.Scan()

}

func TestIncorrectScanner(t *testing.T) {

	source :=
		`
StaticNode(45435GFDHHGF, POST,
StaticNode("string", 45435, POST,
ExistConstraint()))
	`

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected scanner to panic, but it did not")
		} else {
			t.Logf("Recovered from expected panic: %+v", r)
		}
	}()

	s := CreateScanner(source)
	s.Scan()

}
