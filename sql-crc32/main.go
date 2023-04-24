package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/pingcap/test-infra/tools/sql-crc32/cmd"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	cmd.Execute()
}
