package main

import (
	"fmt"
	"github.com/pyroscope-io/client/pyroscope"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	p, err := pyroscope.Start(pyroscope.Config{
		// 项目名字, 只能英文不能中文
		ApplicationName: "test-demo",
		// 按实际自己 pyroscope 服务的地址
		ServerAddress: "http://127.0.0.1:4040",
		Logger:        pyroscope.StandardLogger,
		// 如果开启了 PYROSCOPE_AUTH_INGESTION_ENABLED,
		// 并且按照以上步骤添加了 Key, 那么把 Key 放到此处即可
		AuthToken: "psx-AVmbAtgL9sHsm4YAZJSk-j_C0-EP-5sh7gOIYUjpeP_W",

		// by default all profilers are enabled,
		// but you can select the ones you want to use:
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = p.Stop()
	}()

	// 开启一个协程，每秒打印一次内存使用报告
	go func() {
		for {
			mem := &runtime.MemStats{}
			runtime.ReadMemStats(mem)
			fmt.Fprintf(os.Stderr, "alloc=%d, total=%d\n", mem.Alloc, mem.TotalAlloc)
			time.Sleep(time.Second)
		}
	}()

	ch := make(chan int)
	<-ch
}
