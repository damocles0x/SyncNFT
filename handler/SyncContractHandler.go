package handler

import (
	"SyncNFT/contract"
	"SyncNFT/utils"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strings"
)

//从后往前同步数据,每笔交易只用插入一次
func SyncData(client *ethclient.Client, blockNum int64) {
	contractABI, _ := abi.JSON(strings.NewReader(contract.OracleABI))
	//15000区块之前没有nft数据
	for i := 15000; i < int(blockNum); i++ {
		query := ethereum.FilterQuery{
			Topics: [][]common.Hash{
				{
					contractABI.Events["Transfer"].ID,
					//contractABI.Events["ApprovalForAll"].ID,
				},
			},
			FromBlock: big.NewInt(blockNum - 100),
			ToBlock:   big.NewInt(blockNum),
			//Addresses: []common.Address{common.HexToAddress("0xeB9E4BEd62A82CFe17cBB90ed63d79722D9dA411")},
		}

		filterLogs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Error(err)
		}

		loopFilterLogDesc(client, filterLogs)
	}
}

func loopFilterLogDesc(client *ethclient.Client, datas []types.Log) {
	for i := len(datas) - 1; i >= 0; i-- {
		DealLogMessage(client, datas[i])
	}
}

func DealLogMessage(client *ethclient.Client, l types.Log) {
	switch l.Topics[0].String() {
	/*	case "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925":
		//Approval*/
	case "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef":
		//Transfer
		var (
			from    = "0x0"
			to      = "0x0"
			tokenId = "0"
		)
		//判断Topic的长度
		switch len(l.Topics) {
		case 4:
			//from = topic【1】
			//to = topic【2】
			//tokenID = topic【3】
			from = strings.ToLower(common.HexToAddress(l.Topics[1].Hex()).String())
			to = strings.ToLower(common.HexToAddress(l.Topics[2].Hex()).String())
			tokenId = utils.ParsingUint256(l.Topics[3].Hex())
			break
		case 3:
			//from = topic【1】
			//to = topic【2】
			//tokenID = data[0:64]
			from = strings.ToLower(common.HexToAddress(l.Topics[1].Hex()).String())
			to = strings.ToLower(common.HexToAddress(l.Topics[2].Hex()).String())
			if len(l.Data) == 64 {
				tokenId = utils.ParsingUint256(hex.EncodeToString(l.Data[0:64]))
			} else {
				log.Error("Failed to parse log data", l.Data)
			}

			break
		case 1:
			//默认按照前【0：64】为from，【64：128】为to,[128:192]
			if len(l.Data) == 192 {
				from = hex.EncodeToString(l.Data)[0:64]
				to = hex.EncodeToString(l.Data)[64:128]
				tokenId = hex.EncodeToString(l.Data)[128:192]
			} else {
				log.Error("Failed to parse log data", l.Data)
			}
		}

		fmt.Println(from, to, tokenId)

		/*	case "0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31":
			//ApprovalForAll*/

	}
}
