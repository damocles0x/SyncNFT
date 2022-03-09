package db

import (
	"SyncNFT/config"
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
