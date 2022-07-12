package cmd

import (
	"fmt"

	"github.com/datasage-io/datasage-cli/datasource-ops"
	"github.com/datasage-io/datasage-cli/output"
	pb "github.com/datasage-io/datasage-cli/proto/datasource"
	"github.com/spf13/cobra"
)

var list pb.ListRequest
var first, last, limit int

//datasource represents the datasource of datasage
var listDatasourceCmd = &cobra.Command{
	Use:   "list",
	Short: "Datasource Commands For Manipulating Datasource in Datasage",
	Long:  ` Datasource Commands to do List Data Datasource, Create Datasource and Delete Datasource in Datasage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//Limit
		list.Limit = int64(limit)
		//first
		list.First = int64(first)
		//last
		list.Last = int64(last)
		//Send to Server
		stream, err := datasource.ListDatasource(list)
		if err != nil {
			return err
		}
		response, err := stream.Recv()
		if err != nil {
			return err
		}
		//Count Datasource
		if list.Count {
			fmt.Println("Total Datasource is --- ", response.GetCount())
			return nil
		}
		tbl := output.New("ID", "DATA DOMAIN", "NAME", "DESCRIPTION", "TYPE", "VERSION", "KEY", "CREATEDAT")
		for _, ds := range response.GetListAllDatasources() {
			tbl.AddRow(ds.Id, ds.Datadomain, ds.Name, ds.Description, ds.Type, ds.Version, ds.Key, ds.CreatedAt)
		}
		tbl.Print()
		return nil
	},
}

func init() {
	datasourceCmd.AddCommand(listDatasourceCmd)
	listDatasourceCmd.Flags().IntVarP(&limit, "limit", "", 0, "limit the datasource")
	listDatasourceCmd.Flags().IntVarP(&first, "first", "", 0, "list first the datasource")
	listDatasourceCmd.Flags().IntVarP(&last, "last", "", 0, "list last the datasource")
	listDatasourceCmd.Flags().BoolVarP(&list.Count, "count", "", false, "list count the datasource")
	listDatasourceCmd.Flags().StringVarP(&list.Name, "name", "", "", "List filter by name datasource")
	listDatasourceCmd.Flags().StringArrayVarP(&list.Type, "type", "", nil, "List filter by type datasource")
	listDatasourceCmd.Flags().StringArrayVarP(&list.DataDomain, "domain", "", nil, "List filter by data domain datasource")
}
