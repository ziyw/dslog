package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/ziyw/dslog/dslog"
	"google.golang.org/grpc"
)

const ServerPort = "50051"

type server struct {
	pb.UnimplementedDslogServer
}

func (s *server) SendLog(ctx context.Context, in *pb.LogRequest) (*pb.LogResponse, error) {
	log.Printf("Received: %v %v %v", in.Timestamp.AsTime().Format(time.RFC3339), in.LogType, in.LogMsg)

	// TODO: replace this with connect to postgreSQL

	// err := s.PersistLog(in.GetContent())
	// if err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

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
