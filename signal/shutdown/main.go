package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	// 模拟先关闭 http，再关闭其他依赖的情况
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Second * 10) // 请求的时间比shutdown等待的时间长，看看http会不会处理正在处理的请求，如果中途一直往里加请求会怎么样

		data, _ := json.Marshal(struct {
			Msg string `json:"msg"`
		}{
			Msg: "hello",
		})

		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(data)
	})
}

func main() {

	server := http.Server{
		Addr:    ":12345",
		Handler: nil,
	}

	go func() {
		log.Println("开启 HTTP 服务")
		if err := server.ListenAndServe(); err != nil || err == http.ErrServerClosed {
			log.Fatalf("server listen err:%s", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println("接收到信号", <-ch)

	log.Println("程序开始优雅关闭...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("关闭http服务失败", err)
		return
	}

	fmt.Println("程序优雅关闭成功")
}
