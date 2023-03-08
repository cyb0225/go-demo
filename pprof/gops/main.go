package main

import (
	"github.com/google/gops/agent"
	"log"
	"runtime"
	"time"
)

func main() {
	if err := agent.Listen(agent.Options{
		Addr:                   "127.0.0.1:12345",
		ConfigDir:              "./pprof",
		ShutdownCleanup:        true,
		ReuseSocketAddrAndPort: true,
	}); err != nil {
		log.Fatal(err)
	}
	log.Println("程序启动....")

	_ = make([]int, 1000, 1000)
	runtime.GC()

	_ = make([]int, 1000, 2000)
	runtime.GC()

	time.Sleep(time.Hour)
}
