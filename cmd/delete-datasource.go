package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/datasource"
	pb "github.com/datasage-io/datasage-cli/proto/datasource"
	"github.com/spf13/cobra"
)

var delete pb.DeleteDatasourceRequest
var id int64

//datasource represents the datasource of datasage
var deleteDatasourceCmd = &cobra.Command{
	Use:   "delete",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		delete.Id = id
		fmt.Println("CLI Message -- ", delete)
		//Send to Server
		stream, err := datasource.DeleteDatasource(delete)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		fmt.Println(response.GetMessage())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteDatasourceCmd)
	deleteDatasourceCmd.Flags().Int64VarP(&id, "id", "", 0, "input your datasource id")
}
