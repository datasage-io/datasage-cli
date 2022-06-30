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
func AddDatasource(options pb.AddDatasourceRequest) (pb.Datasource_AddDatasourcesClient, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return nil, err
	}
	//Add Datasource
	response, err := client.AddDatasources(context.Background(), &options)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//ListDatasource - List All Datasource
func ListDatasource(options pb.ListDatasourceRequest) (pb.Datasource_ListDatasourcesClient, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return nil, err
	}
	//List Datasource
	response, err := client.ListDatasources(context.Background(), &options)
	if err != nil {
		return nil, err
	}
	return response, nil
}
