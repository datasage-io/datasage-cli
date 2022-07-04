package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/datasource"
	pb "github.com/datasage-io/datasage-cli/proto/datasource"
	"github.com/spf13/cobra"
)

var create pb.AddDatasourceRequest
var datadomain, dsname, dsdecription, dstype, dsversion, host, port, user, password string

//datasource represents the datasource of datasage
var createDatasourceCmd = &cobra.Command{
	Use:   "add",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Store Data from command line
		create.DataDomain = datadomain
		create.DsName = dsname
		create.DsDescription = dsdecription
		create.DsType = dstype
		create.DsVersion = dsversion
		create.Host = host
		create.Port = port
		create.User = user
		create.Password = password
		fmt.Println("CLI Message -- ", create)
		//Send to Server
		stream, err := datasource.AddDatasource(create)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		fmt.Println(response.GetMessage())
		return nil
	},
}

func init() {
	datasourceCmd.AddCommand(createDatasourceCmd)
	createDatasourceCmd.Flags().StringVar(&datadomain, "datadomain", "", "input your data domain")
	createDatasourceCmd.Flags().StringVar(&dsname, "dsname", "", "input your datasource name")
	createDatasourceCmd.Flags().StringVar(&dsdecription, "dsdecription", "", "input your datasource description")
	createDatasourceCmd.Flags().StringVar(&dstype, "dstype", "", "input your datasource type")
	createDatasourceCmd.Flags().StringVar(&dsversion, "dsversion", "", "input your datasource version")
	createDatasourceCmd.Flags().StringVar(&host, "host", "", "input your host")
	createDatasourceCmd.Flags().StringVar(&port, "port", "", "input your port")
	createDatasourceCmd.Flags().StringVar(&user, "user", "", "input your user")
	createDatasourceCmd.Flags().StringVar(&password, "password", "", "input your password")
}
