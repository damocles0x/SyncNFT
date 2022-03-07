package cmd

import (
	"github.com/spf13/cobra"
)

/**
Used to init data
*/
func InitCmd() *cobra.Command {
	var startNum int64
	var endNum int64
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "i",
		Long:  "It will init data",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	initCmd.Flags().Int64VarP(&startNum, "startNum", "s", 0, "input blockNum")
	initCmd.Flags().Int64VarP(&endNum, "endNum", "e", 0, "input blockNum")
	return initCmd
}
