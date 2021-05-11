package main

import (
	"fmt"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	repo "./src/repo"
	stubs "./src/stubs"
	logger "./src/logger"
	services "./src/services"
)

const (
	PORT = 9000
	CONNECTION_STRING = ""
)
func main() {

	log := logger.Logger{}.GetInstance()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Log("failed to listen: %v", err)
	}

	grpcServer := registerServers();
	log.Log("Starting server..")
	if err := grpcServer.Serve(lis); err != nil {
		log.Log("failed to serve: %s", err)
	}
}
func registerServers() *grpc.Server {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	stubs.RegisterAuthServiceServer(
		grpcServer, 
		services.AuthServer{
			Connection: repo.HolderDB{}.GetConnection(CONNECTION_STRING).(repo.HolderDB),
			Logger: logger.Logger{}.GetInstance(),
		}.Initiate(),
	)
	return grpcServer;
}