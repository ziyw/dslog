package main

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/ziyw/dslog/dslog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterDslogServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestConnection(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("fail dail bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewDslogClient(conn)
	resp, err := client.SendLog(ctx, &pb.LogRequest{Timestamp: timestamppb.Now(), LogType: "ERROR", LogMsg: "This is a test"})
	if err != nil {
		t.Fatalf("sendLog service failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
}
