
package main

import (
	"context"
	pb "datacollector-edge/stages/origins/grpc_client/testing/simple"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewSimpleServiceClient(conn)
	stream, err := client.ServerStreamingRPC(context.Background(), &pb.SimpleInputData{Msg: "Simple Server"})

	for {
		in, err := stream.Recv()
		log.Println("Received value")
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
		log.Println("Got " + in.Msg)
	}
}
