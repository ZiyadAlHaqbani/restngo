package models

import "fmt"

// constrain nodes, where at the end of each step, the response must satisfy all related constraints
type Constraint interface {
	Constrain(node Node) MatchStatus
}

type MatchStatus struct {
	Failed         bool
	Message        string
	Failed_at_node *Node
}

func (match *MatchStatus) ToString() string {
	if !match.Failed {
		return match.Message
	}
	temp := ""
	temp = fmt.Sprintf("%sreason: %+v\n", temp, match.Message)

	return temp
}
