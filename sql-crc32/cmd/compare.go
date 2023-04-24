package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pingcap/log"
	"github.com/pingcap/test-infra/caselib/pkg/mysql"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	upstream        string
	downstream      string
	threads         int
	resultFile      string
	fastFailOpt     bool
	tableNamePrefix string

	compareCmd = &cobra.Command{
		Use:   "compare",
		Short: "compare returns the different tables between 2 mysql compatible databases",
		Run: func(cmd *cobra.Command, args []string) {
			upstreamDb, err := sql.Open("mysql", upstream)
			if err != nil {
				log.Fatal("Failed to open sql session", zap.String("dsn", upstream), zap.Error(err))
			}
			downstreamDb, err := sql.Open("mysql", downstream)
			if err != nil {
				log.Fatal("Failed to open sql session", zap.String("dsn", downstream), zap.Error(err))
			}
			upstreamClient := mysql.NewMySQLClient(upstreamDb)
			downstreamClient := mysql.NewMySQLClient(downstreamDb)
			if len(databases) == 0 {
				databases, err = upstreamClient.GetUserDatabases()
				if err != nil {
					log.Fatal("Failed to get user databases from upstream db", zap.Error(err))
				}
			}
			upstreamDbTableMap, err := upstreamClient.GetTablesInDatabases(databases, tableNamePrefix)
			if err != nil {
				log.Fatal("Failed to get tables from databases on upstream db", zap.Strings("databases", databases), zap.Error(err))
			}
			downstreamDbTableMap, err := downstreamClient.GetTablesInDatabases(databases, tableNamePrefix)
			if err != nil {
				log.Fatal("Failed to get tables from databases on downstream db", zap.Strings("databases", databases), zap.Error(err))
			}
			// Check upstream and downstream db table map
			for database, upstreamTables := range upstreamDbTableMap {
				if downstreamTables, ok := downstreamDbTableMap[database]; ok {
					if len(upstreamTables) != len(downstreamTables) {
						log.Fatal("Table num doesn't match for database",
							zap.String("database", database),
							zap.Int("upstream table num", len(upstreamTables)),
							zap.Int("downstream table num", len(downstreamTables)),
							zap.Strings("upstream tables", upstreamTables),
							zap.Strings("downstream tables", downstreamTables),
						)
					}
				}
			}
			// Check Crc32
			results, err := compare(upstreamClient, downstreamClient, upstreamDbTableMap, threads, fastFailOpt)
			if err != nil {
				log.Fatal("compare failed", zap.Error(err))
			}
			content, err := json.Marshal(results)
			if err != nil {
				log.Fatal("json marshal failed", zap.Error(err))
			}
			err = ioutil.WriteFile(resultFile, content, 0644)
			if err != nil {
				log.Fatal("write result to result.", zap.Error(err))
			}
		},
	}
)

func init() {
	compareCmd.PersistentFlags().StringVarP(&upstream, "upstream", "U", "", "upstream DSN")
	compareCmd.PersistentFlags().StringVarP(&downstream, "downstream", "D", "", "downstream DSN")
	compareCmd.PersistentFlags().IntVarP(&threads, "threads", "T", 10, "threads number")
	compareCmd.PersistentFlags().StringVarP(&resultFile, "result-file", "r", "./result.json", "result file")
	compareCmd.PersistentFlags().BoolVarP(&fastFailOpt, "fast-fail", "f", false, "fast fail when crc32 comparison mismatch")
	compareCmd.PersistentFlags().StringArrayVarP(&databases, "databases", "d", []string{}, "crc32 databases")
	compareCmd.PersistentFlags().StringVarP(&tableNamePrefix, "table-name-prefix", "p", "", "table name prefix")
	rootCmd.AddCommand(compareCmd)
}

const (
	COMPARE_RESULT_UNKNOWN  int = 0
	COMPARE_RESULT_MATCH        = 1
	COMPARE_RESULT_MISMATCH     = 2
)

type CompareResult struct {
	Result          int    `json:"result" yaml:"result"`
	Database        string `json:"database" yaml:"database"`
	Table           string `json:"table" yaml:"table"`
	UpstreamCrc32   uint32 `json:"upstream-crc32" yaml:"upstream-crc32"`
	DownstreamCrc32 uint32 `json:"downstream-crc32" yaml:"downstream-crc32"`
}

