package nodes_test

import (
	builder "htestp/builder"
	"htestp/models"
	nodes "htestp/nodes"
	"net/http"
	"testing"
)

type TestDynamicNode struct {
	*nodes.DynamicNode
	*http.Client
}

func TestExecuteDynamicnode(t *testing.T) {
}

func TestSuccessful(t *testing.T) {
	unsuccessfulBuilder := builder.CreateNewBuilder()
	unsuccessfulBuilder.AddDynamicNodeId("1", "https://dummyjson.com/products", models.GET, nil, nil).
		AddMatchConstraint("doesntexist.field", "welp", models.TypeString)
	unsuccessfulBuilder.Run()

	successfulBuilder := builder.CreateNewBuilder()
	successfulBuilder.AddDynamicNodeId("1", "https://dummyjson.com/products", models.GET, nil, nil).
		AddMatchConstraint("products[0].id", 1, models.TypeFloat)
	successfulBuilder.Run()
	t.Run("an unsuccesful constraint should make the successful function return false", func(t *testing.T) {
		if (*unsuccessfulBuilder.Nodes)["1"].Successful() {
			t.Errorf("Expected %t got %t", false, true)
		}
	})
	t.Run("a successful constraint should make succesful function return true", func(t *testing.T) {
		if !(*successfulBuilder.Nodes)["1"].Successful() {
			t.Errorf("Expected %t got %t", true, false)
		}
	})
}
