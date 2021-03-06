package cmd

import (
	"fmt"
	"io/ioutil"

	p "github.com/datasage-io/datasage-cli/policy-ops"
	pb "github.com/datasage-io/datasage/src/proto/policy"
	"github.com/spf13/cobra"
)

var policy pb.ReadPolicyYAMLFile

//class represents
var readPolicycmd = &cobra.Command{
	Use:   "policy",
	Short: "Policy Commands For Manipulating Policy in Datasage",
	Long:  ` Policy Commands to do Manipulation on Policy in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Store File
		var file string
		if len(args) != 0 {
			file = args[0]
		}
		readFile, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		//To Store in Request
		policy.Policy = string(readFile)

		//Send to Server
		response, err := p.ReadPolicy(policy)
		fmt.Println(response.GetMessage())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(readPolicycmd)
	readPolicycmd.Flags().StringVarP(&policy.Policy, "f", "f", "", "Read Policy YAML File")
}
