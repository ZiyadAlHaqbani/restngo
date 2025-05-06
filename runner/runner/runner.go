package runner

import (
	"htestp/models"
	"log"
	"net/http"
)

func RunHelper(client *http.Client, node models.Node) bool {

	if node == nil {
		return true
	}

	_, err := node.Execute(client)
	if err != nil {
		log.Print(err)
		return false
	}
	if !node.Check() {
		return false
	}

	if len(node.GetNextNodes()) == 1 {
		return RunHelper(client, node.GetNextNodes()[0])
	}

	// branches will still run even if a node in the level fails.
	success := true

	for _, nextNode := range node.GetNextNodes() {
		successful := RunHelper(client, nextNode)
		if !successful {
			success = false
		}
	}

	return success
}
