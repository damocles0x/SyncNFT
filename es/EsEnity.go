package ES

import (
	"time"
)

type EsEnity struct {
	ID              string    `json:"id"`
	CreatedTime     time.Time `json:"created_time"`
	UpdatedTime     time.Time `json:"updated_time"`
	TxFromAddress   string    `json:"tx_from_address"`  // 发送者
	TxToAddress     string    `json:"tx_to_address"`    // 接收者
	TxHash          string    `json:"tx_hash"`          // 交易hash
	TokenId         string    `json:"token_id"`         // tokenId
	ContractAddress string    `json:"contract_address"` // contract表Id

}
