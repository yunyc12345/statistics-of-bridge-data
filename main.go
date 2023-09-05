package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
	"yunyc12345/statistics-of-bridge-data/task"
	"yunyc12345/statistics-of-bridge-data/utils"
)

func main() {
	utils.InitLogger(&utils.LogConf{
		Level: "info",
		Path:  "simple2/app.log",
	})
	wg := &sync.WaitGroup{}

	//wg.Add(1)
	//go eth_alpha_mint_stat(wg)
	//
	//wg.Add(1)
	//go bsc_alpha_mint_stat(wg)
	//
	//wg.Wait()
	//utils.Logger.Info("alpha stat over")

	//
	//wg.Wait()
	//utils.Logger.Info("stat msg over")

	//wg.Add(1)
	//go eth_stat_nftbridge_sender(wg)
	//
	//wg.Add(1)
	//go bsc_stat_nftbridge_sender(wg)
	//
	//wg.Wait()
	//utils.Logger.Info("stat nftbridge sender over")

	//wg.Add(1)
	//go eth_stat_nftbridge_sender_for_receiver_chain(wg)
	//
	//wg.Add(1)
	//go bsc_stat_nftbridge_sender_for_receiver_chain(wg)
	//
	//wg.Wait()
	//utils.Logger.Info("stat nftbridge for receiver chain sender over")

	//wg.Add(1)
	//go polygon_stat_nftbridge_sender(wg)

	//wg.Add(3)
	//go polygon_alpha_mint_stat(wg)

	//go polygon_zkLightClient_mint_stat(wg)
	//
	//go eth_zkLightClient_mint_stat(wg)
	//
	//go bsc_zkLightClient_mint_stat(wg)

	//go polygon_stat_nftbridge_sender_zklightclient(wg)
	//
	//go eth_stat_nftbridge_sender_zklightclient(wg)
	//
	//go bsc_stat_nftbridge_sender_zklightclient(wg)

	//go bsc_faucet_mint_stat(wg)

	//go eth_stat_msg(wg)
	//go bsc_stat_msg(wg)
	//go polygon_stat_msg(wg)

	//wg.Add(2)
	//go bsc_stat_nftbridge_sender_lubun_loyalty(wg)
	//
	//go bsc_luban_loyalty_mint_stat(wg)

	//go coredao_alpha_mint_stat(wg)
	//
	//go coredao_stat_nftbridge_sender_alpha(wg)

	//wg.Add(4)

	//go polygon_opbnb_mint_stat(wg)
	//
	//go coredao_opbnb_mint_stat(wg)
	//
	//go eth_opbnb_mint_stat(wg)
	//
	//go bsc_opbnb_mint_stat(wg)

	//go eth_stat_nftbridge_sender_for_receiver_opbnb_chain(wg)
	//
	//go bsc_stat_nftbridge_sender_for_receiver_opbnb_chain(wg)
	//wg.Add(5)
	//go polygon_stat_nftbridge_sender_for_receiver_opbnb_chain(wg)
	//
	//go coredao_stat_nftbridge_sender_for_receiver_opbnb_chain(wg)

	//go coredao_stat_msg(wg)

	//go celo_stat_msg(wg)

	//go bsc_stat_msg_and_hist(wg)
	//
	//go eth_stat_msg_and_hist(wg)
	//
	//go polygon_stat_msg_and_hist(wg)
	//
	//go coredao_stat_msg_and_hist(wg)
	//
	//go celo_stat_msg_and_hist(wg)
	//wg.Add(1)
	//go coredao_stat_nftbridge_sender_alpha(wg)

	//wg.Add(4)
	//
	//go l0_coredao_stat_nftbridge_sender_for_receiver_opbnb_chain(wg)
	//
	//go l0_bsc_stat_nftbridge_sender_for_receiver_opbnb_chain(wg)
	//
	//go l0_polygon_stat_nftbridge_sender_for_receiver_opbnb_chain(wg)
	//
	//go l0_celo_stat_nftbridge_sender_for_receiver_opbnb_chain(wg)
	//wg.Add(7)
	//
	//go l0_bsc_stat_nftbridge_sender_alpha(wg)
	//go l0_coredao_stat_nftbridge_sender_alpha(wg)
	//go l0_eth_stat_nftbridge_sender_alpha(wg)
	//go l0_polygon_stat_nftbridge_sender_alpha(wg)
	//
	//go l0_bsc_stat_nftbridge_sender_zklightclient(wg)
	//go l0_eth_stat_nftbridge_sender_zklightclient(wg)
	//go l0_polygon_stat_nftbridge_sender_zklightclient(wg)

	//wg.Add(1)
	//go l0_bsc_stat_nftbridge_sender_lubun_loyalty(wg)

	//wg.Add(1)
	//go coredao_alpha_mint_stat(wg)

	wg.Add(1)
	//go l0_mantle_stat_nftbridge_receiver_nft(wg)
	//go bsc_stat_nftbridge_sender_sbt(wg)
	//go combo_to_opbnb_claim_origin(wg)
	go combo_to_opbnb_claim2(wg)

	wg.Wait()
	utils.Logger.Info("stat nftbridge for receiver chain sender over")
}

func main3242355() {

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

	//b, err := utils.JudgHaveHisData(config.Server.CsvFilePath)
	//if err != nil {
	//	utils.Logger.Info("init 2023-05-11 data not exist")
	//}
	//
	//if !b {
	//	utils.Logger.Info("There is no initial data for 2023-05-11 00:00. Initialization process will begin.")
	//	w := &sync.WaitGroup{}
	//
	//	go handle.TriggerNftStat(config.MainnetNftMinterStatistics, w, config.Server.CsvFilePath, "2023-05-11", "2023.05.11", nil)
	//
	//	go handle.TriggerNftBridgeStat(config.MainnetNftbridgeSendersStatistics, w, config.Server.CsvFilePath, "2023-05-11", "2023.05.11", nil)
	//
	//	go handle.TriggerMsgStat(config.MainnetMsgSendersStatistics, w, config.Server.CsvFilePath, "2023-05-11", "2023.05.11", nil, false)
	//
	//	time.Sleep(time.Second * 20)
	//	w.Wait()
	//	err := utils.ZipStat("2023-05-11", config.Server.CsvFilePath)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	utils.Logger.Info("init handle over")
	//}

	utils.Logger.Info("start server")

	crontab := cron.New(cron.WithSeconds())

	task.StatMsgSenderTask(config, crontab)

	task.StatNftMinterTask(config, crontab)

	task.StatNftBridgeSenderTask(config, crontab)

	task.StatDomainAddressTask(config, crontab)

	task.StatNftBridgeSenderForReceiverChainTask(config, crontab)

	task.ReceiverNftTask(config, crontab)

	task.StatClaimNftSender(config, crontab)

	crontab.Start()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello")
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
			go task.DomainAddressTask(config, p)
			go task.NftBridgeSenderForReceiverChainTask(config, p, &req)
			go task.StatReceiverNftTask(config, p, &req)
			go task.StatClaimNftSenderTask(config, p, &req)
		}

		return
	})

	r.Run(config.Server.Port) // listen and serve on 0.0.0.0:8080

}
