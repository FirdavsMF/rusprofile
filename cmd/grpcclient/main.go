package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	grpc2 "rusprofile/pkg/grpc"
)

// Тестовый grpc клиент, запрашивающий данные об ООО "Яндекс"
func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:8090", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := grpc2.NewRusprofileClient(conn)
	request := &grpc2.GetCompanyRequest{
		Inn: 7736207543,
	}
	response, err := client.GetCompany(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response)
}
