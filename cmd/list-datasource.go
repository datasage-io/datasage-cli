package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/datasource"
	ds "github.com/datasage-io/datasage-cli/proto/datasource"
	"github.com/spf13/cobra"
)

//datasource represents the datasource of datasage
var listDatasourceCmd = &cobra.Command{
	Use:   "list",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := datasource.ListDatasource(ds.ListDatasourceRequest{
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "measroot",
		})
		if err != nil {
			return err
		}
		fmt.Println("Response is -- ", response)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listDatasourceCmd)
}
