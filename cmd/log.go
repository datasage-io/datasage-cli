package cmd

import (
	"github.com/spf13/cobra"
)

//log cmd for datasource
var logCmd = &cobra.Command{
	Use:   "logs",
	Short: "Log Commands For Display Log for Datasage",
	Long:  ` Log Commands For Display Log for Datasage `,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
