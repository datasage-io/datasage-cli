package cmd

import (
	c "github.com/datasage-io/datasage-cli/class-ops"
	"github.com/datasage-io/datasage-cli/output"
	pb "github.com/datasage-io/datasage-cli/proto/class"
	"github.com/spf13/cobra"
)

var listClass pb.ListClassRequest

//Class represents the class of datasage
var listClassCmd = &cobra.Command{
	Use:   "list",
	Short: "Class Commands For Manipulating Class in Datasage",
	Long:  ` Class Commands to do List Class, Create Class and Delete Class in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		n := len(args)
		if n > 0 {
			listTag.Tag = args[0]
		}
		//Send to Server
		stream, err := c.ListClass(listClass)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		tbl := output.New("ID", "NAME", "DESCRIPTION", "TAG", "GENERATEDBY", "CREATEDAT")
		for _, c := range response.GetClassResponse() {
			tbl.AddRow(c.ClassId, c.ClassName, c.ClassDescription, c.ClassTag, c.GeneratedBy, c.CreatedAt)
		}
		tbl.Print()
		return nil
	},
}

func init() {
	classCmd.AddCommand(listClassCmd)
	listClassCmd.Flags().StringVarP(&listClass.Class, "all", "l", "", "input your all to get all class ")
}
