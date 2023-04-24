package cmd

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/pingcap/log"
	"github.com/pingcap/test-infra/caselib/pkg/consistency"
	"github.com/pingcap/test-infra/caselib/pkg/mysql"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	host      string
	port      int
	username  string
	password  string
	databases []string
	table     string
	all       bool

	crc32Cmd = &cobra.Command{
		Use:   "crc32",
		Short: "crc32 returns crc32 checksum of databases or tables",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, port))
			if err != nil {
				log.Fatal("Failed to open sql session", zap.Error(err))
			}
			crc32Checker := consistency.Crc32Checker{MySQLClient: mysql.NewMySQLClient(db)}
			if all {
				res, err := crc32Checker.Crc32UserDatabases(threads)
				if err != nil {
					log.Fatal("Failed to get user databases crc32", zap.Error(err))
				}
				log.Info("Crc32 checksum finished", zap.Uint32("Crc32", res))
				return
			}
			if len(databases) != 0 {
				res, err := crc32Checker.Crc32Databases(databases, threads)
				if err != nil {
					log.Fatal("Failed to get databases crc32", zap.Error(err))
				}
				log.Info("Crc32 checksum finished", zap.Uint32("Crc32", res))
				return
			}
			if len(table) > 0 {
				dbAndTable := strings.Split(table, ".")
				if len(dbAndTable) != 2 {
					log.Fatal("Wrong parameter, required db.table format", zap.String("input table", table))
				}
				res, err := crc32Checker.Crc32Table(dbAndTable[0], dbAndTable[1])
				if err != nil {
					log.Fatal("Failed to get table crc32", zap.Error(err))
				}
				log.Info("Crc32 checksum finished", zap.Uint32("Crc32", res))
				return
			}
			log.Fatal("At least specify one of the parameters: --all, --databases, --table")
		},
	}
)

func Execute() error {
	return crc32Cmd.Execute()
}

func init() {
	crc32Cmd.PersistentFlags().StringVarP(&host, "host", "H", "127.0.0.1", "hostname")
	crc32Cmd.PersistentFlags().IntVarP(&port, "port", "P", 4000, "port")
	crc32Cmd.PersistentFlags().StringVarP(&username, "username", "u", "root", "username")
	crc32Cmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password")
	crc32Cmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "crc32 all user databases")
	crc32Cmd.PersistentFlags().StringArrayVarP(&databases, "databases", "d", []string{}, "crc32 databases")
	crc32Cmd.PersistentFlags().StringVarP(&table, "table", "t", "", "crc32 tables, example: db.table")
	crc32Cmd.PersistentFlags().IntVarP(&threads, "threads", "T", 10, "threads number")
	rootCmd.AddCommand(crc32Cmd)
}
