package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/ziyw/dslog/dslog"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedDslogServer
}

func (s *server) AddLog(ctx context.Context, in *pb.LogRequest) (*pb.LogResponse, error) {
	log.Printf("Received: %v", in.GetContent())

	err := s.PersistLog(in.GetContent())
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return &pb.LogResponse{Status: "OKAY"}, nil
}

// TODO: replace by using log generated time, not receiving time
func (s *server) PersistLog(content string) error {
	cur := time.Now().UTC()
	filename := fmt.Sprintf(cur.Format("2006-01-02T15:00UTC"))
	fmt.Println(filename)

	file, err := os.OpenFile(filename+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error open file: ", err)
		return err
	}
	defer file.Close()

	if _, err := file.WriteString("\n" + content); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDslogServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
