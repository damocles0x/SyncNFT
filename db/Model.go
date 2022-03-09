package db

import "time"

type NFT struct {
	Id            int64     `gorm:"id"` // id
	CreatedTime   time.Time `gorm:"created_time"`
	UpdatedTime   time.Time `gorm:"updated_time"`
	TxFromAddress string    `gorm:"tx_from_address"` // 发送者
	TxToAddress   string    `gorm:"tx_to_address"`   // 接收者
	TxHash        string    `gorm:"tx_hash"`         // 交易hash
	TokenId       string    `gorm:"token_id"`        // tokenId
	ContractId    string    `gorm:"contract_id"`     // contract表Id
}

func (NFT) TableName() string {
	return "nft"
}

type CONTRACT struct {
	Id              int64     `gorm:"id"` // id
	CreatedTime     time.Time `gorm:"created_time"`
	UpdatedTime     time.Time `gorm:"updated_time"`
	TxFromAddress   string    `gorm:"tx_from_address"`  // 创建人
	TxHash          string    `gorm:"tx_hash"`          // 交易hash
	ContractAddress string    `gorm:"contract_address"` // 合约地址
	ContractName    string    `gorm:"contract_name"`    // 合约名称
	ContractSymbol  string    `gorm:"contract_symbol"`  // 合约symbol
}

func (CONTRACT) TableName() string {
	return "contract"
}
