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
func AddDatasource(options pb.AddRequest) (pb.AddResponse, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return pb.AddResponse{}, err
	}
	//Add Datasource
	response, err := client.AddDatasource(context.Background(), &options)
	if err != nil {
		return pb.AddResponse{StatusCode: response.GetStatusCode(), Message: response.GetMessage()}, err
	}
	return pb.AddResponse{StatusCode: response.GetStatusCode(), Message: response.GetMessage()}, err
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
func DeleteDatasource(options pb.DeleteRequest) (pb.DeleteResponse, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return pb.DeleteResponse{Message: ""}, err
	}
	//Delete Datasource
	response, err := client.DeleteDatasource(context.Background(), &options)
	if err != nil {
		return pb.DeleteResponse{StatusCode: response.GetStatusCode(), Message: response.GetMessage()}, err
	}
	return pb.DeleteResponse{StatusCode: response.GetStatusCode(), Message: response.GetMessage()}, nil
}

//GetDatasourceLogs - Datasource Log
func GetDatasourceLogs(options pb.LogRequest) (pb.LogResponse, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return pb.LogResponse{}, err
	}
	//Get Datasource Logs
	response, err := client.LogDatasource(context.Background(), &options)
	return pb.LogResponse{DatasourceLog: response.GetDatasourceLog()}, nil
}

//ScanDatasource - To Scan New datasource
func ScanDatasource(options pb.ScanRequest) (pb.ScanResponse, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return pb.ScanResponse{}, err
	}
	//Scan Datasource
	response, err := client.Scan(context.Background(), &options)
	if err != nil {
		return pb.ScanResponse{StatusCode: response.GetStatusCode(), Message: response.GetMessage()}, err
	}
	return pb.ScanResponse{StatusCode: response.GetStatusCode(), Message: response.GetMessage()}, err
}

//ApplyRecommendedPolicy - Apply Recommended Policy
func ApplyPolicy(options pb.ApplyPolicyRequest) (pb.ApplyPolicyResponse, error) {
	//Connect grpc datasource Client
	client, err := connectClient()
	if err != nil {
		return pb.ApplyPolicyResponse{}, err
	}
	//Scan Datasource
	response, err := client.ApplyPolicy(context.Background(), &options)
	if err != nil {
		return pb.ApplyPolicyResponse{StatusCode: response.GetStatusCode(), Message: response.GetMessage()}, err
	}
	return pb.ApplyPolicyResponse{StatusCode: response.GetStatusCode(), Message: response.GetMessage()}, err
}
