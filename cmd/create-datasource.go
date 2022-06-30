package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/datasource"
	ds "github.com/datasage-io/datasage-cli/proto/datasource"
	"github.com/spf13/cobra"
)

//datasource represents the datasource of datasage
var createDatasourceCmd = &cobra.Command{
	Use:   "create",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := datasource.AddDatasource(ds.AddDatasourceRequest{
			DataDomain:    "datasage-cli",
			DsName:        "AWS",
			DsDescription: "Datasage CLI description",
			DsType:        "MySQL",
			DsVersion:     "8",
			DsKey:         "1258fghfg87fghf365",
			Host:          "localhost",
			Port:          "3306",
			User:          "root",
			Password:      "measroot",
		})
		if err != nil {
			return err
		}
		fmt.Println("Response is -- ", response)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createDatasourceCmd)
}
