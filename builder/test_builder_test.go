package builder_test

import (
	"htestp/builder"
	httphandler "htestp/http_handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddChildDynamicNode(t *testing.T) {
	t.Run("Test adding a child dynamic node to a static node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddStaticNodeId("node1", "http://example.com", "GET", nil)
		builder.AddChildDynamicNode("node2", "node1", "http://example.com", "GET", nil, nil)
		assert.Equal(t, (*builder.Nodes)["node1"].GetNextNodes()[0].GetID(), "node2")
	})
	t.Run("Test adding a child dynamic node to a dynamic node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddDynamicNodeId("node1", "http://example.com", "GET", nil, nil)
		builder.AddChildDynamicNode("node2", "node1", "http://example.com", "GET", nil, nil)
		assert.Equal(t, (*builder.Nodes)["node1"].GetNextNodes()[0].GetID(), "node2")
	})
	t.Run("Test adding a child dynamic node to a non-existent node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		assert.Panics(t, func() {
			builder.AddChildDynamicNode("node2", "node1", "http://example.com", "GET", nil, nil)
		})
	})
	t.Run("Test adding a child dynamic node with a duplicate ID", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddDynamicNodeId("node1", "http://example.com", "GET", nil, nil)
		assert.Panics(t, func() {
			builder.AddChildDynamicNode("node1", "node1", "http://example.com", "GET", nil, nil)
		})
	})
	t.Run("Test adding multiple child dynamic nodes to a static node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddStaticNodeId("node1", "http://example.com", "GET", nil)
		builder.AddChildDynamicNode("node2", "node1", "http://example.com", "GET", nil, nil)
		builder.AddChildDynamicNode("node3", "node1", "http://example.com", "GET", nil, nil)
		assert.Equal(t, len((*builder.Nodes)["node1"].GetNextNodes()), 2)
		assert.Equal(t, (*builder.Nodes)["node1"].GetNextNodes()[0].GetID(), "node2")
		assert.Equal(t, (*builder.Nodes)["node1"].GetNextNodes()[1].GetID(), "node3")
	})
	t.Run("Test adding multiple child dynamic nodes to a dynamic node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddDynamicNodeId("node1", "http://example.com", "GET", nil, nil)
		builder.AddChildDynamicNode("node2", "node1", "http://example.com", "GET", nil, nil)
		builder.AddChildDynamicNode("node3", "node1", "http://example.com", "GET", nil, nil)
		assert.Equal(t, len((*builder.Nodes)["node1"].GetNextNodes()), 2)
	})
}

