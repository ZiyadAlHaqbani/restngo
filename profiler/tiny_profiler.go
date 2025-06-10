package profilers

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Scope struct {
	id    string
	start int64
	end   int64
	depth uint8
}

type ChromeTraceEvent struct {
	Name string `json:"name"`
	Cat  string `json:"cat"`
	Ph   string `json:"ph"`
	Ts   int64  `json:"ts"`
	Dur  int64  `json:"dur"`
	Pid  int    `json:"pid"`
	Tid  int    `json:"tid"`
}

var Scopes []Scope
var depth uint8 = 0

func ProfileScope(id string) func() {

	scope := Scope{
		id:    id,
		start: time.Now().UnixMicro(),
		depth: depth,
	}
	depth++
	return func() {
		scope.end = time.Now().UnixMicro()
		Scopes = append(Scopes, scope)
		depth--
	}
}

func dumpScopesToChromeTrace(filename string, scopes []Scope) error {
	events := []ChromeTraceEvent{}

	for _, scope := range scopes {
		events = append(events, ChromeTraceEvent{
			Name: scope.id,
			Cat:  "function",
			Ph:   "X", // "Complete" event
			Ts:   scope.start,
			Dur:  (scope.end - scope.start),
			Pid:  1,                // Fake process ID
			Tid:  int(scope.depth), // Use depth as thread ID (fake)
		})
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	trace := map[string]interface{}{
		"traceEvents": events,
	}

	enc := json.NewEncoder(f)
	return enc.Encode(trace)
}

func DumpTrace(filename string) {

	err := dumpScopesToChromeTrace(filename, Scopes)
	if err != nil {
		log.Panicf("ERROR: %+v", err)
	}
}
