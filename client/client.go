package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/ziyw/dslog/dslog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const INFO = "INFO"

type dslog struct {
	conn   *grpc.ClientConn
	client pb.DslogClient
}

func (d *dslog) Run() error {
	serverAddr := "localhost:50051"
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	client := pb.NewDslogClient(conn)

	d.conn = conn
	d.client = client
	return nil
}

func (d *dslog) Stop() {
	fmt.Println("stop log client")
	d.conn.Close()
}

func main() {

	var logClient dslog
	logClient.Run()
	defer logClient.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tr := &pb.TimeRange{
		StartTime: timestamppb.New(time.Now().Add(-time.Hour * 3)),
		EndTime:   timestamppb.New(time.Now())}
	stream, err := logClient.client.GetByTimeRange(ctx, tr)
	if err != nil {
		log.Fatal("error geting stream", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("%v.GetByTimeRange(_) = _, %v", logClient.client, err)
		}
		fmt.Println(msg)
	}

}
