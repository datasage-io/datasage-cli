package cmd

import (
	"github.com/spf13/cobra"
)

//datasource represents the datasource of datasage
var datasourceCmd = &cobra.Command{
	Use:   "datasource",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(datasourceCmd)
}
