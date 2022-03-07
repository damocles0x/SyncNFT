package main

import (
	"SyncNFT/log"
	"SyncNFT/utils"
	"fmt"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./errorLog", "log", time.Hour*24*14, time.Hour*24)
	//15844391 27073
	//mutex := sync.Mutex{}
	//client := utils.GetClient(&mutex)
	//handler.SyncData(client,15844391,&mutex)https://bscscan.com/tx/0x92abcb7c1bc67a1a82ecdd46c83eaafb17c122558396adb993d360adc30f0801#eventlog
	//query := ethereum.FilterQuery{
	//	FromBlock: big.NewInt(
	//		17349509),
	//	ToBlock: big.NewInt(
	//		17349509),
	//	Addresses:  []common.Address{common.HexToAddress("0xeB9E4BEd62A82CFe17cBB90ed63d79722D9dA411")},
	//}
	//
	//client, _ := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545")
	////client, _ := ethclient.Dial("https://bsc-dataseed.binance.org/")
	//filterLogs, _ := client.FilterLogs(context.Background(), query)
	//
	//
	//for _,log := range filterLogs {
	//	if log.Topics[0].String() =="0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"{
	//		//fmt.Println(log.Topics[0].String())
	//		//fmt.Println(common.HexToAddress(log.Topics[1].String()))
	//		//fmt.Println(common.HexToAddress(log.Topics[2].String()))
	//		//fmt.Println(log.Topics[3].String())
	//		fmt.Println(common.Bytes2Hex(log.Data[64:128]))
	//		common.BytesToAddress()
	//	}
	//}

	//parseUint, err := strconv.ParseUint("0x000000000000000000000000000000000000000000000000000000000000c9e2", 16, 32)

	//uint256 := utils.ParsingUint256("0x000000000000000000000000000000000000000000000000000000000000c9e2")
	parseUint := utils.ParsingUint256("0x0000000000000000000000000000000000000000000000000000000000000000")

	fmt.Print(parseUint)
}
