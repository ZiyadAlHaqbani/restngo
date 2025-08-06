package models

import (
	"fmt"
)

// constrain nodes, where at the end of each step, the response must satisfy all/some related constraints based on the runner configs.
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

			// restrict the size of the matched value string to reduce clutter
			matchedValue_str := fmt.Sprintf("%v", match.MatchedValue)
			if len(matchedValue_str) >= 100 {
				matchedValue_str = fmt.Sprintf("%.100s", matchedValue_str)
				matchedValue_str += "..."
			}

			temp += fmt.Sprintf("found Value '%s' with expected type '%s'", matchedValue_str, match.ValueType)
		} else {
			temp += match.Message
		}
		return temp
	}
	temp := ""
	temp = fmt.Sprintf("%sreason: %+v\n", temp, match.Message)

	return temp
}
