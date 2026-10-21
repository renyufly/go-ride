package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-ride/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", handleTripPreview)

	server := &http.Server{
		Addr: httpAddr,
		Handler: mux,
	}

	/* Graceful shutdown  */
	// 监听可能出现的错误信息
	serverErrors := make(chan error, 1)

	// 在 goroutine上启动server
	go func() {
		log.Printf("HTTP server error: %s", httpAddr)
		serverErrors <- server.ListenAndServe()

	}()

	shutdown := make(chan os.Signal, 1)
	// `SIGINT`：例如终端按下 `Ctrl+C`
	// `SIGTERM`：例如 Docker 或 Kubernetes 请求终止容器
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// 处理channel
	select{
	case err := <- serverErrors:
		log.Printf("Error starting the server: %v", err)
	case sig := <- shutdown:
		log.Printf("Server is shuting down due to %v", sig)

		// 设置停机超时时间, 如果超时，执行预先定义的取消、持久化或强制退出策略
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Could not shut down the server gracefully :%v", err)
			server.Close()
		}
	}


}
