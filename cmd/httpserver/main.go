package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"rusprofile/pkg/config"
	grpc2 "rusprofile/pkg/grpc"
)

func main() {
	cfg, err := config.GetHttpServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = grpc2.RegisterRusprofileHandlerFromEndpoint(ctx, mux, "grpcserver:"+cfg.Grpc.Port, opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":"+cfg.Http.Port, mux))

}
