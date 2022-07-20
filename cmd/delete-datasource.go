package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/datasource-ops"
	pb "github.com/datasage-io/datasage/src/proto/datasource"
	"github.com/spf13/cobra"
)

var delete pb.DeleteRequest

//datasource represents the datasource of datasage
var deleteDatasourceCmd = &cobra.Command{
	Use:   "delete",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Store ID
		for _, val := range args {
			delete.Id = append(delete.Id, val)
		}

		//Delete All
		if delete.IsDeleteAll {
			delete.IsDeleteAll = true
		}
		//Send to Server
		response, err := datasource.DeleteDatasource(delete)
		if err != nil {
			return err
		}
		fmt.Println(response.GetMessage())
		return nil
	},
}

func init() {
	datasourceCmd.AddCommand(deleteDatasourceCmd)
	deleteDatasourceCmd.Flags().StringArrayVarP(&delete.Id, "id", "d", nil, "input your datasource id's")
	deleteDatasourceCmd.Flags().BoolVarP(&delete.IsDeleteAll, "all", "", false, "delete all datasource's")
}
