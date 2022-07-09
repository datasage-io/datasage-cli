package class

import (
	"context"
	"os"

	pb "github.com/datasage-io/datasage-cli/proto/class"
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
func AddClass(options pb.CreateClassRequest) (pb.Class_AddClassClient, error) {
	//Connect grpc Client
	client, err := connectClient()
	if err != nil {
		return nil, err
	}
	//Add Class
	response, err := client.AddClass(context.Background(), &options)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//ListClass - List All Class
func ListClass(options pb.ListClassRequest) (pb.Class_ListClassClient, error) {
	//Connect grpc  Client
	client, err := connectClient()
	if err != nil {
		return nil, err
	}
	//List Class
	response, err := client.ListClass(context.Background(), &options)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//DeleteClass - Delete a Class
func DeleteClass(options pb.DeleteClassRequest) (pb.Class_DeleteClassClient, error) {
	//Connect grpc  Client
	client, err := connectClient()
	if err != nil {
		return nil, err
	}
	//Delete Class
	response, err := client.DeleteClass(context.Background(), &options)
	if err != nil {
		return nil, err
	}
	return response, nil
}
