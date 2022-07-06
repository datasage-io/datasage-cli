package cmd

import (
	"github.com/datasage-io/datasage-cli/datasource-ops"
	"github.com/datasage-io/datasage-cli/output"
	pb "github.com/datasage-io/datasage-cli/proto/datasource"
	"github.com/spf13/cobra"
)

var list pb.ListDatasourceRequest
var rhost, rport, ruser, rpassword, first, last, all string

//datasource represents the datasource of datasage
var listDatasourceCmd = &cobra.Command{
	Use:   "list",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Store command Argument Value
		// list.Host = rhost
		// list.Port = rport
		// list.User = ruser
		// list.Password = rpassword
		// list.First = first
		// list.Last = last
		list.All = all
		//Send to Server
		stream, err := datasource.ListDatasource(list)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		tbl := output.New("ID", "DATA DOMAIN", "NAME", "DESCRIPTION", "TYPE", "VERSION", "KEY", "CREATEDAT", "DELETEDAT")
		for _, ds := range response.GetListAllDatasources() {
			tbl.AddRow(ds.Id, ds.DataDomain, ds.DsDescription, ds.DsType, ds.DsVersion, ds.DsKey, ds.CreatedAt, ds.Deleted)
		}
		tbl.Print()
		return nil
	},
}

func init() {
	datasourceCmd.AddCommand(listDatasourceCmd)
	// listDatasourceCmd.Flags().StringVar(&rhost, "host", "", "input your host")
	// listDatasourceCmd.Flags().StringVar(&rport, "port", "", "input your port")
	// listDatasourceCmd.Flags().StringVar(&ruser, "user", "", "input your user")
	// listDatasourceCmd.Flags().StringVar(&rpassword, "password", "", "input your password")
	// listDatasourceCmd.Flags().StringVar(&first, "first", "", "input your first")
	// listDatasourceCmd.Flags().StringVar(&last, "last", "", "input your last")
	listDatasourceCmd.Flags().StringVar(&all, "all", "", "input your all")
}
