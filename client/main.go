package main

import (
	"context"
	"log"
	"time"

	pb "productinfo/client/ecommerce"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	name := "Apple iPhone 11"
	description := "Description of Apple device"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description})

	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}

	log.Printf("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})

	if err != nil {
		log.Fatalf("Couldl not get product: %v", err)
	}

	log.Printf("Product: ", product.String())

}
