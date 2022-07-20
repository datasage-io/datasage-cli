package cmd

import (
	"fmt"

	c "github.com/datasage-io/datasage-cli/class-ops"
	"github.com/datasage-io/datasage-cli/output"
	pb "github.com/datasage-io/datasage/src/proto/class"
	"github.com/spf13/cobra"
)

var listClass pb.ListRequest
var lastClass, limitClass int

//Class represents the class of datasage
var listClassCmd = &cobra.Command{
	Use:   "list",
	Short: "Class Commands For Manipulating Class in Datasage",
	Long:  ` Class Commands to do List Class, Create Class and Delete Class in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//Limit
		listClass.Limit = int64(limitClass)
		//last
		listClass.Last = int64(lastClass)
		//To Store Array of ID's
		if len(args) != 0 {
			for _, val := range args {
				listClass.Id = append(listClass.Id, val)
			}
		}
		//Send to Server
		response, err := c.ListClass(listClass)
		if err != nil {
			return err
		}
		//Count Datasource
		if listClass.Count {
			fmt.Println("Total Class is --- ", response.GetCount())
			return nil
		}
		tbl := output.New("ID", "NAME", "DESCRIPTION", "TAG", "CREATEAT")
		for _, c := range response.GetClassResponse() {
			tbl.AddRow(c.Id, c.Name, c.Description, c.Tag, c.CreatedAt)
		}
		tbl.Print()
		return nil
	},
}

func init() {
	classCmd.AddCommand(listClassCmd)
	listClassCmd.Flags().IntVarP(&limitClass, "limit", "", 0, "limit the class")
	listClassCmd.Flags().IntVarP(&lastClass, "last", "", 0, "list last the class")
	listClassCmd.Flags().BoolVarP(&listClass.Count, "count", "", false, "list count the class")
	listClassCmd.Flags().StringVarP(&listClass.Name, "name", "", "", "List filter by name class")
	listClassCmd.Flags().StringArrayVarP(&listClass.Tag, "tag", "", nil, "List filter by type tag")
	listClassCmd.Flags().StringArrayVarP(&listClass.Id, "id", "", nil, "Get Tag By Id")
}
