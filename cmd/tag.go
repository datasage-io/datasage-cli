package cmd

import (
	"github.com/spf13/cobra"
)

//tag
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Tag Commands For Tag in policy of Datasage",
	Long:  ` Tag Commands to do List Tag Create Tag and Delete Tag in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
