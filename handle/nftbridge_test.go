package handle

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"testing"
	"time"
	"yunyc12345/statistics-of-bridge-data/contracts"
	"yunyc12345/statistics-of-bridge-data/utils"
)

func TestCobeeStat(t *testing.T) {

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

		fmt.Printf("chain: bsc, height: %v ~ %v", curHeight, nextHeight)

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
			fmt.Printf("chain: bsc, try: %v, get logs error :%v", try, err)
			time.Sleep(5 * time.Second)
			logs, err = c.FilterLogs(ctx, query)
		}

		if err != nil || try > 10 {
			fmt.Printf("chain: bsc, try: %v, get logs error :%v", try, err)
		}

		fmt.Printf("chain: bsc, logs len: %v", len(logs))

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
