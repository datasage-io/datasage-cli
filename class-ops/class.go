package class

import (
	"context"
	"os"

	pb "github.com/datasage-io/datasage/src/proto/class"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

//connectClient - To connect class Client
func connectClient() (pb.ClassClient, error) {
	gRPC := "localhost:8089"

	if val, ok := os.LookupEnv("DATASAGE_SERVICE"); ok {
		gRPC = val
	}
	var client pb.ClassClient
	connection, err := grpc.Dial(gRPC, grpc.WithInsecure())
	if err != nil {
		log.Error().Msg("Error while connecting to grpc " + err.Error())
		return client, err
	}
	client = pb.NewClassClient(connection)
	return client, nil
}

//AddClass - To Add New Class
func AddClass(options pb.CreateRequest) (*pb.MessageResponse, error) {
	//Connect grpc Client
	client, err := connectClient()
	if err != nil {
		return &pb.MessageResponse{Message: ""}, err
	}
	//Add Class
	response, err := client.AddClass(context.Background(), &options)
	if err != nil {
		return &pb.MessageResponse{Message: ""}, err
	}
	return &pb.MessageResponse{Message: response.GetMessage()}, nil
}

//ListClass - List All Class
func ListClass(options pb.ListRequest) (*pb.ListResponse, error) {
	//Connect grpc  Client
	client, err := connectClient()
	if err != nil {
		return &pb.ListResponse{ClassResponse: nil, Count: 0}, err
	}
	//List Class
	response, err := client.ListClass(context.Background(), &options)
	if err != nil {
		return &pb.ListResponse{ClassResponse: nil, Count: 0}, err
	}
	return &pb.ListResponse{ClassResponse: response.GetClassResponse(), Count: response.GetCount()}, nil
}

//DeleteClass - Delete a Class
func DeleteClass(options pb.DeleteRequest) (*pb.MessageResponse, error) {
	//Connect grpc  Client
	client, err := connectClient()
	if err != nil {
		return &pb.MessageResponse{Message: ""}, err
	}
	//Delete Class
	response, err := client.DeleteClass(context.Background(), &options)
	if err != nil {
		return &pb.MessageResponse{Message: ""}, err
	}
	return &pb.MessageResponse{Message: response.GetMessage()}, nil
}
