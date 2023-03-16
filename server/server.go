package server

import (
	"example/pb"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Start 啟動Server
func Start() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("grpc Server啟動失敗,", err)
	}

	server := grpc.NewServer(
		// 設定 middleware
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			CheckUser(),
		)),
	)

	pb.RegisterServiceServer(server, &Service{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	log.Println("grpc Server啟動")

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Println("grpc Server異常關閉,", err)
		}
	}()

	// 接收程序關閉通知
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)

	<-ch
	signal.Stop(ch)

	// 優雅關閉
	server.GracefulStop()
}
