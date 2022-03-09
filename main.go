package main

import (
	"SyncNFT/cmd"
	"SyncNFT/log"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./errorLog", "log", time.Hour*24*14, time.Hour*24)
	cmd.Execute()
	////15844391 27073
	//client := utils.GetClient()
	//address := db.GetContractAddress()
	//resultMap := utils.StringArrayToMap(address)
	//handler.SyncData(client,
	//	15898400, resultMap)
	////handler.CrawlData(1,274)

}
