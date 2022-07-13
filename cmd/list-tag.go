package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/output"
	pb "github.com/datasage-io/datasage-cli/proto/tag"
	"github.com/datasage-io/datasage-cli/tag-ops"
	"github.com/spf13/cobra"
)

var listTag pb.ListRequest
var firstTag, lastTag, limitTag int

//datasource represents the datasource of datasage
var listTagCmd = &cobra.Command{
	Use:   "list",
	Short: "Tag Commands For Manipulating Tag in Datasage",
	Long:  ` Tag Commands to do List Tag, Create Tag and Delete Tag in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//Limit
		listTag.Limit = int64(limitTag)
		//first
		listTag.First = int64(firstTag)
		//last
		listTag.Last = int64(lastTag)
		//To Store Array of ID's
		if len(args) != 0 {
			for _, val := range args {
				listTag.Id = append(listTag.Id, val)
			}
		}
		//Send to Server
		response, err := tag.ListTag(listTag)
		if err != nil {
			return err
		}
		//Count Datasource
		if listTag.Count {
			fmt.Println("Total Tag is --- ", response.GetCount())
			return nil
		}
		tbl := output.New("ID", "NAME", "DESCRIPTION", "CLASS", "CREATEDAT")
		for _, t := range response.GetTagResponse() {
			tbl.AddRow(t.Id, t.Name, t.Description, t.Class, t.CreatedAt)
		}
		tbl.Print()
		return nil
	},
}

func init() {
	tagCmd.AddCommand(listTagCmd)
	listTagCmd.Flags().IntVarP(&limit, "limit", "", 0, "limit the tag")
	listTagCmd.Flags().IntVarP(&first, "first", "", 0, "list first the tag")
	listTagCmd.Flags().IntVarP(&last, "last", "", 0, "list last the tag")
	listTagCmd.Flags().BoolVarP(&listTag.Count, "count", "", false, "list count the tag")
	listTagCmd.Flags().StringVarP(&listTag.Name, "name", "", "", "List filter by name tag")
	listTagCmd.Flags().StringArrayVarP(&listTag.Class, "class", "", nil, "List filter by type class")
	listTagCmd.Flags().StringArrayVarP(&listTag.Id, "id", "", nil, "Get Tag By Id")
}
