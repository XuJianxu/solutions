package cmd

import (
	"database/sql"

	"github.com/pingcap/log"
	"github.com/pingcap/test-infra/caselib/pkg/consistency"
	"github.com/pingcap/test-infra/caselib/pkg/mysql"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	upstreamTs     string
	downstreamTs   string
	tidbCompareCmd = &cobra.Command{
		Use:   "tidb-compare",
		Short: "compare returns the different tables between 2 tidb databases",
		Run: func(cmd *cobra.Command, args []string) {
			upstreamDb, err := sql.Open("mysql", upstream)
			if err != nil {
				log.Fatal("Failed to open sql session", zap.String("dsn", upstream), zap.Error(err))
			}
			downstreamDb, err := sql.Open("mysql", downstream)
			if err != nil {
				log.Fatal("Failed to open sql session", zap.String("dsn", downstream), zap.Error(err))
			}
			if len(databases) != 0 {
				err := consistency.CompareCRC32CheckSum(upstreamDb, downstreamDb, upstreamTs, downstreamTs, threads, databases)
				if err != nil {
					log.Fatal("Failed to tidb-compare", zap.Error(err))
				}
			} else {
				client := mysql.NewMySQLClient(upstreamDb)
				dbs, err := client.GetUserDatabases()
				if err != nil {
					log.Fatal("Failed to get user databases", zap.Error(err))
				}
				err = consistency.CompareCRC32CheckSum(upstreamDb, downstreamDb, upstreamTs, downstreamTs, threads, dbs)
				if err != nil {
					log.Fatal("Failed to tidb-compare", zap.Error(err))
				}
			}
		},
	}
)

func init() {
	tidbCompareCmd.PersistentFlags().StringVarP(&upstream, "upstream", "U", "", "upstream DSN")
	tidbCompareCmd.PersistentFlags().StringVarP(&downstream, "downstream", "D", "", "downstream DSN")
	tidbCompareCmd.PersistentFlags().IntVarP(&threads, "threads", "T", 10, "threads number")
	tidbCompareCmd.PersistentFlags().StringArrayVarP(&databases, "databases", "d", []string{}, "crc32 databases")
	tidbCompareCmd.PersistentFlags().StringVarP(&upstreamTs, "upstream-ts", "", "", "upstream snapshot ts")
	tidbCompareCmd.PersistentFlags().StringVarP(&downstreamTs, "downstream-ts", "", "", "downstream snapshot ts")
	rootCmd.AddCommand(tidbCompareCmd)
}
