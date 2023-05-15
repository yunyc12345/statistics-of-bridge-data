package task

import (
	"github.com/robfig/cron/v3"
	"sync"
	"time"
	"yunyc12345/statistics-of-bridge-data/handle"
	"yunyc12345/statistics-of-bridge-data/utils"
)

func StatMsgSenderTask(config utils.Config, crontab *cron.Cron) error {
	spec := "0 0 8 * * *"
	_, err := crontab.AddFunc(spec, func() {
		MsgSenderTask(config, time.Now(), nil)
	})
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	return nil
}

func MsgSenderTask(config utils.Config, now time.Time, req *utils.RunParams) error {
	csvRootPath := config.Server.CsvFilePath
	//now := time.Now()

	utils.Logger.Infof("stat msg task start: %v", now.String())

	yd := now.AddDate(0, 0, -1)
	ydDataStr := yd.Format("2006-01-02")
	ydFileDataStr := yd.Format("2006.01.02")
	heightMap, err := utils.GetPrevChainEndHeight(csvRootPath, ydDataStr)
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	nowDataStr := now.Format("2006-01-02")
	nowFileDataStr := now.Format("2006.01.02")
	lists := make([]*sync.Map, 0)
	newChainList := make([]utils.Chain, 0)

	for _, chain := range config.MainnetMsgSendersStatistics {
		ydEndHeight := heightMap[chain.Name]
		ydCsvFileName := utils.CreateMsgCsvFileName(chain.Name, ydFileDataStr, ydEndHeight)
		ydCsvFilePath := utils.CreateCsvPath(ydCsvFileName, csvRootPath, ydDataStr, "stat_msg_sender")
		list, ll, err := utils.GetYesterdayCsvData(ydCsvFilePath)
		if err != nil {
			utils.Logger.Errorf("chain: %v, get yesterday csv: %v, error: %v", chain.Name, ydCsvFilePath, err)
		}

		utils.Logger.Infof("chain: %v, msg list len: %v", chain.Name, ll)

		endHeight, err := utils.GetChainCurHeight(chain)
		if err != nil {
			utils.Logger.Errorf("chain: %v, get cur height error", chain.Name)
		}
		if req != nil {
			switch chain.Name {
			case req.Bsc.ChainName:
				endHeight = req.Bsc.EndHeight
			case req.Eth.ChainName:
				endHeight = req.Eth.EndHeight
			}
		}

		c := chain.ToMsgNew(ydEndHeight, endHeight)
		newChainList = append(newChainList, c)
		lists = append(lists, list)
	}

	w := &sync.WaitGroup{}
	handle.TriggerMsgStat(newChainList, w, csvRootPath, nowDataStr, nowFileDataStr, lists)
	time.Sleep(5 * time.Second)
	w.Wait()
	utils.Logger.Infof("stat msg sender task over, time-consuming: %v", time.Now().Sub(now).String())

	return nil
}

func StatNftMinterTask(config utils.Config, crontab *cron.Cron) error {

	spec := "0 0 8 * * *"
	_, err := crontab.AddFunc(spec, func() {
		NftMinterTask(config, time.Now(), nil)
	})
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	return nil
}

