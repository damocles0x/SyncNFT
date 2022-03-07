package cmd

import (
	"SyncNFT/handler"
	"github.com/spf13/cobra"
)

func GetOracleDataCmd() *cobra.Command {
	var startNum, page int64
	scanCmd := &cobra.Command{
		Use:   "data",
		Short: "o",
		Long:  "It will init oracle data by bscscan ",
		RunE: func(cmd *cobra.Command, args []string) error {
			handler.CrawlData(startNum, page)
			return nil
		},
	}

	scanCmd.Flags().Int64VarP(&startNum, "start", "s", 1, "input start num")
	scanCmd.Flags().Int64VarP(&page, "page", "p", 1, "input page num")
	return scanCmd
}
