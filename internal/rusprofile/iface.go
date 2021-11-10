package rusprofile

import (
	grpc "rusprofile/pkg/grpc"
)

type Service interface {
	GetCompany(inn int64) (*grpc.Company, error)
}
