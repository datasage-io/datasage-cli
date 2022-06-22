package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "datasage",
	Short: "A CLI Utility to help manage Datasage",
	Long: `CLI Utility to help manage Datasage
	
Datasage is tool to protect your data.
	  `,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
