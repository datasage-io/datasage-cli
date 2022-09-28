package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/datasage-io/datasage-cli/datasource-ops"
	"github.com/datasage-io/datasage-cli/utils/constants"
	pb "github.com/datasage-io/datasage/src/proto/datasource"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
)

var create pb.AddRequest
var datadomain, name, decription, dstype, version, host, port, user, password string

// datasource represents the datasource of datasage
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
		//  ==== Start ====  //
		//Check Status
		dsStatus, err := datasource.GetStatus(pb.StatusRequest{DsName: create.Name})
		if err != nil {
			fmt.Println(constants.FailedWhileFetchingStatus)
			return nil
		}
		if dsStatus.GetDsStatus() == "" {
			//Add Datasource
			response, err := addDatasource()
			if err != nil {
				fmt.Println(response)
				return nil
			}
			fmt.Println(response)
			fmt.Println(constants.DataSourceScanInprogress)
			//Scan Datasource
			scanResult, err := scanDatasource(create.Name)
			if err != nil {
				fmt.Println(scanResult)
				return nil
			}
			fmt.Println(constants.DataSourceInitialScanCompleted)
			fmt.Println(constants.DataSourcePeriodicScanProcess)
			fmt.Println(constants.DataSourcePeriodicScanWaiting)
			//Recommended Policy
			Recommendedpolicy(create.Name)
		} else {
			switch dsStatus.GetDsStatus() {
			case constants.DataSourceScanInprogress:
				//Scan Datasource
				scanResult, err := scanDatasource(create.Name)
				if err != nil {
					fmt.Println(scanResult)
					return nil
				}
				fmt.Println(scanResult)
				fmt.Println(constants.DataSourcePeriodicScanProcess)
				fmt.Println(constants.DataSourcePeriodicScanWaiting)
				//Recommended Policy
				Recommendedpolicy(create.Name)
			case constants.DataSourcePeriodicScanProcess, constants.DataSourcePeriodicScanWaiting:
				//Recommended Policy
				Recommendedpolicy(create.Name)
			}
		}
		//  ====  End  ====  //
		return nil
	},
}

// Add Datsource
func addDatasource() (string, error) {
	response, err := datasource.AddDatasource(create)
	if err != nil {
		return constants.DataSourceAddFailed, err
	}
	return response.GetMessage(), nil
}

// Scan Datasource
func scanDatasource(dsName string) (string, error) {
	scanResult, err := datasource.ScanDatasource(pb.ScanRequest{Name: dsName})
	if err != nil {
		return scanResult.GetMessage(), err
	}
	return scanResult.GetMessage(), nil
}

// Recommended Policy
func Recommendedpolicy(dsName string) {
	//Periodically Check
	for i := 0; i < viper.GetInt("datasource.numberOfIteration"); i++ {
		//Waiting for status
		time.Sleep(time.Duration(viper.GetInt64("datasource.sleepTime")) * time.Second)

		fmt.Println(constants.DataSourcePeriodicScanCompleted)
		fmt.Println("There are Serveral Recommended Policy for Datasource ")
		recommendedPolicies, err := datasource.GetRecommendedPolicies(pb.RecommendedpolicyRequest{DsName: create.Name})
		if err != nil {
			fmt.Println(constants.DefaultPoliciesIdentificationFailed)
		}
		fmt.Println(constants.DefaultPoliciesIdentified)
		for _, policy := range recommendedPolicies.GetPolicy() {
			fmt.Println(policy.PolicyId, ".", policy.PolicyName)
		}
		fmt.Print("Do you want to Apply the Recommended policies? Yes/No :: ")
		var choice string
		fmt.Scanln(&choice)
		if strings.ToLower(choice) == "yes" || strings.ToLower(choice) == "y" {
			//Get Policy Id to apply recommended policy
			var policyIds []int64
			for {
				var id int64
				fmt.Print("Enter the id of Recommended policy :: ")
				fmt.Scanln(&id)
				policyIds = append(policyIds, id)
				fmt.Print("Do you want to apply one more Recommended Policy - Yes/No :: ")
				var more string
				fmt.Scanln(&more)
				if strings.ToLower(more) == "no" || strings.ToLower(more) == "n" {
					break
				}
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
