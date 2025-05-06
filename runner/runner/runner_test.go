package runner

import (
	"htestp/nodes"
	"net/http"
	"testing"
)

func TestUnsuccessfulRunner(t *testing.T) {

	unsuccessfulNode := &nodes.MockNode{
		ShouldSucceed: false,
	}

	RunHelper(http.DefaultClient, unsuccessfulNode)

	if unsuccessfulNode.Successful() {
		t.Errorf("expected false, got true")
	}
}

func TestSuccessfulRunner(t *testing.T) {

	successfulNode := &nodes.MockNode{
		ShouldSucceed: true,
	}

	RunHelper(http.DefaultClient, successfulNode)

	if !successfulNode.Successful() {
		t.Errorf("expected true, got false")
	}
}
