package main

import (
	"context"
	"fmt"
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

func (dslog *dslog) Info(logMsg string) {
	// TODO: move slog display to server UI side
	// buf := new(bytes.Buffer)
	// logger := slog.New(slog.NewTextHandler(buf, nil))
	// logger.Info(msg)
	// infoMsg := buf.String()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ts := timestamppb.Now()
	dslog.client.SendLog(ctx, &pb.LogRequest{Timestamp: ts, LogType: INFO, LogMsg: logMsg})
}

var logClient dslog

func main() {
	logClient.Run()
	defer logClient.Stop()

	logClient.Info("This is the first correct message")
}
