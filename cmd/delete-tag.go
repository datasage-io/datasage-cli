package cmd

import (
	"fmt"

	pb "github.com/datasage-io/datasage-cli/proto/tag"
	"github.com/datasage-io/datasage-cli/tag-ops"
	"github.com/spf13/cobra"
)

var deleteTag pb.DeleteTagRequest

//datasource represents the datasource of datasage
var deleteTagCmd = &cobra.Command{
	Use:   "delete",
	Short: "Tag Commands For Manipulating Tag in Datasage",
	Long:  ` Tag Commands to do List Tag Data , Create Tag and Delete Tag in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		n := len(args)
		if n > 0 {
			deleteTag.Id = args[0]
		}
		fmt.Println("CLI Message -- ", deleteTag)
		//Send to Server
		stream, err := tag.DeleteTag(deleteTag)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		fmt.Println(response.GetMessage())
		return nil
	},
}

func init() {
	tagCmd.AddCommand(deleteTagCmd)
	deleteTagCmd.Flags().StringVarP(&deleteTag.Id, "id", "d", "", "input your Tag id")
}
