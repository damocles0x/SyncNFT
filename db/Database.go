package db

import (
	"SyncNFT/config"
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func SaveContracts(data *[]CONTRACT) {
	res := config.DB.Create(data)
	if res.Error != nil {
		log.Error(res.Error)
	}
}

func GetContractAddress() *[]string {
	var addressArray []string
	res := config.DB.Table("contract").Select("contract_address").Find(&addressArray)
	if res.Error != nil {
		log.Error(res.Error)
	}
	return &addressArray
}

/**
处理同步数据时插入nft的逻辑
*/
func InsertNFT(nft *NFT) {
	//判断transfer里面的合约地址是否为初始化那一部分的nft合约地址
	enity := NFT{}
	res := config.DB.Table("nft").Select("id").Where("token_id = ? and contract_address = ?", nft.TokenId, nft.ContractAddress).Find(&enity)
	if res.RowsAffected == 0 {
		nft.CreatedTime = time.Now()
		nft.UpdatedTime = time.Now()
		config.DB.Create(nft)
	}
}

func InsertNFTBatch(nfts *[]NFT) error {
	var buffer bytes.Buffer
	sql := "INSERT INTO `nft` (`created_time`,`updated_time`,`tx_from_address`,`tx_to_address`,`tx_hash`,`token_id`,`contract_address`) VALUES"
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	time := time.Now().Format("2006-01-02 15:04:05")
	for i, nft := range *nfts {
		if i == len(*nfts)-1 {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s','%s');", time, time, nft.TxFromAddress, nft.TxToAddress, nft.TxHash, nft.TokenId, nft.ContractAddress))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s','%s'),", time, time, nft.TxFromAddress, nft.TxToAddress, nft.TxHash, nft.TokenId, nft.ContractAddress))
		}
	}
	return config.DB.Exec(buffer.String()).Error
}
