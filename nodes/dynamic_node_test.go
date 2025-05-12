package nodes_test

import (
	builder "htestp/builder"
	"htestp/constraints"
	httphandler "htestp/http_handler"
	"htestp/models"
	nodes "htestp/nodes"
	"htestp/runner/context"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestDynamicNode struct {
	*nodes.DynamicNode
	*http.Client
}

func TestExecuteDynamicnode(t *testing.T) {
	t.Run("test correct query with expected response", func(t *testing.T) {
		testNode := TestDynamicNode{
			DynamicNode: &nodes.DynamicNode{
				InnerNode: nodes.StaticNode{
					Request: httphandler.Request{
						Url:    "https://dummyjson.com/products",
						Method: string(models.GET),
					},
				},
			},
		}
		response, err := testNode.Execute(http.DefaultClient)
		assert.Nil(t, err)
		assert.Equal(t, 200, response.Status)
	})
	t.Run("test incorrect query with expected response", func(t *testing.T) {
		testNode := TestDynamicNode{
			DynamicNode: &nodes.DynamicNode{
				InnerNode: nodes.StaticNode{
					Request: httphandler.Request{
						Url:    "https://dummyjson.com/doesntexist",
						Method: string(models.GET),
					},
				},
			},
		}
		response, err := testNode.Execute(http.DefaultClient)
		assert.Nil(t, err)
		assert.Equal(t, 404, response.Status)
	})
	t.Run("test use of context variables", func(t *testing.T) {
		context.StoreVariable("q", "phone")
		testNode := TestDynamicNode{
			DynamicNode: &nodes.DynamicNode{
				InnerNode: nodes.StaticNode{
					Request: httphandler.Request{
						Url:    "https://dummyjson.com/products/search",
						Method: string(models.GET),
					},
				},
				QueryBuilderFunc: func(storage *map[string]models.TypedVariable) url.Values {
					searchProduct := (*storage)["q"].Value
					return url.Values{"q": {searchProduct.(string)}}
				},
			},
		}
		response, err := testNode.Execute(http.DefaultClient)
		assert.Nil(t, err)
		assert.Equal(t, 200, response.Status)
		assert.Equal(t, 101.0, response.Body.(map[string]interface{})["products"].([]interface{})[0].(map[string]interface{})["id"])
	})
}
func TestCheck(t *testing.T) {
	t.Run("Check returns the result of InnerNode's Check", func(t *testing.T) {
		testNode := TestDynamicNode{
			DynamicNode: &nodes.DynamicNode{},
			Client:      &http.Client{},
		}
		testNode.InnerNode = nodes.StaticNode{
			Request: httphandler.Request{
				Url:    "https://dummyjson.com/products",
				Method: string(models.GET),
			},
		}
		testNode.InnerNode.Execute(testNode.Client)
		assert.Equal(t, testNode.InnerNode.Check(), testNode.Check())
	})
	t.Run("Check returns false if InnerNode's Check returns false", func(t *testing.T) {
		testNode := TestDynamicNode{
			DynamicNode: &nodes.DynamicNode{},
			Client:      &http.Client{},
		}
		testNode.InnerNode = nodes.StaticNode{
			Request: httphandler.Request{
				Url:    "https://dummyjson.com/products",
				Method: string(models.GET),
			},
		}
		testNode.AddConstraint(&constraints.Match_Constraint{
			Field:    "products[0].id",
			Type:     models.TypeFloat,
			Expected: 0.0,
		})
		testNode.InnerNode.Execute(testNode.Client)
		assert.Equal(t, false, testNode.Check())
	})
	t.Run("Check returns true if InnerNode's Check returns true", func(t *testing.T) {
		testNode := TestDynamicNode{
			DynamicNode: &nodes.DynamicNode{},
			Client:      &http.Client{},
		}
		testNode.InnerNode = nodes.StaticNode{
			Request: httphandler.Request{
				Url:    "https://dummyjson.com/products",
				Method: string(models.GET),
			},
		}
		testNode.AddConstraint(&constraints.Match_Constraint{
			Field:    "products[0].id",
			Type:     models.TypeFloat,
			Expected: 1.0,
		})
		testNode.InnerNode.Execute(testNode.Client)
		assert.Equal(t, true, testNode.Check())
	})
}
func TestGetResp(t *testing.T) {
	t.Run("GetResp returns the response from InnerNode", func(t *testing.T) {
		testNode := TestDynamicNode{
			DynamicNode: &nodes.DynamicNode{},
			Client:      &http.Client{},
		}
		testNode.InnerNode = nodes.StaticNode{
			Request: httphandler.Request{
				Url:    "https://dummyjson.com/products",
				Method: string(models.GET),
			},
		}
		testNode.InnerNode.Execute(testNode.Client)
		assert.Equal(t, testNode.InnerNode.GetResp(), testNode.GetResp())
	})
}

func TestAddConstraint(t *testing.T) {
	t.Run("AddConstraint adds to InnerNode's constraints", func(t *testing.T) {
		testNode := nodes.DynamicNode{}
		constraint := constraints.Match_Constraint{
			Field:    "title",
			Type:     models.TypeString,
			Expected: "iPhone 9",
		}
		testNode.AddConstraint(&constraint)
		assert.Equal(t, 1, len(testNode.InnerNode.Constraints))
		assert.Equal(t, &constraint, testNode.InnerNode.Constraints[0])
	})
	t.Run("AddConstraint adds nil to InnerNode's constraints", func(t *testing.T) {
		testNode := nodes.DynamicNode{}
		testNode.AddConstraint(nil)
		assert.Equal(t, 1, len(testNode.InnerNode.Constraints))
		assert.Nil(t, testNode.InnerNode.Constraints[0])
	})
}

func TestGetNextNodes(t *testing.T) {
	t.Run("Added multiple nodes should be returned exactly when using GetNextNodes", func(t *testing.T) {
		StaticNodeOne := nodes.StaticNode{
			Request: httphandler.Request{
				Url:    "google.com",
				Method: string(models.GET),
			},
		}
		StaticNodeTwo := nodes.StaticNode{
			Request: httphandler.Request{
				Url:    "youtube.com",
				Method: string(models.GET),
			},
		}

		ExpectedNodes := []models.Node{&StaticNodeOne, &StaticNodeTwo}

		testBuilder := builder.CreateNewBuilder()
		testBuilder.AddDynamicNodeId("1", "trendyol.com", models.GET, nil, nil)
		testBuilder.AddStaticNodeBranch("google.com", models.GET, nil)
		testBuilder.AddStaticNodeBranch("youtube.com", models.GET, nil)

		assert.Equal(t, ExpectedNodes, (*testBuilder.Nodes)["1"].GetNextNodes())
	})
	// TODO: dont add nodes and check
}
func TestAddNode(t *testing.T) {
	t.Run("added dynamic node to the end by AddNode should match last node", func(t *testing.T) {
		testBuilder := builder.CreateNewBuilder()
		testBuilder.AddDynamicNodeId("1", "", models.GET, nil, nil)
		testNode := nodes.DynamicNode{
			InnerNode: nodes.StaticNode{
				Request: httphandler.Request{
					Url:    "https://dummyjson.com/products",
					Method: string(models.GET),
				},
			},
		}
		(*testBuilder.Nodes)["1"].AddNode(&testNode)

		assert.Equal(t,
			(*testBuilder.Nodes)["1"].(*nodes.DynamicNode).Next[0], &testNode)
	})
	t.Run("added nil to the end by AddNode should match nil", func(t *testing.T) {
		testBuilder := builder.CreateNewBuilder()
		testBuilder.AddDynamicNodeId("1", "", models.GET, nil, nil)

		(*testBuilder.Nodes)["1"].AddNode(nil)

		assert.Equal(t, (*testBuilder.Nodes)["1"].(*nodes.DynamicNode).Next[0], nil)
	})
}
func TestSuccessful(t *testing.T) {
	// TODO: Refactor to use testify
	t.Run("an unexecuted dynamic node should make the successful function return true", func(t *testing.T) {
		unExecutedBuilder := builder.CreateNewBuilder()
		unExecutedBuilder.AddDynamicNodeId("1", "https://dummyjson.com/products", models.GET, nil, nil).
			AddMatchConstraint("products[0].id", 1, models.TypeFloat)
		if !(*unExecutedBuilder.Nodes)["1"].Successful() {
			t.Errorf("Expected %t got %t", true, false)
		}
	})
	t.Run("an unsuccesful constraint should make the successful function return false", func(t *testing.T) {
		unsuccessfulBuilder := builder.CreateNewBuilder()
		unsuccessfulBuilder.AddDynamicNodeId("1", "https://dummyjson.com/products", models.GET, nil, nil).
			AddMatchConstraint("doesntexist.field", "welp", models.TypeString)
		unsuccessfulBuilder.Run()
		if (*unsuccessfulBuilder.Nodes)["1"].Successful() {
			t.Errorf("Expected %t got %t", false, true)
		}
	})
	t.Run("a successful constraint should make succesful function return true", func(t *testing.T) {
		successfulBuilder := builder.CreateNewBuilder()
		successfulBuilder.AddDynamicNodeId("1", "https://dummyjson.com/products", models.GET, nil, nil).
			AddMatchConstraint("products[0].id", 1.0, models.TypeFloat)
		successfulBuilder.Run()
		if !(*successfulBuilder.Nodes)["1"].Successful() {
			t.Errorf("Expected %t got %t", true, false)
		}
	})
}
