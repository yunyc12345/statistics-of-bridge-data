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

func TriggerNftBridgeReceiverNftStat(chains []utils.Chain, w *sync.WaitGroup, filePath, dataStr, fileDataStr string, hisList map[string]*sync.Map) {
	fp := filePath + "/" + dataStr + "/stat_nftbridge_receiver_nft"
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

			tokenNAme := ""
			for _, s := range c.StatContract {
				tokenNAme += s.Name + "|"
			}

			list := &sync.Map{}
			if hisList != nil && len(hisList) != 0 {
				list = hisList[c.Name+"-"+tokenNAme]
			}

			StatNftBridgeForReceiverNftHandler(c, list)

			name := c.Name + "-" + tokenNAme + "-" + "stat_receiver" + "-" + strconv.FormatUint(c.EndHeight, 10) + "-" + fileDataStr
			ll, err := utils.ToCsv(list, fp, name)
			if err != nil {
				utils.Logger.Errorf("stat nft bridge receiverNft chain: %v, sender to csv err: %v", c.Name, err)
			}
			utils.Logger.Infof("chain: %v, receiverNft chain: %v, new nft bridge list len: %v", c.Name, c.Name, ll)

		}(chain)

	}

	time.Sleep(time.Second * 10)
	wg.Wait()

}

func StatNftBridgeForReceiverNftHandler(chain utils.Chain, list *sync.Map) {
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

	event := abi.Events["ReceiveNFT"]
	curHeight, nextHeight := chain.StartHeight, chain.StartHeight

	internal := uint64(3000)
	if chain.Name == "polygon mainnet" {
		internal = uint64(2000)
	}

	contractAddr := chain.L0NftBridge

	utils.Logger.Infof("chain: %v, contract: %v", cd.Info.Name, contractAddr)
	utils.Logger.Infof("chain: %v, topic: %v", cd.Info.Name, event.ID.Hex())

	targetMap := make(map[string]bool)
	for _, s := range chain.StatContract {
		targetMap[strings.ToLower(s.Addr)] = true
	}

	for curHeight < chain.EndHeight {
		if curHeight+internal >= chain.EndHeight {
			nextHeight = chain.EndHeight
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
			receiver := eventData[5].(common.Address).Hex()
			//srcContractAddr := strings.ToLower(eventData[0].(common.Address).Hex())
			dstContractAddr := strings.ToLower(eventData[1].(common.Address).Hex())

			if _, ok := targetMap[dstContractAddr]; ok {
				list.Store(receiver, utils.Member)
			}

			//for _, contract := range chain.StatContract {
			//	if strings.ToLower(contract.Addr) == srcContractAddr {
			//		list.Store(receiver, utils.Member)
			//		break
			//	}
			//}

		}

		time.Sleep(1 * time.Second)

		curHeight = nextHeight
	}

}
