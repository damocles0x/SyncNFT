package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
	"os"
	"path"
	"strconv"
	"sync"
)

var (
	APPVIPER *viper.Viper
	DB       *gorm.DB
	EsCli    *elastic.Client
	Mutex    *sync.Mutex
	Redis    *redis.Client
)

func init() {
	APPVIPER = initAppConfig()
	DB = initDB()
	Mutex = initMutex()
	//EsCli = initElasticSearch()
	Redis = initRedis()
}

func initAppConfig() *viper.Viper {
	workDir, _ := os.Getwd()
	appViper := viper.New()
	appViper.SetConfigName("application")
	appViper.SetConfigType("yml")
	appViper.AddConfigPath(path.Join(workDir, "config"))
	err := appViper.ReadInConfig()
	if err != nil {

	}
	return appViper
}
func initMutex() *sync.Mutex {
	mutex := sync.Mutex{}
	return &mutex
}
func initDB() *gorm.DB {
	host := APPVIPER.GetString("database.host")
	port := APPVIPER.GetString("database.port")
	database := APPVIPER.GetString("database.databaseName")
	username := APPVIPER.GetString("database.username")
	password := APPVIPER.GetString("database.password")
	charset := APPVIPER.GetString("database.charset")
	loc := APPVIPER.GetString("database.loc")

	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)

	db, err := gorm.Open(mysql.Open(sqlStr), &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)})
	if err != nil {
		panic("connected error" + err.Error())
	} else {
		log.Infoln("connect db success")
	}
	return db
}

func initClient() *ethclient.Client {
	url := APPVIPER.GetString("node.url1")
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Error("client faild:", err)
	}

	log.Infoln("connect client success")
	return client
}

func initClients() []*ethclient.Client {
	var clients []*ethclient.Client
	for i := 1; i < 14; i++ {
		url := APPVIPER.GetString("node.url" + strconv.Itoa(i))
		client, err := ethclient.Dial(url)
		if err != nil {
			log.Error("client faild:", i, err)
		} else {
			clients = append(clients, client)
		}
	}
	log.Infoln("connect client success")
	return clients
}

func initElasticSearch() *elastic.Client {
	url := APPVIPER.GetString("elasticSearch.host")
	cli, err := elastic.NewSimpleClient(
		elastic.SetURL(url),
	)
	if err != nil {
		panic("new es client failed, err=" + err.Error())
	}
	return cli
}

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     APPVIPER.GetString("redis.host") + ":" + APPVIPER.GetString("redis.port"),
		Password: APPVIPER.GetString("redis.password"), // no password set
		DB:       APPVIPER.GetInt("redis.db"),          // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Error(err)
	}
	return rdb
}
