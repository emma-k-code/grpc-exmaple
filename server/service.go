package server

import (
	"context"

	"example/pb"
)

// Service 範例Service
type Service struct {
	pb.UnimplementedServiceServer
}

// Sum 相加
func (*Service) Sum(ctx context.Context, request *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{Result: request.A + request.B}, nil
}
