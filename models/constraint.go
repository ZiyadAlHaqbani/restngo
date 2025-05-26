package models

import (
	"fmt"
)

// constrain nodes, where at the end of each step, the response must satisfy all related constraints
type Constraint interface {
	Constrain(node Node) MatchStatus
	ToString() string
}

type MatchStatus struct {
	Failed         bool
	Message        string
	Failed_at_node *Node //	not used so far, potential for future error handling

	//	used to allow '_Store' constraints to store values in global 'Storage' context
	MatchedValue interface{}
	ValueType    MatchType
}

func (match *MatchStatus) ToString() string {
	if !match.Failed {
		temp := ""
		if match.MatchedValue != nil {
			temp += fmt.Sprintf("found Value '%+v' with expected type '%s'", match.MatchedValue, match.ValueType)
		}
		return temp
	}
	temp := ""
	temp = fmt.Sprintf("%sreason: %+v\n", temp, match.Message)

	return temp
}
