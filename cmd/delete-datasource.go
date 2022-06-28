package cmd

import (
	"github.com/spf13/cobra"
)

//datasource represents the datasource of datasage
var delete_datasource = &cobra.Command{
	Use:   "datasource",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(delete_datasource)
	delete_datasource.Flags().String("delete", "", "delete an existing datasource")
}
