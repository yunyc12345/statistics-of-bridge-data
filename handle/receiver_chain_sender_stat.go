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

func TriggerNftBridgeReceiverStat(chains []utils.Chain, w *sync.WaitGroup, filePath, dataStr, fileDataStr string, hisList map[string]*sync.Map) {
	fp := filePath + "/" + dataStr + "/stat_nftbridge_receiver_chain_sender"
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

			for i, statChain := range c.StatChains {
				wg.Add(1)

				go func(sc utils.StatChain, index int) {
					defer wg.Done()

					if len(sc.StatContract) == 0 {
						list := &sync.Map{}
						if hisList != nil && len(hisList) != 0 {
							list = hisList[c.Name+"-"+sc.Name+"-"+"all"]
						}

						StatNftBridgeForReceiverHandler(c, list, sc)

						// {chain_name}-{receiver_chain_name}-{contract_name}-{type}-{end_height}-2023/05/11
						name := c.Name + "-" + sc.Name + "-" + "all" + "-" + "stat_sender" + "-" + strconv.FormatUint(sc.EndHeight, 10) + "-" + fileDataStr
						ll, err := utils.ToCsv(list, fp, name)
						if err != nil {
							utils.Logger.Errorf("stat nft bridge receiver chain: %v, sender to csv err: %v", sc.Name, err)
						}
						utils.Logger.Infof("chain: %v, receiver chain: %v, new nft bridge list len: %v", c.Name, sc.Name, ll)
					} else {

						go func(sc2 utils.StatChain) {
							for _, s := range sc2.StatContract {
								list := &sync.Map{}
								if hisList != nil && len(hisList) != 0 {
									list = hisList[c.Name+"-"+sc.Name+"-"+s.Name]
								}

								StatNftBridgeForReceiverAndContractHandler(c, list, sc, s)

								// {chain_name}-{receiver_chain_name}-{contract_name}-{type}-{end_height}-2023/05/11
								name := c.Name + "-" + sc.Name + "-" + s.Name + "-" + "stat_sender" + "-" + strconv.FormatUint(sc.EndHeight, 10) + "-" + fileDataStr
								ll, err := utils.ToCsv(list, fp, name)
								if err != nil {
									utils.Logger.Errorf("stat nft bridge receiver chain: %v, sender to csv err: %v", sc.Name, err)
								}
								utils.Logger.Infof("chain: %v, receiver chain: %v, new nft bridge list len: %v", c.Name, sc.Name, ll)
							}
						}(sc)

					}

				}(statChain, i)
			}
		}(chain)

	}

	time.Sleep(time.Second * 10)
	wg.Wait()

}

func StatNftBridgeForReceiverHandler(chain utils.Chain, list *sync.Map, sc utils.StatChain) {
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
	curHeight, nextHeight := sc.StartHeight, sc.StartHeight

	internal := uint64(3000)
	if chain.Name == "polygon mainnet" {
		internal = uint64(2000)
	}

	contractAddr := chain.NftBridge
	//if cd.Info.Name == "eth mainnet" {
	//	contractAddr = chain.NftBridge
	//}
	//
	//if cd.Info.Name == "bsc mainnet" {
	//	contractAddr = chain.NftBridge
	//}

	utils.Logger.Infof("chain: %v, contract: %v", cd.Info.Name, contractAddr)
	utils.Logger.Infof("chain: %v, topic: %v", cd.Info.Name, event.ID.Hex())

	for curHeight < sc.EndHeight {
		if curHeight+internal >= sc.EndHeight {
			nextHeight = sc.EndHeight
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
			appChainId := uint64(eventData[2].(uint16))
			//srcContractAddr := strings.ToLower(eventData[0].(common.Address).Hex())

			for _, u := range sc.ZkAppChainId {
				if u == appChainId {
					list.Store(sender, utils.Member)
				}
			}
			for _, u := range sc.L0AppChainID {
				if u == appChainId {
					list.Store(sender, utils.Member)
				}
			}

		}

		time.Sleep(1 * time.Second)

		curHeight = nextHeight
	}

}

