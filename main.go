package main

import (
	"SyncNFT/db"
	"SyncNFT/handler"
	"SyncNFT/log"
	"SyncNFT/utils"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./errorLog", "log", time.Hour*24*14, time.Hour*24)
	//cmd.Execute() 15545400

	//调试用
	//15844391 27073
	client := utils.GetClient()
	address := db.GetContractAddress()
	resultMap := utils.StringArrayToMap(address)
	handler.SyncData(client,
		15545400, resultMap)
	//handler.CrawlData(1,274)

}
