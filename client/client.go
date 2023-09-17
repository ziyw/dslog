package main

import (
	"bytes"
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/ziyw/dslog/dslog"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "server address")
)

func DslogInfo(msg string) (string, error) {
	buf := new(bytes.Buffer)
	logger := slog.New(slog.NewTextHandler(buf, nil))
	logger.Info(msg)
	return buf.String(), nil
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:% v", err)
	}
	defer conn.Close()
	c := pb.NewDslogClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	logContent, _ := DslogInfo("This is formated log")

	r, err := c.AddLog(ctx, &pb.LogRequest{Content: *&logContent})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("Response: %s", r.GetStatus())
}