func StatNftBridgeForReceiverAndContractHandler(chain utils.Chain, list *sync.Map, sc utils.StatChain, target utils.StatContract) {
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
	curHeight, nextHeight := sc.StartHeight, sc.StartHeight

	internal := uint64(3000)
	if chain.Name == "polygon mainnet" || chain.Name == "l0 polygon mainnet" || chain.Name == "celo" || chain.Name == "l0 celo" {
		internal = uint64(2000)
	}

	//if cd.Info.Name == "eth mainnet" {
	//	contractAddr = chain.NftBridge
	//}
	//
	//if cd.Info.Name == "bsc mainnet" {
	//	contractAddr = chain.NftBridge
	//}

	utils.Logger.Infof("chain: %v, topic: %v", cd.Info.Name, event.ID.Hex())

	for curHeight < sc.EndHeight {
		if curHeight+internal >= sc.EndHeight {
			nextHeight = sc.EndHeight
		} else {
			nextHeight += internal
		}

		utils.Logger.Infof("chain: %v, height: %v ~ %v", cd.Info.Name, curHeight, nextHeight)

		query := ethereum.FilterQuery{
			BlockHash: nil,
			FromBlock: big.NewInt(int64(curHeight)), // The error will occur if logs cannot be pulled from the latest block, with a difference of 32.
			ToBlock:   big.NewInt(int64(nextHeight)),
			Addresses: []common.Address{common.HexToAddress(chain.L0NftBridge)},
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
			appChainId := uint64(eventData[2].(uint16))
			srcContractAddr := strings.ToLower(eventData[0].(common.Address).Hex())
			for _, u := range sc.L0AppChainID {
				if u == appChainId && srcContractAddr == strings.ToLower(target.Addr) {
					list.Store(sender, utils.Member)
				}
			}

		}

		query2 := ethereum.FilterQuery{
			BlockHash: nil,
			FromBlock: big.NewInt(int64(curHeight)), // The error will occur if logs cannot be pulled from the latest block, with a difference of 32.
			ToBlock:   big.NewInt(int64(nextHeight)),
			Addresses: []common.Address{common.HexToAddress(chain.NftBridge)},
			Topics:    [][]common.Hash{{event.ID}},
		}

		try2 := 0
		logs2, err := cd.Cli.FilterLogs(ctx, query2)
		for err != nil && try2 <= 10 {
			try++
			utils.Logger.Errorf("chain: %v, try: %v, get logs error :%v", cd.Info.Name, try, err)
			time.Sleep(5 * time.Second)
			logs, err = cd.Cli.FilterLogs(ctx, query)
		}

		if err != nil || try2 > 10 {
			utils.Logger.Errorf("chain: %v, try: %v, get logs error :%v", cd.Info.Name, try, err)
		}

		utils.Logger.Infof("chain: %v, logs len: %v", cd.Info.Name, len(logs))

		for _, l := range logs2 {
			eventData, err := abi.Unpack(event.Name, l.Data)
			if err != nil {
				utils.Logger.Errorf("chain: %v, height: %v ~ %v, parse log err: %v", cd.Info.Name, curHeight, nextHeight, err)
				return
			}
			sender := strings.ToLower(eventData[3].(common.Address).Hex())
			appChainId := uint64(eventData[2].(uint16))
			srcContractAddr := strings.ToLower(eventData[0].(common.Address).Hex())
			for _, u := range sc.ZkAppChainId {
				if u == appChainId && srcContractAddr == strings.ToLower(target.Addr) {
					list.Store(sender, utils.Member)
				}
			}

		}

		time.Sleep(1 * time.Second)

		curHeight = nextHeight
	}

}

func StatNftBridgeReceiversHandler(chain utils.Chain, list *sync.Map) {
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
	curHeight, nextHeight := chain.StartHeight, chain.StartHeight

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
			sender := strings.ToLower(eventData[3].(common.Address).Hex())

			list.Store(sender, utils.Member)

		}

		time.Sleep(1 * time.Second)

		curHeight = nextHeight
	}
}
