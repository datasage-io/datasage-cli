package cmd

import (
	"fmt"

	pb "github.com/datasage-io/datasage-cli/proto/tag"
	"github.com/datasage-io/datasage-cli/tag-ops"
	"github.com/spf13/cobra"
)

var deleteTag pb.DeleteRequest

//Tag represents the Tag of datasage
var deleteTagCmd = &cobra.Command{
	Use:   "delete",
	Short: "Tag Commands For Manipulating Tag in Datasage",
	Long:  ` Tag Commands to do List Tag Data , Create Tag and Delete Tag in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Store ID
		if len(args) != 0 {
			for _, val := range args {
				deleteTag.Id = append(deleteTag.Id, val)
			}
		} else {
			fmt.Println("Invalid Command")
		}

		//Delete All
		if deleteTag.IsDeleteAll {
			deleteTag.IsDeleteAll = true
		}
		//Send to Server
		response, err := tag.DeleteTag(deleteTag)
		if err != nil {
			return err
		}
		fmt.Println(response.GetMessage())
		return nil
	},
}

func init() {
	tagCmd.AddCommand(deleteTagCmd)
	deleteTagCmd.Flags().StringArrayVarP(&deleteTag.Id, "id", "d", nil, "input your tag id's")
	deleteTagCmd.Flags().BoolVarP(&deleteTag.IsDeleteAll, "all", "", false, "delete all tag's")
}
