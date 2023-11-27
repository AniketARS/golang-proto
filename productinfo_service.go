package main

import (
	"context"
	pb "productinfo/ecommerce"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	productMap map[string]*pb.Product
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Error while generating Product ID: %v", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, ok := s.productMap[in.Value]
	if ok {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product(%v) does not exist.", in.Value)
}

func (s *server) AllProduct(ctx context.Context, empty *pb.Empty) (*pb.Products, error) {
	// products := make([]pb.Product, 0, len(s.productMap))
	products := []*pb.Product{}
	for _, value := range s.productMap {
		products = append(products, value)
	}
	return &pb.Products{Products: products}, status.New(codes.OK, "").Err()
}
