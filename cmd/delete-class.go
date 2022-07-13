package cmd

import (
	"fmt"

	c "github.com/datasage-io/datasage-cli/class-ops"
	pb "github.com/datasage-io/datasage-cli/proto/class"
	"github.com/spf13/cobra"
)

var deleteClass pb.DeleteRequest

//Class represents the class of datasage
var deleteClassCmd = &cobra.Command{
	Use:   "delete",
	Short: "Class Commands For Manipulating Class in Datasage",
	Long:  ` Class Commands to do List Class Data , Create Class and Delete Class in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Store ID
		for _, val := range args {
			deleteClass.Id = append(deleteClass.Id, val)
		}

		//Delete All
		if deleteClass.IsDeleteAll {
			deleteClass.IsDeleteAll = true
		}
		//Send to Server
		response, err := c.DeleteClass(deleteClass)
		if err != nil {
			return err
		}
		fmt.Println(response.GetMessage())
		return nil
	},
}

func init() {
	classCmd.AddCommand(deleteClassCmd)
	deleteClassCmd.Flags().StringArrayVarP(&delete.Id, "id", "d", nil, "input your class id's")
	deleteClassCmd.Flags().BoolVarP(&delete.IsDeleteAll, "all", "", false, "delete all class's")
}
