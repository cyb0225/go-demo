package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	server := &http.Server{
		Addr:    ":12345",
		Handler: nil,
	}

	_ = server.ListenAndServe()
}
