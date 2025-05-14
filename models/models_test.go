package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexTraversal(t *testing.T) {
	test_traversal := IndexTraversal(0)

	assert.Equal(t, true, test_traversal.By_index)
	assert.Equal(t, false, test_traversal.By_field)
}

func TestFieldTraversal(t *testing.T) {
	test_traversal := FieldTraversal("test")

	assert.Equal(t, true, test_traversal.By_field)
	assert.Equal(t, false, test_traversal.By_index)
}
