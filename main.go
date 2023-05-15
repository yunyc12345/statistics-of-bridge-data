package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"yunyc12345/statistics-of-bridge-data/contracts"
	"yunyc12345/statistics-of-bridge-data/handle"
	"yunyc12345/statistics-of-bridge-data/task"
	"yunyc12345/statistics-of-bridge-data/utils"
)

func main2() {
	utils.InitLogger(&utils.LogConf{
		Level: "info",
		Path:  "log/app.log",
	})

	c, err := ethclient.Dial("https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e")
	if err != nil {
		panic("")
	}
	abi, err := contracts.NftBridgeMetaData.GetAbi()
	if err != nil {
		panic("")
		return
	}

	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-05-15/stat_nftbridge_sender/bsc mainnet-cobee-stat_sender-28207389-2023.05.15")
	if err != nil {
		panic("")
	}

	event := abi.Events["TransferNFT"]
	curHeight, nextHeight := uint64(27532007), uint64(27532007)
	internal := uint64(10000)
	for curHeight < 27815520 {
		if curHeight+internal >= 27815520 {
			nextHeight = 27815520
		} else {
			nextHeight += internal
		}

		fmt.Printf("chain: bsc, height: %v ~ %v\n", curHeight, nextHeight)

		query := ethereum.FilterQuery{
			BlockHash: nil,
			FromBlock: big.NewInt(int64(curHeight)), // The error will occur if logs cannot be pulled from the latest block, with a difference of 32.
			ToBlock:   big.NewInt(int64(nextHeight)),
			Addresses: []common.Address{common.HexToAddress("0xE09828f0DA805523878Be66EA2a70240d312001e")},
			Topics:    [][]common.Hash{{event.ID}},
		}
		ctx := context.Background()
		try := 0
		logs, err := c.FilterLogs(ctx, query)
		for err != nil && try <= 10 {
			try++
			fmt.Printf("chain: bsc, try: %v, get logs error :%v\n", try, err)
			time.Sleep(5 * time.Second)
			logs, err = c.FilterLogs(ctx, query)
		}

		if err != nil || try > 10 {
			fmt.Printf("chain: bsc, try: %v, get logs error :%v\n", try, err)
		}

		fmt.Printf("chain: bsc, logs len: %v\n", len(logs))

		for _, l := range logs {
			eventData, err := abi.Unpack(event.Name, l.Data)
			if err != nil {
				panic("")
			}
			sender := strings.ToLower(eventData[3].(common.Address).Hex())
			tokenAddress := strings.ToLower(eventData[0].(common.Address).Hex())
			if tokenAddress == strings.ToLower("0x9e8c1e7b35f646a606644a5532c6103c647938cf") {
				list.Store(sender, utils.Member)
			}

		}

		time.Sleep(1 * time.Second)

		curHeight = nextHeight

	}

	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-cobee-stat_sender-28207389-2023.05.15")
	if err != nil {
		utils.Logger.Errorf("stat nft bridge sender to csv err: %v", err)
	}
}

func main() {

	utils.InitLogger(&utils.LogConf{
		Level: "info",
		Path:  "log/app.log",
	})
	cfp := os.Getenv("CONFIG_FILE_PATH")

	fmt.Printf("config path: %v\n", cfp)

	jsonFile, err := os.Open(cfp)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config utils.Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}

	err = utils.CreateCsvDir(config.Server.CsvFilePath)
	if err != nil {
		panic(err)
	}

	b, err := utils.JudgHaveHisData(config.Server.CsvFilePath)
	if err != nil {
		utils.Logger.Info("init 2023-05-11 data not exist")
	}

	if !b {
		utils.Logger.Info("There is no initial data for 2023-05-11 00:00. Initialization process will begin.")
		w := &sync.WaitGroup{}

		go handle.TriggerNftStat(config.MainnetNftMinterStatistics, w, config.Server.CsvFilePath, "2023-05-11", "2023.05.11", nil)

		go handle.TriggerNftBridgeStat(config.MainnetNftbridgeSendersStatistics, w, config.Server.CsvFilePath, "2023-05-11", "2023.05.11", nil)

		go handle.TriggerMsgStat(config.MainnetMsgSendersStatistics, w, config.Server.CsvFilePath, "2023-05-11", "2023.05.11", nil)

		time.Sleep(time.Second * 20)
		w.Wait()
		err := utils.ZipStat("2023-05-11", config.Server.CsvFilePath)
		if err != nil {
			panic(err)
		}

		utils.Logger.Info("init handle over")
	}

	utils.Logger.Info("start server")

	crontab := cron.New(cron.WithSeconds())

	task.StatMsgSenderTask(config, crontab)

	task.StatNftMinterTask(config, crontab)

	task.StatNftBridgeSenderTask(config, crontab)

	crontab.Start()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})

	r.GET("/csv/", func(c *gin.Context) {
		dirStr := time.Now().Format("2006-01-02")
		fileName := dirStr + ".zip"
		csvRootFilePath := config.Server.CsvFilePath
		err = utils.ZipStat(dirStr, csvRootFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"param": dirStr,
			})
			return
		}
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(csvRootFilePath + "/" + fileName)
		return
	})

	r.GET("/csv/:data", func(c *gin.Context) {
		dirStr := c.Param("data")
		_, err := time.Parse("2006-01-02", dirStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"param": dirStr,
			})
			return
		}
		fileName := dirStr + ".zip"
		csvRootFilePath := config.Server.CsvFilePath
		err = utils.ZipStat(dirStr, csvRootFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(csvRootFilePath + "/" + fileName)
		return
	})

	r.POST("/run/:data", func(c *gin.Context) {
		dirStr := c.Param("data")
		p, err := time.Parse("2006-01-02", dirStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"param": dirStr,
			})
			return
		}

		var req utils.RunParams
		err = c.ShouldBindJSON(&req) // 解析req参数
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		csvRootFilePath := config.Server.CsvFilePath
		exists, err := utils.PathExists(csvRootFilePath + "/" + dirStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"param": dirStr,
			})
			return
		}

		if !exists {
			go task.MsgSenderTask(config, p, &req)
			go task.NftMinterTask(config, p, &req)
			go task.NftBridgeSenderTask(config, p, &req)
		}

		return
	})

	r.Run(config.Server.Port) // listen and serve on 0.0.0.0:8080

}
