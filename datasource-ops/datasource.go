package datasource

import (
	"context"
	"os"

	pb "github.com/datasage-io/datasage-cli/proto/datasource"
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
func AddDatasource(options pb.AddRequest) (pb.Datasource_AddDatasourceClient, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return nil, err
	}
	//Add Datasource
	response, err := client.AddDatasource(context.Background(), &options)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//ListDatasource - List All Datasource
func ListDatasource(options pb.ListRequest) (pb.Datasource_ListDatasourceClient, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return nil, err
	}
	//List Datasource
	response, err := client.ListDatasource(context.Background(), &options)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//DeleteDatasource - Delete a Datasource
func DeleteDatasource(options pb.DeleteRequest) (pb.Datasource_DeleteDatasourceClient, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return nil, err
	}
	//Delete Datasource
	response, err := client.DeleteDatasource(context.Background(), &options)
	if err != nil {
		return nil, err
	}
	return response, nil
}
