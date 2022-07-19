package cmd

import (
	"github.com/spf13/cobra"
)

//class
var classCmd = &cobra.Command{
	Use:   "class",
	Short: "Class Commands For Class in policy of Datasage",
	Long:  ` Class Commands to do List Class Create Class and Delete Class in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(classCmd)
}
