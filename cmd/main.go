package main

import (
	"fmt"
	"log"
	"net"

	"github.com/14jasimmtp/product-svc/pkg/config"
	"github.com/14jasimmtp/product-svc/pkg/db"
	"github.com/14jasimmtp/product-svc/pkg/product/pb"
	"github.com/14jasimmtp/product-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal("error in config", err)
	}

	h := db.Connection(c.DB_URL)

	lis, err := net.Listen("tcp", c.PORT)
	if err != nil {
		log.Fatal("error in listening", err)
	}

	fmt.Println("Product Service Started")

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failes to serve grpc client :", err)
	}
}
