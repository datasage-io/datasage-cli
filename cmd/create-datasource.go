package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/datasage-io/datasage-cli/datasource-ops"
	pb "github.com/datasage-io/datasage/src/proto/datasource"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
)

var create pb.AddRequest
var datadomain, name, decription, dstype, version, host, port, user, password string

//datasource represents the datasource of datasage
var createDatasourceCmd = &cobra.Command{
	Use:   "add",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//RecommendedPolicy
		var recommedPolicy = []string{"GDPR Audit", "PII Audit", "SOC 2 Audit", "HIPAA Audit", "UDI Audit"}
		//To Store Data from command line
		create.DataDomain = datadomain
		create.Name = name
		create.Description = decription
		create.Type = dstype
		create.Version = version
		create.Host = host
		create.Port = port
		create.User = user
		create.Password = password
		//Operation Start From Here
		fmt.Println("Submitting to add datasource")
		response, err := datasource.AddDatasource(create)
		if err != nil {
			return nil
		}
		if response.GetStatusCode() == codes.OK.String() {
			fmt.Println("Datasource saved successfully")
			time.Sleep(2 * time.Second)
			fmt.Println("Submitting For Scan Datasource")
			scanResult, err := datasource.ScanDatasource(pb.ScanRequest{Name: create.Name})
			if err != nil {
				fmt.Printf("Status Code - %v :: Error Message - %v", scanResult.GetStatusCode(), scanResult.GetMessage())
				return nil
			}
			if scanResult.GetStatusCode() == codes.OK.String() {
				fmt.Println("Scan Completed")
				time.Sleep(2 * time.Second)
				fmt.Println("System Recommend following policies are")

				for i, policy := range recommedPolicy {
					fmt.Println(i+1, " - ", policy)
				}
				fmt.Println("Do you want to Apply the policies? Yes/No")
				var choice string
				fmt.Scanf("%s", &choice)
				if strings.ToLower(choice) == "yes" || strings.ToLower(choice) == "y" {
					//Get Policy Id to apply recommended policy
					var policyIds []int64
					fmt.Println("Enter the id of Recommended policy")
					reader := bufio.NewReader(os.Stdin)
					input, err := reader.ReadString('\n')
					if err != nil {
						fmt.Println("An error occured while reading input. Please try again", err)
						return nil
					}
					for _, val := range strings.Fields(input) {
						id, _ := strconv.ParseInt(val, 10, 64)
						policyIds = append(policyIds, id)
					}

					recommendedresult, err := datasource.ApplyPolicy(pb.ApplyPolicyRequest{Id: policyIds})
					if err != nil {
						fmt.Printf("Status Code - %v :: Error Message - %v", recommendedresult.GetStatusCode(), recommendedresult.GetMessage())
					}
					if recommendedresult.GetStatusCode() == codes.OK.String() {
						time.Sleep(2 * time.Second)
						fmt.Println("Policy has been applied")
					}
				} else {
					fmt.Println("Process Completed")
				}

			}
		}
		return nil
	},
}

func init() {
	datasourceCmd.AddCommand(createDatasourceCmd)
	createDatasourceCmd.Flags().StringVarP(&datadomain, "datadomain", "", "", "input your data domain")
	createDatasourceCmd.Flags().StringVarP(&name, "name", "", "", "input your datasource name")
	createDatasourceCmd.Flags().StringVarP(&decription, "description", "", "", "input your datasource description")
	createDatasourceCmd.Flags().StringVarP(&dstype, "type", "", "", "input your datasource type")
	createDatasourceCmd.Flags().StringVarP(&version, "version", "", "", "input your datasource version")
	createDatasourceCmd.Flags().StringVarP(&host, "host", "", "", "input your host")
	createDatasourceCmd.Flags().StringVarP(&port, "port", "", "", "input your port")
	createDatasourceCmd.Flags().StringVarP(&user, "user", "", "", "input your user")
	createDatasourceCmd.Flags().StringVarP(&password, "password", "", "", "input your password")
}
