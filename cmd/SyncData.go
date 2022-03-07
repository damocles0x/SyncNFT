package cmd

import (
	"github.com/spf13/cobra"
)

/**
Used to sync data
*/
func SyncCmd() *cobra.Command {
	var startNum int64
	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "s",
		Long:  "It will sync the latest block ",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	syncCmd.Flags().Int64VarP(&startNum, "startNum", "s", 0, "input blockNum")
	return syncCmd
}
