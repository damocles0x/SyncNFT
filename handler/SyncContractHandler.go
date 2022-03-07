package handler

import (
	"SyncNFT/contract"
	"SyncNFT/utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strings"
	"sync"
)

//从后往前同步数据
func SyncData(client *ethclient.Client, blockNum int64, mutex *sync.Mutex) {
	contractABI, _ := abi.JSON(strings.NewReader(contract.OracleABI))

	//15000区块之前没有nft数据
	for i := 15000; i < int(blockNum); i++ {
		query := ethereum.FilterQuery{
			Topics: [][]common.Hash{
				{
					contractABI.Events["Transfer"].ID,
					contractABI.Events["ApprovalForAll"].ID,
				},
			},
			FromBlock: big.NewInt(blockNum - 100),
			ToBlock:   big.NewInt(blockNum),
		}

		filterLogs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Error(err)
		}

		loopFilterLogDesc(client, filterLogs, mutex)
	}
}

func loopFilterLogDesc(client *ethclient.Client, datas []types.Log, mutex *sync.Mutex) {
	for i := len(datas) - 1; i > 0; i-- {
		DealLogMessage(client, datas[i])
	}
}

func DealLogMessage(client *ethclient.Client, l types.Log) {
	switch l.Topics[0].String() {
	//case "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925":
	//	//Approval
	//	var (
	//		i = uint256.NewInt(0)
	//	)
	//	if len(l.Topics) == 4 {
	//		//tokenID= topic【3】
	//		u := parsingUint256(l.Topics[3].Hex(), l.TxHash.String())
	//		if u != nil {
	//			i = u
	//		}
	//	} else {
	//		//tokenID = l.data
	//		toString := hex.EncodeToString(l.Data)
	//		u := parsingUint256(toString, l.TxHash.String())
	//		if u != nil {
	//			i = u
	//		}
	//	}
	//	data := ES.EsEnity{
	//		ID:            StringTohash(strings.ToLower(l.Address.String()) + i.ToBig().String()),
	//		TokenId:       i.ToBig().String(),
	//		Owner:         strings.ToLower(common.HexToAddress(l.Topics[1].String()).String()),
	//		OracleAddr:    strings.ToLower(l.Address.String()),
	//		TokenApproval: strings.ToLower(common.HexToAddress(l.Topics[2].String()).String()),
	//	}
	//	//fmt.Println("Approval", l.BlockNumber)
	//	ES.UpdateData(&data)

	case "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef":
		//Transfer
		//先判断Topic的长度标准为4 ，from to tokenId
		var (
			from    = "0x0"
			to      = "0x0"
			tokenId = "0"
		)
		if len(l.Topics) == 4 {
			//tokenID = topic【3】

			from = strings.ToLower(common.HexToAddress(l.Topics[1].Hex()).String())
			to = strings.ToLower(common.HexToAddress(l.Topics[2].Hex()).String())
			tokenId = utils.ParsingUint256(l.Topics[3].Hex())
		}

		fmt.Println(from, to, tokenId)

		//var (
		//	i   = uint256.NewInt(0)
		//	uri = "Undefined"
		//)
		//
		//
		//if len(l.Topics) == 4 {
		//	//tokenID= topic【3】
		//	u := parsingUint256(l.Topics[3].Hex(), l.TxHash.String())
		//	if u != nil {
		//		i = u
		//	}
		//} else {
		//	//tokenID = l.data
		//	toString := hex.EncodeToString(l.Data)
		//	u := parsingUint256(toString, l.TxHash.String())
		//	if u != nil {
		//		i = u
		//	}
		//}
		//
		//uri = getTokenUrI(client, l.Address.String(), i.ToBig())
		//data := ES.EsEnity{
		//	ID:         StringTohash(strings.ToLower(l.Address.String()) + i.ToBig().String()),
		//	TokenId:    i.ToBig().String(),
		//	TokenUri:   uri,
		//	Owner:      strings.ToLower(common.HexToAddress(l.Topics[2].Hex()).String()),
		//	OracleAddr: strings.ToLower(l.Address.String()),
		//}
		//fmt.Println("Transfer", l.BlockNumber)
		//ES.SaveOrUpdateData(&data)
		//case "0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31":
		//	//ApprovalForAll
		//	i := uint256.NewInt(0)
		//	if len(l.Topics) == 4 {
		//		//tokenID= topic【3】
		//		u := parsingUint256(l.Topics[3].Hex(), l.TxHash.String())
		//		if u != nil {
		//			i = u
		//		}
		//	} else {
		//		//tokenID = l.data
		//		toString := hex.EncodeToString(l.Data)
		//		u := parsingUint256(toString, l.TxHash.String())
		//		if u != nil {
		//			i = u
		//		}
		//	}
		//	data := db.ORACLE_DATA{
		//		Address:     strings.ToLower(l.Address.String()),
		//		ApprovalAll: strings.ToLower(common.HexToAddress(l.Topics[2].String()).String()),
		//	}
		//	//fmt.Println("ApprovalForAll", l.BlockNumber)
		//	go db.UpdateOracleApprove(&data, i.ToBig().String())
	}
}
