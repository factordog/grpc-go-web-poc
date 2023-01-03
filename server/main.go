package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpcserver/api/grpctime"
	"log"
	"net"
	"time"
)

const (
	listenAddress = "0.0.0.0:9090"
)

type timeService struct {
	// NOTE: this has to be here to satisfy the interface, otherwise the compiler will complain
	pb.UnimplementedTimeServiceServer
}

func (t *timeService) GetCurrentTime(ctx context.Context, req *pb.GetCurrentTimeRequest) (*pb.GetCurrentTimeResponse, error) {
	log.Println("Got time here!")
	return &pb.GetCurrentTimeResponse{CurrentTime: time.Now().String()}, nil
}

func (t *timeService) GetGreet(ctx context.Context, req *pb.GetGreetRequest) (*pb.GetGreetResponse, error) {
	log.Println("Greetings!")
	return &pb.GetGreetResponse{Greet: "Hello!"}, nil
}

func (t *timeService) StreamTimeUpdates(rect *pb.GetCurrentTimeRequest, stream pb.TimeService_StreamTimeUpdatesServer) error {
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration", i)
		if err := stream.Send(&pb.GetCurrentTimeResponse{CurrentTime: time.Now().String()}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}

	return nil
}

func main() {
	log.Printf("Time service starting on %s", listenAddress)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	fmt.Println(lis, s)
	pb.RegisterTimeServiceServer(s, &timeService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
