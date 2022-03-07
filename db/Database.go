package db

import (
	"SyncNFT/config"
	log "github.com/sirupsen/logrus"
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
