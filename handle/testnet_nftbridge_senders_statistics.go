package handle

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"sync"
	"time"
	"yunyc12345/statistics-of-bridge-data/contracts"
	"yunyc12345/statistics-of-bridge-data/utils"
)

func TestnetNftbridgeSendersStatistics(chains []utils.Chain, w *sync.WaitGroup) {
	w.Add(1)
	defer w.Done()
	utils.Logger.Infof("demand1 handle start")
	list := &sync.Map{}
	wg := sync.WaitGroup{}
	for _, chain := range chains {
		c := chain
		go handleDemand1(c, list, &wg)
	}
	time.Sleep(time.Second * 10)
	wg.Wait()

	err := utils.ToCsv(list, "demand1")
	if err != nil {
		utils.Logger.Errorf("demand1 to csv err: %v", err)
	}
}

func handleDemand1(chain utils.Chain, list *sync.Map, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
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
	internal := 1000

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
			Addresses: []common.Address{common.HexToAddress(cd.Info.NftBridge)},
			Topics:    [][]common.Hash{{event.ID}},
		}

		try := 0
		logs, err := cd.Cli.FilterLogs(ctx, query)
		for err != nil && try <= 5 {
			try++
			utils.Logger.Errorf("chain: %v, try: %v, get logs error :%v", cd.Info.Name, try, err)
			logs, err = cd.Cli.FilterLogs(ctx, query)
		}

		if err != nil || try > 5 {
			utils.Logger.Errorf("chain: %v, try: %v, get logs error :%v", cd.Info.Name, try, err)
		}

		for _, l := range logs {
			eventData, err := abi.Unpack(event.Name, l.Data)
			if err != nil {
				utils.Logger.Errorf("chain: %v, height: %v ~ %v, parse log err: %v", cd.Info.Name, curHeight, nextHeight, err)
				return
			}
			sender := strings.ToLower(eventData[0].(common.Address).Hex())
			list.Store(sender, utils.Member)
		}

		time.Sleep(3 * time.Second)

		curHeight = nextHeight
	}

}
