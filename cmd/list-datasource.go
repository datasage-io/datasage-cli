package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/datasource"
	pb "github.com/datasage-io/datasage-cli/proto/datasource"
	"github.com/spf13/cobra"
)

var list pb.ListDatasourceRequest

//datasource represents the datasource of datasage
var listDatasourceCmd = &cobra.Command{
	Use:   "list",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		list.Host = "localhost"
		list.Port = "3306"
		list.User = "root"
		list.Password = "measroot"
		stream, err := datasource.ListDatasource(list)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		fmt.Println("Response is -- ", response.GetListAllDatasources())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listDatasourceCmd)
}
