package policy

import (
	"context"
	"os"

	pb "github.com/datasage-io/datasage-cli/proto/policy"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

//connectClient - To connect Policy Client
func connectClient() (pb.PolicyClient, error) {
	gRPC := "localhost:8089"

	if val, ok := os.LookupEnv("DATASAGE_SERVICE"); ok {
		gRPC = val
	}
	var client pb.PolicyClient
	connection, err := grpc.Dial(gRPC, grpc.WithInsecure())
	if err != nil {
		log.Error().Msg("Error while connecting to grpc " + err.Error())
		return client, err
	}
	client = pb.NewPolicyClient(connection)
	return client, nil
}

//Read YAML File from CLI
func ReadPolicy(options pb.ReadPolicyYAMLFile) (*pb.PolicyResponse, error) {
	//Connect grpc Client
	client, err := connectClient()
	if err != nil {
		return &pb.PolicyResponse{Message: ""}, err
	}
	//Add Class
	response, err := client.ReadPolicy(context.Background(), &options)
	if err != nil {
		return &pb.PolicyResponse{Message: ""}, err
	}
	return &pb.PolicyResponse{Message: response.GetMessage()}, nil
}
