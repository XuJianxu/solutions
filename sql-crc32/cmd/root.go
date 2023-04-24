package cmd

import (
	"github.com/pingcap/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sql-crc32",
	Short: "sql-crc32 provides a set of tools related to crc32 & sql",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger, props, err := log.InitLogger(&log.Config{Level: logLevel})
		if err != nil {
			panic(err)
		}
		log.ReplaceGlobals(logger, props)
	},
}

var logLevel string

func init() {
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "log level: debug, info, warn, error")
}
