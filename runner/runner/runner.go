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

func GetBranches(head models.Node) []models.Node {
	list_of_branches := []models.Node{}
	list_of_branches = append(list_of_branches, head)
	getBranchesHelper(head, list_of_branches)
	return list_of_branches
}

func getBranchesHelper(node models.Node, list []models.Node) {

	if len(node.GetNextNodes()) == 0 {
		return
	}

	if len(node.GetNextNodes()) == 1 {
		getBranchesHelper(node.GetNextNodes()[0], list)
	} else {
		for _, next := range node.GetNextNodes() {
			list = append(list, next)
			getBranchesHelper(next, list)
		}
	}

}
