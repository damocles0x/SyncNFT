package cmd

import (
	"SyncNFT/db"
	"SyncNFT/handler"
	"SyncNFT/utils"
	"github.com/spf13/cobra"
)

/**
Used to init data
*/
func InitCmd() *cobra.Command {
	var startNum int64
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "i",
		Long:  "It will init data",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := utils.GetClient()
			address := db.GetContractAddress()
			resultMap := utils.StringArrayToMap(address)
			handler.SyncData(client,
				startNum, resultMap)
			return nil
		},
	}

	initCmd.Flags().Int64VarP(&startNum, "startNum", "s", 0, "input blockNum")
	return initCmd
}
