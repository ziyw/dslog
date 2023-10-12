package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/ziyw/dslog/dslog"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var repo Repo

func toLogMessage(logEntry LogEntry) pb.LogMessage {
	return pb.LogMessage{
		Timestamp: timestamppb.New(logEntry.createdAt),
		LogType:   logEntry.logType,
		LogMsg:    logEntry.logMsg,
	}
}

const ServerPort = ":50051"

type server struct {
	pb.UnimplementedDslogServer
}

func (s *server) Send(ctx context.Context, in *pb.LogMessage) (*pb.SendResponse, error) {
	log.Printf("Received: %v %v %v", in.Timestamp.AsTime().Format(time.RFC3339), in.LogType, in.LogMsg)
	id, err := repo.Insert(in.Timestamp.AsTime(), in.LogType, in.LogMsg)
	if err != nil {
		return nil, err
	}
	idStr := fmt.Sprintf("%d", id)
	return &pb.SendResponse{Id: idStr}, nil
}

// GetByTimeRange(*TimeRange, Dslog_GetByTimeRangeServer) error

func (s *server) GetByTimeRange(in *pb.TimeRange, stream pb.Dslog_GetByTimeRangeServer) error {
	entries := repo.GetByTimeRange(in.StartTime.AsTime(), in.EndTime.AsTime())
	for _, e := range entries {
		var entry pb.LogMessage = toLogMessage(e)
		stream.Send(&entry)
	}
	return nil
}

func main() {
	repo.Connect()
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