func TestAddChildDynamicNodeRaw(t *testing.T) {
	t.Run("Test adding a child dynamic node to a static node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddStaticNodeId("node1", "http://example.com", "GET", nil)
		builder.AddChildDynamicNodeRaw("node2", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"}, nil, nil)
		assert.Equal(t, (*builder.Nodes)["node1"].GetNextNodes()[0].GetID(), "node2")
	})
	t.Run("Test adding a child dynamic node to a dynamic node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddDynamicNodeId("node1", "http://example.com", "GET", nil, nil)
		builder.AddChildDynamicNodeRaw("node2", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"}, nil, nil)
		assert.Equal(t, (*builder.Nodes)["node1"].GetNextNodes()[0].GetID(), "node2")
	})
	t.Run("Test adding a child dynamic node to a non-existent node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		assert.Panics(t, func() {
			builder.AddChildDynamicNodeRaw("node2", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"}, nil, nil)
		})
	})
	t.Run("Test adding a child dynamic node with a duplicate ID", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddDynamicNodeId("node1", "http://example.com", "GET", nil, nil)
		assert.Panics(t, func() {
			builder.AddChildDynamicNodeRaw("node1", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"}, nil, nil)
		})
	})
	t.Run("Test adding multiple child dynamic nodes to a static node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddStaticNodeId("node1", "http://example.com", "GET", nil)
		builder.AddChildDynamicNodeRaw("node2", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"}, nil, nil)
		builder.AddChildDynamicNodeRaw("node3", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"}, nil, nil)
		assert.Equal(t, len((*builder.Nodes)["node1"].GetNextNodes()), 2)
	})
	t.Run("Test adding multiple child dynamic nodes to a dynamic node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddDynamicNodeId("node1", "http://example.com", "GET", nil, nil)
		builder.AddChildDynamicNodeRaw("node2", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"}, nil, nil)
		builder.AddChildDynamicNodeRaw("node3", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"}, nil, nil)
		assert.Equal(t, len((*builder.Nodes)["node1"].GetNextNodes()), 2)
	})
}

func TestAddChildStaticNode(t *testing.T) {
	t.Run("Test adding a child static node to a static node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddStaticNodeId("node1", "http://example.com", "GET", nil)
		builder.AddChildStaticNode("node2", "node1", "http://example.com", "GET", nil)
		assert.Equal(t, (*builder.Nodes)["node1"].GetNextNodes()[0].GetID(), "node2")
	})
	t.Run("Test adding a child static node to a dynamic node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddDynamicNodeId("node1", "http://example.com", "GET", nil, nil)
		builder.AddChildStaticNode("node2", "node1", "http://example.com", "GET", nil)
		assert.Equal(t, (*builder.Nodes)["node1"].GetNextNodes()[0].GetID(), "node2")
	})
	t.Run("Test adding a child static node to a non-existent node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		assert.Panics(t, func() {
			builder.AddChildStaticNode("node2", "node1", "http://example.com", "GET", nil)
		})
	})
	t.Run("Test adding a child static node with a duplicate ID", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddStaticNodeId("node1", "http://example.com", "GET", nil)
		assert.Panics(t, func() {
			builder.AddChildStaticNode("node1", "node1", "http://example.com", "GET", nil)
		})
	})
	t.Run("Test adding multiple child static nodes to a static node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddStaticNodeId("node1", "http://example.com", "GET", nil)
		builder.AddChildStaticNode("node2", "node1", "http://example.com", "GET", nil)
		builder.AddChildStaticNode("node3", "node1", "http://example.com", "GET", nil)
		assert.Equal(t, len((*builder.Nodes)["node1"].GetNextNodes()), 2)
	})
	t.Run("Test adding multiple child static nodes to a dynamic node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddDynamicNodeId("node1", "http://example.com", "GET", nil, nil)
		builder.AddChildStaticNode("node2", "node1", "http://example.com", "GET", nil)
		builder.AddChildStaticNode("node3", "node1", "http://example.com", "GET", nil)
		assert.Equal(t, len((*builder.Nodes)["node1"].GetNextNodes()), 2)
	})
}

func TestAddChildStaticNodeRaw(t *testing.T) {
	t.Run("Test adding a child static node to a static node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddStaticNodeId("node1", "http://example.com", "GET", nil)
		builder.AddChildStaticNodeRaw("node2", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"})
		assert.Equal(t, (*builder.Nodes)["node1"].GetNextNodes()[0].GetID(), "node2")
	})
	t.Run("Test adding a child static node to a dynamic node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddDynamicNodeId("node1", "http://example.com", "GET", nil, nil)
		builder.AddChildStaticNodeRaw("node2", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"})
		assert.Equal(t, (*builder.Nodes)["node1"].GetNextNodes()[0].GetID(), "node2")
	})
	t.Run("Test adding a child static node to a non-existent node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		assert.Panics(t, func() {
			builder.AddChildStaticNodeRaw("node2", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"})
		})
	})
	t.Run("Test adding a child static node with a duplicate ID", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddStaticNodeId("node1", "http://example.com", "GET", nil)
		assert.Panics(t, func() {
			builder.AddChildStaticNodeRaw("node1", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"})
		})
	})
	t.Run("Test adding multiple child static nodes to a static node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddStaticNodeId("node1", "http://example.com", "GET", nil)
		builder.AddChildStaticNodeRaw("node2", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"})
		builder.AddChildStaticNodeRaw("node3", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"})
		assert.Equal(t, len((*builder.Nodes)["node1"].GetNextNodes()), 2)
	})
	t.Run("Test adding multiple child static nodes to a dynamic node", func(t *testing.T) {
		builder := builder.CreateNewBuilder()
		builder.AddDynamicNodeId("node1", "http://example.com", "GET", nil, nil)
		builder.AddChildStaticNodeRaw("node2", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"})
		builder.AddChildStaticNodeRaw("node3", "node1", httphandler.Request{Url: "http://example.com", Method: "GET"})
		assert.Equal(t, len((*builder.Nodes)["node1"].GetNextNodes()), 2)
	})
}
