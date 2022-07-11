package cmd

import (
	"fmt"

	c "github.com/datasage-io/datasage-cli/class-ops"
	pb "github.com/datasage-io/datasage-cli/proto/class"
	"github.com/spf13/cobra"
)

var class pb.CreateClassRequest
var classname, classdescription, classtag string

//class represents
var createClassCmd = &cobra.Command{
	Use:   "add",
	Short: "Class Commands For Manipulating Class in Datasage",
	Long:  ` Class Commands to do List CLass, Create Class and Delete Class in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Store Data from command line
		class.ClassName = classname
		class.ClassDescription = classdescription
		class.ClassTag = classtag

		fmt.Println("CLI Message -- ", class)
		//Send to Server
		stream, err := c.AddClass(class)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		if err != nil {
			return err
		}
		fmt.Println(response.GetMessage())
		return nil
	},
}

func init() {
	classCmd.AddCommand(createClassCmd)
	createClassCmd.Flags().StringVarP(&classname, "name", "n", "", "input your class name")
	createClassCmd.Flags().StringVarP(&classdescription, "description", "d", "", "input your class description")
	createClassCmd.Flags().StringVarP(&classtag, "tag", "t", "", "input your class tag")
}