func NftMinterTask(config utils.Config, now time.Time, req *utils.RunParams) error {
	csvRootPath := config.Server.CsvFilePath
	//now := time.Now()

	utils.Logger.Infof("stat nft minter task start: %v", now.String())

	yd := now.AddDate(0, 0, -1)
	ydDataStr := yd.Format("2006-01-02")
	ydFileDataStr := yd.Format("2006.01.02")

	heightMap, err := utils.GetPrevChainTokensEndHeight(csvRootPath, ydDataStr)
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	nowDataStr := now.Format("2006-01-02")
	nowFileDataStr := now.Format("2006.01.02")
	lists := make(map[string]*sync.Map)
	newChainList := make([]utils.Chain, 0)

	for _, chain := range config.MainnetNftMinterStatistics {
		ts := make([]utils.Token, 0)
		for _, token := range chain.Tokens {
			ydEndHeight := heightMap[chain.Name+"-"+token.Name]
			ydCsvFileName := utils.CreateNftCsvFileName(chain.Name, ydFileDataStr, token.Name, ydEndHeight)
			ydCsvFilePath := utils.CreateCsvPath(ydCsvFileName, csvRootPath, ydDataStr, "stat_nft_minter")
			list, ll, err := utils.GetYesterdayCsvData(ydCsvFilePath)
			if err != nil {
				utils.Logger.Errorf("chain: %v, get yesterday csv: %v, error: %v", chain.Name, ydCsvFilePath, err)
			}

			utils.Logger.Infof("chain: %v, token: %v, nft minter list len: %v", chain.Name, token.Name, ll)

			endHeight, err := utils.GetChainCurHeight(chain)
			if err != nil {
				utils.Logger.Errorf("chain: %v, get cur height error", chain.Name)
			}

			if req != nil {
				switch chain.Name {
				case req.Bsc.ChainName:
					endHeight = req.Bsc.EndHeight
				case req.Eth.ChainName:
					endHeight = req.Eth.EndHeight
				}
			}

			ts = append(ts, token.NewToken(ydEndHeight, endHeight))
			lists[chain.Name+"-"+token.Name] = list
		}
		c := chain.ToTokenNew(ts)
		newChainList = append(newChainList, c)
	}

	w := &sync.WaitGroup{}
	handle.TriggerNftStat(newChainList, w, csvRootPath, nowDataStr, nowFileDataStr, lists)
	time.Sleep(5 * time.Second)
	w.Wait()
	utils.Logger.Infof("stat nft minter task over, time-consuming: %v", time.Now().Sub(now).String())
	return nil
}

func StatNftBridgeSenderTask(config utils.Config, crontab *cron.Cron) error {
	spec := "0 0 8 * * *"
	_, err := crontab.AddFunc(spec, func() {
		NftBridgeSenderTask(config, time.Now(), nil)
	})
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	return nil
}

func NftBridgeSenderTask(config utils.Config, now time.Time, req *utils.RunParams) error {
	csvRootPath := config.Server.CsvFilePath
	//now := time.Now()

	utils.Logger.Infof("stat nft minter task start: %v", now.String())

	yd := now.AddDate(0, 0, -1)
	ydDataStr := yd.Format("2006-01-02")
	ydFileDataStr := yd.Format("2006.01.02")

	heightMap, err := utils.GetPrevChainTokensEndHeight(csvRootPath, ydDataStr)
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	nowDataStr := now.Format("2006-01-02")
	nowFileDataStr := now.Format("2006.01.02")
	lists := make(map[string]*sync.Map)
	newChainList := make([]utils.Chain, 0)

	for _, chain := range config.MainnetNftbridgeSendersStatistics {
		ts := make([]utils.Token, 0)
		for _, token := range chain.Tokens {
			mapKey := chain.Name + "-" + token.Name
			utils.Logger.Infof("chain: %v, token: %v, mapkey: %v", chain.Name, token.Name, mapKey)
			ydEndHeight := heightMap[mapKey]
			ydCsvFileName := utils.CreateNftBridgeCsvFileName(chain.Name, ydFileDataStr, token.Name, ydEndHeight)
			ydCsvFilePath := utils.CreateCsvPath(ydCsvFileName, csvRootPath, ydDataStr, "stat_nftbridge_sender")
			list, ll, err := utils.GetYesterdayCsvData(ydCsvFilePath)
			if err != nil {
				utils.Logger.Errorf("chain: %v, get yesterday csv: %v, error: %v", chain.Name, ydCsvFilePath, err)
			}

			utils.Logger.Infof("chain: %v, token: %v, nft bridge list len: %v", chain.Name, token.Name, ll)

			endHeight, err := utils.GetChainCurHeight(chain)
			if err != nil {
				utils.Logger.Errorf("chain: %v, get cur height error", chain.Name)
			}

			if req != nil {
				switch chain.Name {
				case req.Bsc.ChainName:
					endHeight = req.Bsc.EndHeight
				case req.Eth.ChainName:
					endHeight = req.Eth.EndHeight
				}
			}

			ts = append(ts, token.NewToken(ydEndHeight, endHeight))
			lists[mapKey] = list
		}
		c := chain.ToTokenNew(ts)
		newChainList = append(newChainList, c)
	}

	w := &sync.WaitGroup{}
	handle.TriggerNftBridgeStat(newChainList, w, csvRootPath, nowDataStr, nowFileDataStr, lists)
	time.Sleep(5 * time.Second)
	w.Wait()
	utils.Logger.Infof("stat nft bridge sender task over, time-consuming: %v", time.Now().Sub(now).String())
	return nil
}
