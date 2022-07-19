package cmd

import (
	"fmt"

	c "github.com/datasage-io/datasage-cli/class-ops"
	pb "github.com/datasage-io/datasage-cli/proto/class"
	"github.com/spf13/cobra"
)

var class pb.CreateRequest
var classname, classdescription string

//class represents
var createClassCmd = &cobra.Command{
	Use:   "add",
	Short: "Class Commands For Manipulating Class in Datasage",
	Long:  ` Class Commands to do List CLass, Create Class and Delete Class in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Store Data from command line
		class.Name = classname
		class.Description = classdescription
		//to Sotre Class names
		for _, val := range args {
			class.Tag = append(class.Tag, val)
		}

		//Send to Server
		response, err := c.AddClass(class)
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
	createClassCmd.Flags().StringArrayVarP(&class.Tag, "tag", "t", nil, "input your class tag")
}
