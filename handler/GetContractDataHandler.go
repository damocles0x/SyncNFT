package handler

import (
	"SyncNFT/db"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//get 721Token message by bsc
func CrawlData(from int64, page int64) {
	for i := from; i <= page; i++ {
		url := "https://bscscan.com/tokens-nft?ps=100&p=" + strconv.Itoa(int(i))
		get, err := http.Get(url)
		if err != nil {
			log.Error(err)
			i -= 1
			continue
		}
		r(get.Body, i)
	}
}

func r(r io.Reader, i int64) {
	datas := []db.CONTRACT{}
	reader, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Error(err)
	}
	reader.Find(".text-primary").Each(func(i int, s *goquery.Selection) {
		//var symbol string
		attr, _ := s.Attr("href")
		if attr != "https://etherscan.io/" && s.Text() != "Etherscan" {
			attr = attr[7:]
			data := db.CONTRACT{
				ContractAddress: strings.ToLower(attr),
				ContractName:    s.Text(),
				CreatedTime:     time.Now(),
				UpdatedTime:     time.Now(),
			}
			datas = append(datas, data)
		}
	})
	db.SaveContracts(&datas)
	fmt.Println(i)
}
