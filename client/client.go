package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/ziyw/dslog/dslog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr   = flag.String("addr", "localhost:50051", "server address")
	logstr = flag.String("log-string", "This is an log", "log string content")
)

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

	r, err := c.AddLog(ctx, &pb.LogRequest{Content: *logstr})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("AddingLogResponse: %s", r.GetStatus())
}