func compare(upstreamClient, downstreamClient *mysql.MySQLClient, dbTableMap map[string][]string, threads int, fastFail bool) ([]CompareResult, error) {
	dbTables := make([]mysql.DbTable, 0)
	for database, tables := range dbTableMap {
		for _, table := range tables {
			dbTables = append(dbTables, mysql.DbTable{Database: database, Table: table})
		}
	}
	totalTables := len(dbTables)
	log.Info("total tables", zap.Int("table number", totalTables))
	if threads > totalTables {
		threads = totalTables
	}
	dbTableChan := make(chan mysql.DbTable, threads)
	resultChan := make(chan CompareResult, threads)
	errChan := make(chan error)
	for tid := 0; tid < threads; tid++ {
		go compareWorker(upstreamClient, downstreamClient, dbTableChan, resultChan, errChan)
	}
	results := make([]CompareResult, 0)
	failed := false
	for i, dbTable := range dbTables {
		go func() {
			log.Info(fmt.Sprintf("Dealing table progress: %d/%d", i+1, totalTables),
				zap.String("database", dbTable.Database),
				zap.String("table", dbTable.Table))
			dbTableChan <- dbTable
		}()
		log.Debug("Request result", zap.Int("tableId", i+1))
		select {
		case err := <-errChan:
			return nil, err
		case result := <-resultChan:
			if result.Result == COMPARE_RESULT_MATCH {
				log.Info("compare result match",
					zap.String("database", result.Database),
					zap.String("table", result.Table),
					zap.Uint32("upstream crc32", result.UpstreamCrc32),
					zap.Uint32("downstream crc32", result.DownstreamCrc32))
			} else {
				failed = true
				if fastFail {
					log.Fatal("compare result mismatch or unknown",
						zap.String("database", result.Database),
						zap.String("table", result.Table),
						zap.Uint32("upstream crc32", result.UpstreamCrc32),
						zap.Uint32("downstream crc32", result.DownstreamCrc32))
				}
				log.Error("compare result mismatch or unknown",
					zap.String("database", result.Database),
					zap.String("table", result.Table),
					zap.Uint32("upstream crc32", result.UpstreamCrc32),
					zap.Uint32("downstream crc32", result.DownstreamCrc32))
			}
			results = append(results, result)
		}
	}
	close(dbTableChan)
	close(resultChan)
	if failed {
		log.Error("Crc32 check failed. Detailed result refer to result file", zap.String("result file path", resultFile))
	} else {
		log.Info("Crc32 check passed.")
	}
	return results, nil
}

func compareWorker(upstreamClient, downstreamClient *mysql.MySQLClient, dbTableChan <-chan mysql.DbTable, compareResultChan chan<- CompareResult, errChan chan<- error) {
	for dbTable := range dbTableChan {
		upstreamDbTableChan := make(chan mysql.DbTable)
		upstreamChecksumChan := make(chan uint32)
		upstreamErrChan := make(chan error)
		downstreamDbTableChan := make(chan mysql.DbTable)
		downstreamChecksumChan := make(chan uint32)
		downstreamErrChan := make(chan error)
		go func() {
			upstreamDbTableChan <- dbTable
			downstreamDbTableChan <- dbTable
		}()

		go upstreamClient.Crc32TableWorker(upstreamDbTableChan, upstreamChecksumChan, upstreamErrChan)
		go downstreamClient.Crc32TableWorker(downstreamDbTableChan, downstreamChecksumChan, downstreamErrChan)

		select {
		case err := <-upstreamErrChan:
			errChan <- err
		case err := <-downstreamErrChan:
			errChan <- err
		case downstreamChecksum := <-downstreamChecksumChan:
			upstreamChecksum := <-upstreamChecksumChan
			result := CompareResult{
				Result:   COMPARE_RESULT_UNKNOWN,
				Database: dbTable.Database,
				Table:    dbTable.Table,
			}
			result.UpstreamCrc32 = upstreamChecksum
			result.DownstreamCrc32 = downstreamChecksum
			if upstreamChecksum != downstreamChecksum {
				result.Result = COMPARE_RESULT_MISMATCH
			} else {
				result.Result = COMPARE_RESULT_MATCH
			}
			compareResultChan <- result
		}
	}
}
