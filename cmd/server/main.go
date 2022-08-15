package main

import (
	"flag"
	"fmt"
	"gprc/pb"
	"gprc/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore(), service.NewDiskImageStore("image"), service.NewInMemoryRatingStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listerner, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listerner)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
