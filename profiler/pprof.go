package profilers

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func StartPProf() func() {
	go func() {
		log.Print("pprof listening at http://localhost:6060/debug/pprof/, waiting for 5 seconds so pprof could connect.")
		http.ListenAndServe("localhost:6060", nil)
	}()
	time.Sleep(time.Second * 5)
	return func() {
		select {}
	}
}
