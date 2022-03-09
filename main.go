package main

import (
	"SyncNFT/handler"
	"SyncNFT/log"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./errorLog", "log", time.Hour*24*14, time.Hour*24)
	//15844391 27073
	//client := utils.GetClient()
	dial, _ := ethclient.Dial("https://data-seed-prebsc-2-s2.binance.org:8545/")
	handler.SyncData(dial,
		17349509)

	//https://bscscan.com/tx/0x92abcb7c1bc67a1a82ecdd46c83eaafb17c122558396adb993d360adc30f0801#eventlog
}
