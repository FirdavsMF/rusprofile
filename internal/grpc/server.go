package grpc

import (
	"context"
	"rusprofile/internal/rusprofile"
	grpc "rusprofile/pkg/grpc"
)

type grpcServer struct {
	service rusprofile.Service
	grpc.UnimplementedRusprofileServer
}

func NewGrpcServer(service rusprofile.Service) grpc.RusprofileServer {
	return &grpcServer{service: service}
}

func (g *grpcServer) GetCompany(ctx context.Context, request *grpc.GetCompanyRequest) (*grpc.Company, error) {
	return g.service.GetCompany(request.GetInn())
}
