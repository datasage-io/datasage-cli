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
			listTag.Tag = args[0]
		} else {
			listTag.Tag = "all"
		}
		//Send to Server
		stream, err := tag.ListTag(listTag)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		tbl := output.New("ID", "NAME", "DESCRIPTION", "CLASS")
		for _, t := range response.GetTagResponse() {
			tbl.AddRow(t.TagId, t.TagName, t.TagDescription, t.TagClass)
		}
		tbl.Print()
		return nil
	},
}

func init() {
	tagCmd.AddCommand(listTagCmd)
	listTagCmd.Flags().StringVarP(&listTag.Tag, "list", "l", "", "List Tag")
}
