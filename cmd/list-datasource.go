package cmd

import (
	"github.com/datasage-io/datasage-cli/datasource-ops"
	"github.com/datasage-io/datasage-cli/output"
	pb "github.com/datasage-io/datasage-cli/proto/datasource"
	"github.com/spf13/cobra"
)

var list pb.ListDatasourceRequest

//datasource represents the datasource of datasage
var listDatasourceCmd = &cobra.Command{
	Use:   "list",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//To Store command Argument Value
		n := len(args)
		if n > 0 {
			list.All = args[0]
		}
		//Send to Server
		stream, err := datasource.ListDatasource(list)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		tbl := output.New("ID", "DATA DOMAIN", "NAME", "DESCRIPTION", "TYPE", "VERSION", "KEY", "CREATEDAT", "DELETEDAT")
		for _, ds := range response.GetListAllDatasources() {
			tbl.AddRow(ds.DsId, ds.DsDatadomain, ds.DsDescription, ds.DsType, ds.DsVersion, ds.DsKey, ds.CreatedAt)
		}
		tbl.Print()
		return nil
	},
}

func init() {
	datasourceCmd.AddCommand(listDatasourceCmd)

	listDatasourceCmd.Flags().StringVarP(&list.All, "all", "l", "", "input your all")
}
