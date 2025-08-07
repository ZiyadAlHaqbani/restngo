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

	if len(node.GetNextNodes()) > 1 {
		list = append(list, node)
	}

	for _, child := range node.GetNextNodes() {
		getBranchesHelper(child, list)
	}

}

func GetBranches_(head models.Node) []models.Segment {
	list_of_segments := []models.Segment{}
	getBranchesHelper_(head, head, &list_of_segments)
	return list_of_segments
}

// pass a list pointer instead of a value to fix issues with mutation
func getBranchesHelper_(node models.Node, previous_branch models.Node, list *[]models.Segment) {

	// if this is a leaf node, create a new segment
	if len(node.GetNextNodes()) == 0 {
		*list = append(*list, models.Segment{
			Start: previous_branch,
			End:   node,
		})
		return
	}

	new_branch := previous_branch

	// if this is a branching node, create a new segment
	if len(node.GetNextNodes()) > 1 {
		*list = append(*list, models.Segment{
			Start: previous_branch,
			End:   node,
		})
		new_branch = node
	}

	for _, child := range node.GetNextNodes() {
		getBranchesHelper_(child, new_branch, list)
	}

}
