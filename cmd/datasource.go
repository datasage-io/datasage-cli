package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//datasource represents the datasource of datasage
var datasource = &cobra.Command{
	Use:   "datasource",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		action, _ := cmd.Flags().GetString("list")

		switch action {
		case "list":
			ListDataSources()
		case "create":
			CreateDataSources()
		case "delete":
			DeleteDataSources()
		default:
			fmt.Println("No Action Found")
		}

		return nil
	},
}

func ListDataSources() {
	fmt.Println("List all the datasource")
}

func CreateDataSources() {
	fmt.Println("Create a new datasource")
}

func DeleteDataSources() {
	fmt.Println("Delete an existing datasource")
}

func init() {
	rootCmd.AddCommand(datasource)
	datasource.PersistentFlags().String("list", "", "List all the datasource")
	datasource.PersistentFlags().String("create", "", "Create a new datasource")
	datasource.PersistentFlags().String("delete", "", "delete an existing datasource")
}
