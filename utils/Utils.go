package utils

import (
	"SyncNFT/config"
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
)

func DataToJson(data interface{}) []byte {
	bytesData, err := json.Marshal(data)
	if err != nil {
		log.Error("DataToJson:", err)
	}
	return bytesData
}

func JsonToData(body []byte, entity interface{}) {
	err := json.Unmarshal(body, entity)
	if err != nil {
		log.Error(err)
	}
}
func StringArrayToMap(stringArray *[]string) map[string]byte {
	var result map[string]byte
	result = make(map[string]byte)
	for _, s := range *stringArray {
		result[s] = byte(1)
	}
	return result
}

func ParsingUint256(s string) string {
	var result string
	if s == "0x0000000000000000000000000000000000000000000000000000000000000000" {
		result = "0x0"
	} else {
		//除去0x后面多余的0
		compile := regexp.MustCompile("^[0]+")
		findString := compile.FindString(s[2:])
		result = "0x" + s[len(findString)+2:]
	}
	fromHex, err := uint256.FromHex(result)
	if err != nil {
		log.Error(err)
		fromHex = uint256.NewInt(0)
	}
	return fromHex.ToBig().String()
}

/**
获取客户端链接
*/
func GetClient() *ethclient.Client {
	url := config.APPVIPER.GetString("nodes.node1")
	length := config.APPVIPER.GetInt("nodes.length")
	//调用次数
	y := 0
	//为了保证多个协程调用时url不会重复，要对y加锁
	config.Mutex.Lock()
	y++
	config.Mutex.Unlock()
	url = config.APPVIPER.GetString("nodes.node" + strconv.Itoa(y%length))
	dial, err := ethclient.Dial(url)
	if err != nil {
		log.Error(err)
	}

	return dial
}
