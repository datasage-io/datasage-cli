package cmd

import (
	"fmt"

	c "github.com/datasage-io/datasage-cli/class-ops"
	pb "github.com/datasage-io/datasage-cli/proto/class"
	"github.com/spf13/cobra"
)

var deleteClass pb.DeleteClassRequest

//Class represents the class of datasage
var deleteClassCmd = &cobra.Command{
	Use:   "delete",
	Short: "Class Commands For Manipulating Class in Datasage",
	Long:  ` Class Commands to do List Class Data , Create Class and Delete Class in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		n := len(args)
		if n > 0 {
			deleteClass.Id = args[0]
		}
		fmt.Println("CLI Message -- ", deleteClass)
		//Send to Server
		stream, err := c.DeleteClass(deleteClass)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		fmt.Println(response.GetMessage())
		return nil
	},
}

func init() {
	classCmd.AddCommand(deleteClassCmd)
	deleteClassCmd.Flags().StringVarP(&deleteClass.Id, "id", "d", "", "input your Class id")
}
