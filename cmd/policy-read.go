package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	p "github.com/datasage-io/datasage-cli/policy-ops"
	pb "github.com/datasage-io/datasage/src/proto/policy"
	"github.com/spf13/cobra"
)

// To Decalre Policy
var fileURL string

// To Store Policy Content
var policy pb.ReadPolicyYAMLFile

// class represents
var readPolicycmd = &cobra.Command{
	Use:   "policy",
	Short: "Policy Commands For Manipulating Policy in Datasage",
	Long:  ` Policy Commands to do Manipulation on Policy in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Check file Status
		fileStatus, err := os.Stat(fileURL)
		if err != nil {
			fmt.Println("File Doesn't Exists")
		} else if fileStatus.Size() == 0 {
			fmt.Println("File is empty")
		} else {
			//Read File Path
			data, err := ioutil.ReadFile(fileURL)
			if err != nil {
				fmt.Println("File reading error", err)
				return nil
			}

			//To Store in Request
			policy.Policy = string(data)

			//Send to Server
			response, err := p.ReadPolicy(policy)
			fmt.Println(response.GetMessage())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(readPolicycmd)
	readPolicycmd.Flags().StringVarP(&fileURL, "f", "f", "", "Apply Policy")
}
