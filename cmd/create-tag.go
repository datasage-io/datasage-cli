package cmd

import (
	"fmt"

	pb "github.com/datasage-io/datasage-cli/proto/tag"
	"github.com/datasage-io/datasage-cli/tag-ops"
	"github.com/spf13/cobra"
)

var createTag pb.CreateTagRequest
var tagname, tagdescription, tagclass string

//createTagCmd represents the tag of datasage
var createTagCmd = &cobra.Command{
	Use:   "add",
	Short: "Tag Commands For Tag in policy of Datasage",
	Long:  ` Tag Commands to do List Tag Create Tag and Delete Tag in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Store Data from command line
		createTag.TagName = tagname
		createTag.Description = tagdescription
		createTag.Class = tagclass
		fmt.Println("CLI Message -- ", createTag)
		//Send to Server
		stream, err := tag.AddTag(createTag)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		fmt.Println(response.GetMessage())
		return nil
	},
}

func init() {
	tagCmd.AddCommand(createTagCmd)
	createTagCmd.Flags().StringVar(&tagname, "name", "", "input your tag name")
	createTagCmd.Flags().StringVar(&tagdescription, "description", "", "input your tag description")
	createTagCmd.Flags().StringVar(&tagclass, "class", "", "input your class name")
}
