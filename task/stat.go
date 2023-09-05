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
	day := now.Day()
	//ydMonth := yd.Month()
	//nowMonth := now.Month()
	diffMonth := false
	if day == 2 {
		diffMonth = true
	}

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
			case req.Polygon.ChainName:
				endHeight = req.Polygon.EndHeight
			case req.Celo.ChainName:
				endHeight = req.Celo.EndHeight
			}

		}

		c := chain.ToNewChain(ydEndHeight, endHeight)
		newChainList = append(newChainList, c)
		lists = append(lists, list)
	}

	w := &sync.WaitGroup{}
	handle.TriggerMsgStat(newChainList, w, csvRootPath, nowDataStr, nowFileDataStr, lists, diffMonth)
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

	heightMap, err := utils.GetPrevStatNftEndHeight(csvRootPath, ydDataStr)
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
				case req.Polygon.ChainName:
					endHeight = req.Polygon.EndHeight
				case req.CoreDao.ChainName:
					endHeight = req.CoreDao.EndHeight
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
				case req.Polygon.ChainName:
					endHeight = req.Polygon.EndHeight
				case req.CoreDao.ChainName:
					endHeight = req.CoreDao.EndHeight
				case req.L0Bsc.ChainName:
					endHeight = req.L0Bsc.EndHeight
				case req.L0CoreDao.ChainName:
					endHeight = req.L0CoreDao.EndHeight
				case req.L0Polygon.ChainName:
					endHeight = req.L0Polygon.EndHeight
				case req.L0Eth.ChainName:
					endHeight = req.L0Eth.EndHeight
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

func StatNftBridgeSenderForReceiverChainTask(config utils.Config, crontab *cron.Cron) error {
	spec := "0 0 8 * * *"
	_, err := crontab.AddFunc(spec, func() {
		NftBridgeSenderForReceiverChainTask(config, time.Now(), nil)
	})
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	return nil
}

func NftBridgeSenderForReceiverChainTask(config utils.Config, now time.Time, req *utils.RunParams) error {
	csvRootPath := config.Server.CsvFilePath
	//now := time.Now()

	utils.Logger.Infof("stat nft bridge sender for receiver task start: %v", now.String())

	yd := now.AddDate(0, 0, -1)
	ydDataStr := yd.Format("2006-01-02")
	ydFileDataStr := yd.Format("2006.01.02")

	heightMap, err := utils.GerPrevChainReceiverEndHeight(csvRootPath, ydDataStr)
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	nowDataStr := now.Format("2006-01-02")
	nowFileDataStr := now.Format("2006.01.02")
	lists := make(map[string]*sync.Map)
	newChainList := make([]utils.Chain, 0)

	for _, chain := range config.StatReceiverChains {
		scs := make([]utils.StatChain, 0)

		for _, sc := range chain.StatChains {

			if len(sc.StatContract) == 0 {
				mapKey := chain.Name + "-" + sc.Name + "-" + "all"
				utils.Logger.Infof("chain: %v, receiver chain: %v, mapkey: %v", chain.Name, sc.Name, mapKey)
				ydEndHeight := heightMap[mapKey]
				ydCsvFileName := utils.CreateNftBridgeForReceiverCsvFileName(chain.Name, ydFileDataStr, sc.Name, ydEndHeight, "all")
				ydCsvFilePath := utils.CreateCsvPath(ydCsvFileName, csvRootPath, ydDataStr, "stat_nftbridge_receiver_chain_sender")
				list, ll, err := utils.GetYesterdayCsvData(ydCsvFilePath)
				if err != nil {
					utils.Logger.Errorf("chain: %v, get yesterday csv: %v, error: %v", chain.Name, ydCsvFilePath, err)
				}

				utils.Logger.Infof("chain: %v, receiver chain: %v, nft bridge list len: %v", chain.Name, sc.Name, ll)

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
					case req.Polygon.ChainName:
						endHeight = req.Polygon.EndHeight
					case req.CoreDao.ChainName:
						endHeight = req.CoreDao.EndHeight
					}
				}

				scs = append(scs, sc.ToNewStatChain(ydEndHeight, endHeight))
				lists[mapKey] = list
			} else {

				for _, s := range sc.StatContract {
					mapKey := chain.Name + "-" + sc.Name + "-" + s.Name
					utils.Logger.Infof("chain: %v, receiver chain: %v, mapkey: %v", chain.Name, sc.Name, mapKey)
					ydEndHeight := heightMap[mapKey]
					ydCsvFileName := utils.CreateNftBridgeForReceiverCsvFileName(chain.Name, ydFileDataStr, sc.Name, ydEndHeight, s.Name)
					ydCsvFilePath := utils.CreateCsvPath(ydCsvFileName, csvRootPath, ydDataStr, "stat_nftbridge_receiver_chain_sender")
					list, ll, err := utils.GetYesterdayCsvData(ydCsvFilePath)
					if err != nil {
						utils.Logger.Errorf("chain: %v, get yesterday csv: %v, error: %v", chain.Name, ydCsvFilePath, err)
					}

					utils.Logger.Infof("chain: %v, receiver chain: %v, nft bridge list len: %v", chain.Name, sc.Name, ll)

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
						case req.Polygon.ChainName:
							endHeight = req.Polygon.EndHeight
						case req.CoreDao.ChainName:
							endHeight = req.CoreDao.EndHeight
						case req.L0Bsc.ChainName:
							endHeight = req.L0Bsc.EndHeight
						case req.L0CoreDao.ChainName:
							endHeight = req.L0CoreDao.EndHeight
						case req.L0Polygon.ChainName:
							endHeight = req.L0Polygon.EndHeight
						case req.L0Celo.ChainName:
							endHeight = req.L0Celo.EndHeight
						}
					}
					scs = append(scs, sc.ToNewStatChain2(ydEndHeight, endHeight, s))
					lists[mapKey] = list
				}
			}

		}

		c := chain.ToStatChainNew(scs)
		newChainList = append(newChainList, c)
	}

	//time.Sleep(15 * time.Second)

	w := &sync.WaitGroup{}
	handle.TriggerNftBridgeReceiverStat(newChainList, w, csvRootPath, nowDataStr, nowFileDataStr, lists)
	time.Sleep(5 * time.Second)
	w.Wait()
	utils.Logger.Infof("stat nft bridge sender for receiver task over, time-consuming: %v", time.Now().Sub(now).String())
	return nil
}

