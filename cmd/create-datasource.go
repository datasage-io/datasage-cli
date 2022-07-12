package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/datasource-ops"
	pb "github.com/datasage-io/datasage-cli/proto/datasource"
	"github.com/spf13/cobra"
)

var create pb.AddRequest
var datadomain, name, decription, dstype, version, host, port, user, password string

//datasource represents the datasource of datasage
var createDatasourceCmd = &cobra.Command{
	Use:   "add",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Store Data from command line
		create.DataDomain = datadomain
		create.Name = name
		create.Description = decription
		create.Type = dstype
		create.Version = version
		create.Host = host
		create.Port = port
		create.User = user
		create.Password = password
		//Send to Server
		stream, err := datasource.AddDatasource(create)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		if err != nil {
			return err
		}
		fmt.Println(response.GetMessage())
		return nil
	},
}

func init() {
	datasourceCmd.AddCommand(createDatasourceCmd)
	createDatasourceCmd.Flags().StringVarP(&datadomain, "datadomain", "", "", "input your data domain")
	createDatasourceCmd.Flags().StringVarP(&name, "name", "", "", "input your datasource name")
	createDatasourceCmd.Flags().StringVarP(&decription, "description", "", "", "input your datasource description")
	createDatasourceCmd.Flags().StringVarP(&dstype, "type", "", "", "input your datasource type")
	createDatasourceCmd.Flags().StringVarP(&version, "version", "", "", "input your datasource version")
	createDatasourceCmd.Flags().StringVarP(&host, "host", "", "", "input your host")
	createDatasourceCmd.Flags().StringVarP(&port, "port", "", "", "input your port")
	createDatasourceCmd.Flags().StringVarP(&user, "user", "", "", "input your user")
	createDatasourceCmd.Flags().StringVarP(&password, "password", "", "", "input your password")
}
