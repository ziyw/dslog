package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/ziyw/dslog/dslog"
	"google.golang.org/grpc"
)

var repo Repo

const ServerPort = ":50051"

type server struct {
	pb.UnimplementedDslogServer
}

func (s *server) SendLog(ctx context.Context, in *pb.LogRequest) (*pb.LogResponse, error) {
	log.Printf("Received: %v %v %v", in.Timestamp.AsTime().Format(time.RFC3339), in.LogType, in.LogMsg)
	repo.Insert(in.Timestamp.AsTime(), in.LogType, in.LogMsg)
	return &pb.LogResponse{Status: "OKAY"}, nil
}

func main() {
	err := repo.Connect()
	if err != nil {
		log.Fatal("error setting up Repo", err)
	}
	defer repo.Close()

	lis, err := net.Listen("tcp", ServerPort)
	if err != nil {
		log.Fatalf("server fail to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDslogServer(s, &server{})
	log.Printf("server is listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
