package summary

// import (
// 	"os"

// 	"github.com/rs/zerolog/log"
// 	"google.golang.org/grpc"
// )

// //connectClient - To connect summary Client
// func connectClient() (summary.SummaryClient, error) {
// 	gRPC := "localhost:8089"

// 	if val, ok := os.LookupEnv("DATASAGE_SERVICE"); ok {
// 		gRPC = val
// 	}
// 	var client summary.SummaryClient
// 	connection, err := grpc.Dial(gRPC, grpc.WithInsecure())
// 	if err != nil {
// 		log.Error().Msg("Error while connecting to grpc " + err.Error())
// 		return client, err
// 	}
// 	client = summary.NewSummaryClient(connection)
// 	return client, nil
// }