func StatDomainAddressTask(config utils.Config, crontab *cron.Cron) error {
	spec := "0 0 8 * * *"
	_, err := crontab.AddFunc(spec, func() {
		DomainAddressTask(config, time.Now())
	})
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	return nil
}

func DomainAddressTask(config utils.Config, now time.Time) error {
	csvRootPath := config.Server.CsvFilePath
	//now := time.Now()

	utils.Logger.Infof("stat domain address task start: %v", now.String())

	yd := now.AddDate(0, 0, -1)
	ydDataStr := yd.Format("2006-01-02")
	ydFileDataStr := yd.Format("2006.01.02")

	endTimestampMap, err := utils.GetPrevDomainEndTimestamp(csvRootPath, ydDataStr)
	if err != nil {
		if !utils.ErrorIsNotExist(err) {
			utils.Logger.Errorln(err)
			return err
		} else {
			endTimestampMap = nil
		}
	}

	nowDataStr := now.Format("2006-01-02")
	nowFileDataStr := now.Format("2006.01.02")

	lists := make(map[string]*sync.Map)
	newDomains := make([]utils.StatDomain, 0)

	for _, domain := range config.StatDomains {
		ydEndHeight := domain.StartTimestamp
		if endTimestampMap != nil {
			ydEndHeight = endTimestampMap[domain.Ty]
		}

		ydCsvFileName := utils.CreateDomainCsvFileName(domain.Ty, ydFileDataStr, ydEndHeight)
		ydCsvFilePath := utils.CreateCsvPath(ydCsvFileName, csvRootPath, ydDataStr, "stat_domain")
		list, ll, err := utils.GetYesterdayCsvData(ydCsvFilePath)

		if err != nil {
			utils.Logger.Errorf("domain ty: %v, get yesterday csv: %v, error: %v", domain.Ty, ydCsvFilePath, err)
		}

		utils.Logger.Infof("domain ty: %v, domain list len: %v", domain.Ty, ll)

		lists[domain.Ty] = list
		newDomains = append(newDomains, utils.StatDomain{
			Url:            domain.Url,
			Ty:             domain.Ty,
			StartTimestamp: ydEndHeight,
			EndTimestamp:   now.Unix(),
		})
	}
	w := &sync.WaitGroup{}
	handle.TriggerDomainStat(w, newDomains, csvRootPath, nowDataStr, nowFileDataStr, lists)
	time.Sleep(5 * time.Second)
	w.Wait()
	utils.Logger.Infof("stat domain address task over, time-consuming: %v", time.Now().Sub(now).String())
	return nil
}

func ReceiverNftTask(config utils.Config, crontab *cron.Cron) error {
	spec := "0 0 8 * * *"
	_, err := crontab.AddFunc(spec, func() {
		StatReceiverNftTask(config, time.Now(), nil)
	})
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	return nil
}

