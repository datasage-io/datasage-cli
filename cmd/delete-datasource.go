package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/datasource-ops"
	pb "github.com/datasage-io/datasage-cli/proto/datasource"
	"github.com/spf13/cobra"
)

var delete pb.DeleteDatasourceRequest

//datasource represents the datasource of datasage
var deleteDatasourceCmd = &cobra.Command{
	Use:   "delete",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		n := len(args)
		if n > 0 {
			delete.Id = args[0]
		}
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
	datasourceCmd.AddCommand(deleteDatasourceCmd)
	deleteDatasourceCmd.Flags().StringVarP(&delete.Id, "id", "d", "", "input your datasource id")
}
