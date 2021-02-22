package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/luciano-nascimento/grpc-go-lab/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	//close connection when not used anymore
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddedUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Luciano",
		Email: "l@l.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Luciano",
		Email: "l@l.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status:", stream.Status, "-", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "l1",
			Name:  "luciano1",
			Email: "l1@l.com",
		},
		&pb.User{
			Id:    "l2",
			Name:  "luciano2",
			Email: "l2@l.com",
		},
		&pb.User{
			Id:    "l3",
			Name:  "luciano3",
			Email: "l3@l.com",
		},
		&pb.User{
			Id:    "l4",
			Name:  "luciano4",
			Email: "l4@l.com",
		},
		&pb.User{
			Id:    "l5",
			Name:  "luciano5",
			Email: "l5@l.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)

}

func AddedUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id:    "l1",
			Name:  "luciano1",
			Email: "l1@l.com",
		},
		&pb.User{
			Id:    "l2",
			Name:  "luciano2",
			Email: "l2@l.com",
		},
		&pb.User{
			Id:    "l3",
			Name:  "luciano3",
			Email: "l3@l.com",
		},
		&pb.User{
			Id:    "l4",
			Name:  "luciano4",
			Email: "l4@l.com",
		},
		&pb.User{
			Id:    "l5",
			Name:  "luciano5",
			Email: "l5@l.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}
			fmt.Printf("Recebendo user %v com status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
