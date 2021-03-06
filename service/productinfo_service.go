package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"productinfo/service/ecommerce"
	pb "productinfo/service/ecommerce"

	"github.com/gofrs/uuid"
)

//server abstraction
type server struct {
	productMap map[string]*pb.Product
	orderMap   map[string]*pb.Order
	ecommerce.UnimplementedProductInfoServer
}

func (s *server) GetOrder(ctx context.Context, orderID *pb.OrderID) (*pb.Order, error) {
	ord := s.orderMap[orderID.Id]
	return ord, nil
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating ProductID", err)
	}

	in.Id = out.String()

	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	s.productMap[in.Id] = in

	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()

}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist", in.Value)
}
