package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/tag-ops"
	pb "github.com/datasage-io/datasage/src/proto/tag"
	"github.com/spf13/cobra"
)

var createTag pb.AddRequest
var tagname, tagdescription string

//createTagCmd represents the tag of datasage
var createTagCmd = &cobra.Command{
	Use:   "add",
	Short: "Tag Commands For Tag in policy of Datasage",
	Long:  ` Tag Commands to do List Tag Create Tag and Delete Tag in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Store Data from command line
		createTag.Name = tagname
		createTag.Description = tagdescription
		//to Sotre Class names
		for _, val := range args {
			createTag.Class = append(createTag.Class, val)
		}
		//Send to Server
		response, err := tag.AddTag(createTag)
		if err != nil {
			return err
		}
		fmt.Println(response.GetMessage())
		return nil
	},
}

func init() {
	tagCmd.AddCommand(createTagCmd)
	createTagCmd.Flags().StringVarP(&tagname, "name", "n", "", "input your tag name")
	createTagCmd.Flags().StringVarP(&tagdescription, "description", "d", "", "input your tag description")
	createTagCmd.Flags().StringArrayVarP(&createTag.Class, "class", "c", nil, "input your class name")
}
