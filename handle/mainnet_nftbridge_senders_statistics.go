package handle

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"
	"yunyc12345/statistics-of-bridge-data/contracts"
	"yunyc12345/statistics-of-bridge-data/utils"
)

func TriggerNftBridgeStat(chains []utils.Chain, w *sync.WaitGroup, filePath, dataStr, fileDataStr string, hisList map[string]*sync.Map) {
	fp := filePath + "/" + dataStr + "/stat_nftbridge_sender"
	err := utils.CreateCsvDir(fp)
	if err != nil {
		panic(err)
	}
	j, _ := json.Marshal(chains)
	utils.Logger.Info(string(j))

	w.Add(1)
	defer w.Done()
	utils.Logger.Infof("stat nft bridge sender handle start")

	wg := &sync.WaitGroup{}
	for _, chain := range chains {
		wg.Add(1)
		go func(c utils.Chain) {

			defer wg.Done()

			for i, token := range c.Tokens {

				wg.Add(1)

				go func(t utils.Token, index int) {

					defer wg.Done()
					list := &sync.Map{}

					if hisList != nil && len(hisList) != 0 {
						list = hisList[c.Name+"-"+t.Name]
					}
					StatNftBridgeHandler(c, t, list)

					// {chain_name}-{token_name}-{type}-{end_height}-2023/05/11
					name := c.Name + "-" + t.Name + "-" + "stat_sender" + "-" + strconv.FormatUint(t.EndHeight, 10) + "-" + fileDataStr
					ll, err := utils.ToCsv(list, fp, name)
					if err != nil {
						utils.Logger.Errorf("stat nft bridge sender to csv err: %v", err)
					}
					utils.Logger.Infof("chain: %v, token: %v, new nft bridge list len: %v", c.Name, t.Name, ll)
				}(token, i)

			}
		}(chain)

	}

	time.Sleep(time.Second * 10)
	wg.Wait()

}

func StatNftBridgeHandler(chain utils.Chain, t utils.Token, list *sync.Map) {
	j, _ := json.Marshal(t)
	utils.Logger.Infof("chain: %v, token: %v, json: %v", chain.Name, t.Name, string(j))
	ctx := context.Background()

	cd, err := utils.InitGlobalCliMapAndZKMap(&chain)
	if err != nil {
		utils.Logger.Errorln(err)
		return
	}

	abi, err := contracts.NftBridgeMetaData.GetAbi()
	if err != nil {
		utils.Logger.Errorf("chain: %v, get abi error :%v", cd.Info.Name, err)
		return
	}

	event := abi.Events["TransferNFT"]
	curHeight, nextHeight := t.StartHeight, t.StartHeight

	internal := uint64(10000)

	contractAddr := ""
	if cd.Info.Name == "eth mainnet" {
		contractAddr = "0x1e40CD8569F3c91F5101d54AE01a75574a9ccE60"
	}

	if cd.Info.Name == "bsc mainnet" {
		contractAddr = "0xE09828f0DA805523878Be66EA2a70240d312001e"
	}

	utils.Logger.Infof("chain: %v, contract: %v", cd.Info.Name, contractAddr)
	utils.Logger.Infof("chain: %v, topic: %v", cd.Info.Name, event.ID.Hex())

	for curHeight < t.EndHeight {
		if curHeight+internal >= t.EndHeight {
			nextHeight = t.EndHeight
		} else {
			nextHeight += internal
		}

		utils.Logger.Infof("chain: %v, height: %v ~ %v", cd.Info.Name, curHeight, nextHeight)

		query := ethereum.FilterQuery{
			BlockHash: nil,
			FromBlock: big.NewInt(int64(curHeight)), // The error will occur if logs cannot be pulled from the latest block, with a difference of 32.
			ToBlock:   big.NewInt(int64(nextHeight)),
			Addresses: []common.Address{common.HexToAddress(contractAddr)},
			Topics:    [][]common.Hash{{event.ID}},
		}

		try := 0
		logs, err := cd.Cli.FilterLogs(ctx, query)
		for err != nil && try <= 10 {
			try++
			utils.Logger.Errorf("chain: %v, try: %v, get logs error :%v", cd.Info.Name, try, err)
			time.Sleep(5 * time.Second)
			logs, err = cd.Cli.FilterLogs(ctx, query)
		}

		if err != nil || try > 10 {
			utils.Logger.Errorf("chain: %v, try: %v, get logs error :%v", cd.Info.Name, try, err)
		}

		utils.Logger.Infof("chain: %v, logs len: %v", cd.Info.Name, len(logs))

		for _, l := range logs {
			eventData, err := abi.Unpack(event.Name, l.Data)
			if err != nil {
				utils.Logger.Errorf("chain: %v, height: %v ~ %v, parse log err: %v", cd.Info.Name, curHeight, nextHeight, err)
				return
			}
			sender := strings.ToLower(eventData[3].(common.Address).Hex())
			tokenAddress := strings.ToLower(eventData[0].(common.Address).Hex())
			if tokenAddress == strings.ToLower(t.Address) {
				list.Store(sender, utils.Member)
			}

		}

		time.Sleep(1 * time.Second)

		curHeight = nextHeight
	}

}
