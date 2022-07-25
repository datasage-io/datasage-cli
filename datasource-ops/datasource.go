package datasource

import (
	"context"
	"os"

	pb "github.com/datasage-io/datasage/src/proto/datasource"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

//connectClient - To connect datasource Client
func connectClient() (pb.DatasourceClient, error) {
	gRPC := "localhost:8089"

	if val, ok := os.LookupEnv("DATASAGE_SERVICE"); ok {
		gRPC = val
	}
	var client pb.DatasourceClient
	connection, err := grpc.Dial(gRPC, grpc.WithInsecure())
	if err != nil {
		log.Error().Msg("Error while connecting to grpc " + err.Error())
		return client, err
	}
	client = pb.NewDatasourceClient(connection)
	return client, nil
}

//AddDatasource - To Add New datasource
func AddDatasource(options pb.AddRequest) (pb.MessageResponse, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return pb.MessageResponse{Message: ""}, err
	}
	//Add Datasource
	response, err := client.AddDatasource(context.Background(), &options)
	if err != nil {
		return pb.MessageResponse{Message: ""}, err
	}
	return pb.MessageResponse{Message: response.GetMessage()}, nil
}

//ListDatasource - List All Datasource
func ListDatasource(options pb.ListRequest) (pb.ListResponse, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return pb.ListResponse{ListAllDatasources: nil, Count: 0}, err
	}
	//List Datasource
	response, err := client.ListDatasource(context.Background(), &options)
	if err != nil {
		return pb.ListResponse{ListAllDatasources: nil, Count: 0}, err
	}
	return pb.ListResponse{ListAllDatasources: response.GetListAllDatasources(), Count: response.GetCount()}, nil
}

//DeleteDatasource - Delete a Datasource
func DeleteDatasource(options pb.DeleteRequest) (pb.MessageResponse, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return pb.MessageResponse{Message: ""}, err
	}
	//Delete Datasource
	response, err := client.DeleteDatasource(context.Background(), &options)
	if err != nil {
		return pb.MessageResponse{Message: ""}, err
	}
	return pb.MessageResponse{Message: response.GetMessage()}, nil
}

//GetDatasourceLogs - Datasource Log
func GetDatasourceLogs(options pb.DatasourceLogRequest) (pb.DatasourceLogResponse, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return pb.DatasourceLogResponse{}, err
	}
	//Get Datasource Logs
	response, err := client.LogDatasource(context.Background(), &options)
	return pb.DatasourceLogResponse{DatasourceLog: response.GetDatasourceLog()}, nil
}
