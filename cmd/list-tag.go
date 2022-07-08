package cmd

import (
	"github.com/datasage-io/datasage-cli/output"
	pb "github.com/datasage-io/datasage-cli/proto/tag"
	"github.com/datasage-io/datasage-cli/tag-ops"
	"github.com/spf13/cobra"
)

var listTag pb.ListTagRequest

//datasource represents the datasource of datasage
var listTagCmd = &cobra.Command{
	Use:   "list",
	Short: "Tag Commands For Manipulating Tag in Datasage",
	Long:  ` Tag Commands to do List Tag, Create Tag and Delete Tag in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		n := len(args)
		if n > 0 {
			listTag.Tagname = args[0]
		}
		//Send to Server
		stream, err := tag.ListTag(listTag)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		tbl := output.New("ID", "NAME", "DESCRIPTION", "CLASS", "GENERATEDBY", "CREATEDAT", "UPDATEDAT")
		for _, t := range response.GetTagResponse() {
			tbl.AddRow(t.TagId, t.TagName, t.Description, t.Class, t.GeneratedBy, t.CreatedAt, t.UpdatedAt)
		}
		tbl.Print()
		return nil
	},
}

func init() {
	tagCmd.AddCommand(listTagCmd)
	listTagCmd.Flags().StringVar(&all, "", "", "input your all to get all tag ")
}
