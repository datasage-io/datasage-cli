package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//datasource represents the datasource of datasage
var list_datasource = &cobra.Command{
	Use:   "datasource",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("List all the datasources")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(list_datasource)
	list_datasource.Flags().String("list", "", "List all the datasource")
}
