package cmd

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/datasage-io/datasage-cli/datasource-ops"
	"github.com/datasage-io/datasage-cli/output"
	pb "github.com/datasage-io/datasage/src/proto/datasource"
	"github.com/spf13/cobra"
)

var log pb.DatasourceLogRequest

//Export
var export struct {
	json bool
	csv  bool
}

//log represents the datasource log
var logDatasourceCmd = &cobra.Command{
	Use:   "datasource",
	Short: "Log Datasource Commands For Display Log for Datasage",
	Long:  ` Log Datasource Commands For Display Log for Datasage `,
	RunE: func(cmd *cobra.Command, args []string) error {
		//Send to Server
		response, err := datasource.GetDatasourceLogs(log)
		if err != nil {
			return err
		}

		//To Store Datasource Log
		var datasourceLog []pb.DatasourceLogResponseStruct
		for _, log := range response.GetDatasourceLog() {
			datasourceLog = append(datasourceLog, *log)
		}

		//export in JSON file format
		if export.json {
			//Convert logs into Json file
			file, _ := json.MarshalIndent(datasourceLog, "", "")
			fileName := "datasource_log" + time.Now().UTC().Format("2006-01-02 15:04:05") + ".json"
			_ = ioutil.WriteFile(fileName, file, 0644)
			return nil
		}

		//export in CSV file format
		if export.csv {
			fileName := "system_log_" + time.Now().UTC().Format("2006-01-02 15:04:05") + ".csv"
			file, _ := os.Create(fileName)
			writer := csv.NewWriter(file)
			defer writer.Flush()

			header := []string{"Datasource", "Database", "Table", "Column", "Tags",
				"Classes", "LastScanTime"}
			_ = writer.Write(header)

			for _, logs := range datasourceLog {
				//Classes
				classes := strings.Join(logs.Classes, ", ")
				//Tags
				tags := strings.Join(logs.Tags, ", ")
				var row []string
				row = append(row, logs.Datasource, logs.Database, logs.Table, logs.Column, tags,
					classes, logs.LastScanTime)
				_ = writer.Write(row)
			}
			return nil
		}

		//Print logs In Table
		tbl := output.New("DATASOURCE", "DATABASE", "TABLE", "COLUMN", "TAGS",
			"CLASSES", "LAST-SCAN-TIME")
		for _, logs := range datasourceLog {
			tbl.AddRow(logs.Datasource, logs.Database, logs.Table, logs.Column, logs.Tags,
				logs.Classes, logs.LastScanTime)
		}
		//Print Table
		tbl.Print()
		return nil
	},
}

func init() {
	logCmd.AddCommand(logDatasourceCmd)
	logDatasourceCmd.Flags().StringVarP(&log.Datasource, "datasource", "", "", "List Logs based on Datasource")
	logDatasourceCmd.Flags().StringVarP(&log.Database, "database", "", "", "List Logs based on Database")
	logDatasourceCmd.Flags().StringVarP(&log.Table, "table", "", "", "List Logs based on Table")
	logDatasourceCmd.Flags().StringArrayVarP(&log.Columns, "columns", "", nil, "List Logs based on Columns")

	//Export In JSON - CSV
	logDatasourceCmd.Flags().BoolVar(&export.json, "json", false, "Export the log in JSON Format")
	logDatasourceCmd.Flags().BoolVar(&export.csv, "csv", false, "Export the log in CSV Format")
}
