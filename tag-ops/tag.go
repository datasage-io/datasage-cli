package tag

import (
	"context"
	"os"

	pb "github.com/datasage-io/datasage-cli/proto/tag"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

//connectClient - To connect tag Client
func connectClient() (pb.TagClient, error) {
	gRPC := "localhost:8089"

	if val, ok := os.LookupEnv("DATASAGE_SERVICE"); ok {
		gRPC = val
	}
	var client pb.TagClient
	connection, err := grpc.Dial(gRPC, grpc.WithInsecure())
	if err != nil {
		log.Error().Msg("Error while connecting to grpc " + err.Error())
		return client, err
	}
	client = pb.NewTagClient(connection)
	return client, nil
}

//AddTag - To Add New tag
func AddTag(options pb.AddRequest) (pb.MessageResponse, error) {
	//Connect grpc Client
	client, err := connectClient()
	if err != nil {
		return pb.MessageResponse{Message: ""}, err
	}
	//Add Tag
	response, err := client.AddTag(context.Background(), &options)
	if err != nil {
		return pb.MessageResponse{Message: ""}, err
	}
	return pb.MessageResponse{Message: response.GetMessage()}, nil
}

//ListTag - List All Tag
func ListTag(options pb.ListRequest) (pb.ListResponse, error) {
	//Connect grpc  Client
	client, err := connectClient()
	if err != nil {
		return pb.ListResponse{TagResponse: nil, Count: 0}, err
	}
	//List Tag
	response, err := client.ListTag(context.Background(), &options)
	if err != nil {
		return pb.ListResponse{TagResponse: nil, Count: 0}, err
	}
	return pb.ListResponse{TagResponse: response.GetTagResponse(), Count: response.GetCount()}, nil
}

//DeleteTag - Delete a Tag
func DeleteTag(options pb.DeleteRequest) (pb.MessageResponse, error) {
	//Connect grpc  Client
	client, err := connectClient()
	if err != nil {
		return pb.MessageResponse{Message: ""}, err
	}
	//Delete Tag
	response, err := client.DeleteTag(context.Background(), &options)
	if err != nil {
		return pb.MessageResponse{Message: ""}, err
	}
	return pb.MessageResponse{Message: response.GetMessage()}, nil
}
