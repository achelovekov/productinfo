package main

import (
	"log"
	"net"

	pb "productinfo/service/ecommerce"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	orderMapVal := make(map[string]*pb.Order)

	orderMapVal["id1"] = &pb.Order{Id: "1", Items: []string{"apple"}, Description: "decr1", Price: "145.000", Destination: "MOSCKOW"}
	orderMapVal["id2"] = &pb.Order{Id: "2", Items: []string{"SAMsumg"}, Description: "decr2", Price: "146.000", Destination: "london"}

	srv := &server{
		orderMap: orderMapVal,
	}
	pb.RegisterProductInfoServer(s, srv)

	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
