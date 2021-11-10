package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	grpcService "rusprofile/internal/grpc"
	"rusprofile/internal/rusprofile"
	"rusprofile/pkg/config"
	grpc2 "rusprofile/pkg/grpc"
)

func main() {
	cfg, err := config.GetGrpcServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	service := rusprofile.NewService(cfg.RusProfile.Host)

	rusprofileGrpc := grpcService.NewGrpcServer(service)

	server := grpc.NewServer()

	grpc2.RegisterRusprofileServer(server, rusprofileGrpc)

	ln, err := net.Listen("tcp", ":"+cfg.Grpc.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
