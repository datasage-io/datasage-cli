package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/datasource"
	pb "github.com/datasage-io/datasage-cli/proto/datasource"
	"github.com/spf13/cobra"
)

var create pb.AddDatasourceRequest

//datasource represents the datasource of datasage
var createDatasourceCmd = &cobra.Command{
	Use:   "create",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		create.DataDomain = "datasage-cli"
		create.DsName = "AWS"
		create.DsDescription = "Datasage CLI description"
		create.DsType = "MySQL"
		create.DsVersion = "8"
		create.DsKey = "1258fghfg87fghf365"
		create.Host = "localhost"
		create.Port = "3306"
		create.User = "root"
		create.Password = "measroot"
		stream, err := datasource.AddDatasource(create)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		fmt.Println("Response is -- ", response.GetMessage())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createDatasourceCmd)
}