func StatReceiverNftTask(config utils.Config, now time.Time, req *utils.RunParams) error {
	csvRootPath := config.Server.CsvFilePath
	//now := time.Now()

	utils.Logger.Infof("stat nft minter task start: %v", now.String())

	yd := now.AddDate(0, 0, -1)
	ydDataStr := yd.Format("2006-01-02")
	ydFileDataStr := yd.Format("2006.01.02")

	heightMap, err := utils.GetPrevChainReceiverNftEndHeight(csvRootPath, ydDataStr)
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	nowDataStr := now.Format("2006-01-02")
	nowFileDataStr := now.Format("2006.01.02")
	lists := make(map[string]*sync.Map)
	newChainList := make([]utils.Chain, 0)

	for _, chain := range config.StatReceiverNft {

		tokenNAme := ""
		for _, s := range chain.StatContract {
			tokenNAme += s.Name + "|"
		}
		ydEndHeight := heightMap[chain.Name+"-"+tokenNAme]
		mapKey := chain.Name + "-" + tokenNAme
		ydCsvFileName := utils.CreateReceiverNftCsvFileName(chain.Name, tokenNAme, ydFileDataStr, ydEndHeight)
		ydCsvFilePath := utils.CreateCsvPath(ydCsvFileName, csvRootPath, ydDataStr, "stat_nftbridge_receiver_nft")
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
			case req.Polygon.ChainName:
				endHeight = req.Polygon.EndHeight
			case req.Celo.ChainName:
				endHeight = req.Celo.EndHeight
			case req.L0Mantle.ChainName:
				endHeight = req.L0Mantle.EndHeight
			}

		}

		c := chain.ToNewChain(ydEndHeight, endHeight)
		newChainList = append(newChainList, c)
		lists[mapKey] = list
	}

	w := &sync.WaitGroup{}
	handle.TriggerNftBridgeReceiverNftStat(newChainList, w, csvRootPath, nowDataStr, nowFileDataStr, lists)
	time.Sleep(5 * time.Second)
	w.Wait()
	utils.Logger.Infof("stat msg sender task over, time-consuming: %v", time.Now().Sub(now).String())

	return nil
}

func StatClaimNftSender(config utils.Config, crontab *cron.Cron) error {
	spec := "0 0 8 * * *"
	_, err := crontab.AddFunc(spec, func() {
		StatClaimNftSenderTask(config, time.Now(), nil)
	})
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	return nil
}

func StatClaimNftSenderTask(c utils.Config, now time.Time, req *utils.RunParams) error {
	csvRootPath := c.Server.CsvFilePath
	//now := time.Now()

	utils.Logger.Infof("stat nft minter task start: %v", now.String())

	yd := now.AddDate(0, 0, -1)
	ydDataStr := yd.Format("2006-01-02")
	ydFileDataStr := yd.Format("2006.01.02")

	heightMap, err := utils.GetPrevChainStatClaimNftSenderEndHeight(csvRootPath, ydDataStr)
	if err != nil {
		utils.Logger.Errorln(err)
		return err
	}

	nowDataStr := now.Format("2006-01-02")
	nowFileDataStr := now.Format("2006.01.02")
	lists := make(map[string]*sync.Map)
	newScnss := make([]utils.StatClaimNftSender, 0)

	for _, scns := range c.StatClaimNftSender {
		srcTokenName := ""
		for _, token := range scns.Src.Tokens {
			srcTokenName += token.Name + "|"
		}
		if srcTokenName == "" {
			srcTokenName = "null|"
		}

		dstTokenName := ""
		for _, token := range scns.Dst.Tokens {
			dstTokenName += token.Name + "|"
		}
		if dstTokenName == "" {
			dstTokenName = "null|"
		}

		key := scns.Src.Name + "-" + scns.Dst.Name + "-" + srcTokenName + "-" + dstTokenName
		ydEndHeight := heightMap[key]
		ydCsvFileName := utils.CreateClaimNftSenderCsvName(key, ydFileDataStr, ydEndHeight.SrcHeight, ydEndHeight.DstHeight)
		ydCsvFilePath := utils.CreateCsvPath(ydCsvFileName, csvRootPath, ydDataStr, "stat_claim_nft_sender")
		list, ll, err := utils.GetYesterdayCsvData(ydCsvFilePath)
		if err != nil {
			utils.Logger.Errorf("key: %v, get yesterday csv: %v, error: %v", key, ydCsvFilePath, err)
		}

		utils.Logger.Infof("key: %v, msg list len: %v", key, ll)

		srcEndHeight, err := utils.GetChainCurHeight(scns.Src)
		if err != nil {
			utils.Logger.Errorf("src chain: %v, get cur height error", scns.Src.Name)
		}

		dstEndHeight, err := utils.GetChainCurHeight(scns.Dst)
		if err != nil {
			utils.Logger.Errorf("dst chain: %v, get cur height error", scns.Dst.Name)
		}

		if req != nil {
			switch scns.Src.Name {
			case req.ComboTestNet.ChainName:
				srcEndHeight = req.ComboTestNet.EndHeight
			case req.OpbnbTestNet.ChainName:
				srcEndHeight = req.OpbnbTestNet.EndHeight
			}

			switch scns.Dst.Name {
			case req.ComboTestNet.ChainName:
				dstEndHeight = req.ComboTestNet.EndHeight
			case req.OpbnbTestNet.ChainName:
				dstEndHeight = req.OpbnbTestNet.EndHeight
			}

		}

		srcC := scns.Src.ToNewChain(ydEndHeight.SrcHeight, srcEndHeight)
		dstC := scns.Dst.ToNewChain(ydEndHeight.DstHeight, dstEndHeight)
		newScnss = append(newScnss, scns.ToNew(srcC, dstC))
		lists[key] = list
	}

	w := &sync.WaitGroup{}
	handle.TriggerClaimNftSenderStat(newScnss, w, csvRootPath, nowDataStr, nowFileDataStr, lists)
	time.Sleep(5 * time.Second)
	w.Wait()
	utils.Logger.Infof("stat claim nft sender task over, time-consuming: %v", time.Now().Sub(now).String())

	return nil
}
