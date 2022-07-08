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
func AddTag(options pb.CreateTagRequest) (pb.Tag_AddTagClient, error) {
	//Connect grpc Client
	client, err := connectClient()
	if err != nil {
		return nil, err
	}
	//Add Datasource
	response, err := client.AddTag(context.Background(), &options)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//ListTag - List All Tag
func ListTag(options pb.ListTagRequest) (pb.Tag_ListTagClient, error) {
	//Connect grpc  Client
	client, err := connectClient()
	if err != nil {
		return nil, err
	}
	//List Tag
	response, err := client.ListTag(context.Background(), &options)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//DeleteTag - Delete a Tag
func DeleteTag(options pb.DeleteTagRequest) (pb.Tag_DeleteTagClient, error) {
	//Connect grpc  Client
	client, err := connectClient()
	if err != nil {
		return nil, err
	}
	//Delete Tag
	response, err := client.DeleteTag(context.Background(), &options)
	if err != nil {
		return nil, err
	}
	return response, nil
}
