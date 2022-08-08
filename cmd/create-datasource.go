package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/datasage-io/datasage-cli/datasource-ops"
	"github.com/datasage-io/datasage-cli/utils/constants"
	pb "github.com/datasage-io/datasage/src/proto/datasource"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		response, err := datasource.AddDatasource(create)
		if err != nil {
			fmt.Println(constants.DataSourceAddFailed)
			return nil
		}
		if response.GetStatusCode() == codes.OK.String() {
			fmt.Println(constants.DataSourceAddedSucessful)
			fmt.Println(constants.DataSourceScanInprogress)
			scanResult, err := datasource.ScanDatasource(pb.ScanRequest{Name: create.Name})
			if err != nil {
				fmt.Println(constants.DataSourceInitialScanFailed)
				return nil
			}
			if scanResult.GetStatusCode() == codes.OK.String() {
				fmt.Println(constants.DataSourceInitialScanCompleted)
				fmt.Println(constants.DataSourcePeriodicScanProcess)
				fmt.Println(constants.DataSourcePeriodicScanWaiting)

				//Periodically Check
				for i := 0; i < viper.GetInt("datasource.numberOfIteration"); i++ {
					//Waiting for status
					time.Sleep(time.Duration(viper.GetInt64("datasource.sleepTime")) * time.Second)
					dsStatus, err := datasource.GetStatus(pb.StatusRequest{DsName: create.Name})
					if err != nil {
						fmt.Println(constants.DataSourcePeriodicScanFailed)
					} else if dsStatus.GetStatusCode() == codes.OK.String() {
						fmt.Println(constants.DataSourcePeriodicScanCompleted)
						fmt.Println("There are Serveral Recommended Policy for Datasource ")
						recommendedPolicies, err := datasource.GetRecommendedPolicies(empty.Empty{})
						if err != nil {
							fmt.Println(constants.DefaultPoliciesIdentificationFailed)
						}
						fmt.Println(constants.DefaultPoliciesIdentified)
						for i, policy := range recommendedPolicies.GetPolicyName() {
							fmt.Println(i+1, ".", policy)
						}
						fmt.Println("Do you want to Apply the policies? Yes/No")
						var choice string
						fmt.Scanf("%s", &choice)
						if strings.ToLower(choice) == "yes" || strings.ToLower(choice) == "y" {
							//Get Policy Id to apply recommended policy
							var policyIds []int64
							fmt.Println("Enter the id of Recommended policy")
							reader := bufio.NewReader(os.Stdin)
							input, _ := reader.ReadString('\n')
							for _, val := range strings.Fields(input) {
								id, _ := strconv.ParseInt(val, 10, 64)
								policyIds = append(policyIds, id)
							}
							fmt.Println(constants.ApplyrecommendingPolicy)
							policyAppliedResult, err := datasource.ApplyPolicy(pb.ApplyPolicyRequest{Id: policyIds, DsName: create.Name})
							if err != nil {
								fmt.Printf("Status Code - %v :: Error Message - %v", policyAppliedResult.GetStatusCode(), policyAppliedResult.GetMessage())
							}
							if policyAppliedResult.GetStatusCode() == codes.OK.String() {
								fmt.Println("Policy has been applied")
								break
							}
						} else {
							fmt.Println("Process Completed")
							break
						}
					}
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
