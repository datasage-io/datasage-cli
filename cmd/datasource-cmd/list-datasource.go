package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/datasource"
	pb "github.com/datasage-io/datasage-cli/proto/datasource"
	"github.com/spf13/cobra"
)

var list pb.ListDatasourceRequest
var rhost, rport, ruser, rpassword string

//datasource represents the datasource of datasage
var listDatasourceCmd = &cobra.Command{
	Use:   "list",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		list.Host = rhost
		list.Port = rport
		list.User = ruser
		list.Password = rpassword
		fmt.Println("CLI Message -- ", list)
		//Send to Server
		stream, err := datasource.ListDatasource(list)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		fmt.Println(response.GetListAllDatasources())
		return nil
	},
}

func init() {
	datasourceCmd.AddCommand(listDatasourceCmd)
	listDatasourceCmd.Flags().StringVar(&rhost, "host", "", "input your host")
	listDatasourceCmd.Flags().StringVar(&rport, "port", "", "input your port")
	listDatasourceCmd.Flags().StringVar(&ruser, "user", "", "input your user")
	listDatasourceCmd.Flags().StringVar(&rpassword, "password", "", "input your password")
}
