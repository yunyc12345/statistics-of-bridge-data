package handle

import (
	"context"
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

func TriggerMsgStat(chains []utils.Chain, w *sync.WaitGroup, filePath, dataStr, fileDataStr string, hisList []*sync.Map, diffMonth bool) {
	fp := filePath + "/" + dataStr + "/stat_msg_sender"
	err := utils.CreateCsvDir(fp)
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}

	w.Add(1)
	defer w.Done()
	utils.Logger.Infof("stat msg sender handle start")

	wg := &sync.WaitGroup{}
	for i, chain := range chains {

		c := chain
		wg.Add(1)
		go func(c utils.Chain, index int) {
			defer wg.Done()
			list := &sync.Map{}
			if hisList != nil && len(hisList) != 0 {
				list = hisList[index]
			}
			if diffMonth {
				list = &sync.Map{}
			}
			StatMsgSenderHandler(c, list)

			// {chain_name}-{type}-{end_height}-2023/05/11
			name := c.Name + "-" + "msg" + "-" + strconv.FormatUint(c.EndHeight, 10) + "-" + fileDataStr
			ll, err := utils.ToCsv(list, fp, name)
			if err != nil {
				utils.Logger.Errorf("stat msg to csv err: %v", err)
			}
			utils.Logger.Infof("chain: %v, new msg list len: %v", c.Name, ll)
		}(c, i)

	}

	time.Sleep(time.Second * 10)
	wg.Wait()

}

func StatMsgSenderHandler(chain utils.Chain, list *sync.Map) {
	//wg.Add(1)
	//defer wg.Done()
	ctx := context.Background()

	cd, err := utils.InitGlobalCliMapAndZKMap(&chain)
	if err != nil {
		utils.Logger.Errorln(err)
		return
	}

	abi, err := contracts.MailerMetaData.GetAbi()
	if err != nil {
		utils.Logger.Errorf("chain: %v, get abi error :%v", cd.Info.Name, err)
		return
	}

	event := abi.Events["MessageSend"]
	curHeight, nextHeight := chain.StartHeight, chain.StartHeight
	internal := uint64(3000)
	if chain.Name == "polygon mainnet" {
		internal = uint64(2000)
	}

	if chain.Name == "celo" {
		internal = uint64(1000)
	}

	//contractAddr := cd.Info.NftBridge
	//contractGfAddr := cd.Info.MailboxGreenfield
	//if contractAddr == "" {
	//	if cd.Info.Name == "eth mainnet" {
	//		contractAddr = "0x1e40CD8569F3c91F5101d54AE01a75574a9ccE60"
	//	}
	//
	//	if cd.Info.Name == "bsc mainnet" {
	//		contractAddr = "0xE09828f0DA805523878Be66EA2a70240d312001e"
	//	}
	//
	//}

	utils.Logger.Infof("chain: %v, gf: %v", cd.Info.Name, cd.Info.MailerGreenfield)
	utils.Logger.Infof("chain: %v, nomarl: %v", cd.Info.Name, cd.Info.Mailer)
	utils.Logger.Infof("chain: %v, event: %v", cd.Info.Name, event.ID.Hex())

	for curHeight < chain.EndHeight {
		if curHeight+internal >= chain.EndHeight {
			nextHeight = chain.EndHeight
		} else {
			nextHeight += internal
		}

		utils.Logger.Infof("chain: %v, height: %v ~ %v", cd.Info.Name, curHeight, nextHeight)

		addrs := make([]common.Address, 0)
		for _, s := range cd.Info.Mailer {
			addrs = append(addrs, common.HexToAddress(s))
		}

		for _, s := range cd.Info.MailerGreenfield {
			addrs = append(addrs, common.HexToAddress(s))
		}

		query := ethereum.FilterQuery{
			BlockHash: nil,
			FromBlock: big.NewInt(int64(curHeight)), // The error will occur if logs cannot be pulled from the latest block, with a difference of 32.
			ToBlock:   big.NewInt(int64(nextHeight)),
			Addresses: addrs,
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
			sender := strings.ToLower(eventData[0].(common.Address).Hex())
			list.Store(sender, utils.Member)
		}

		time.Sleep(1 * time.Second)

		curHeight = nextHeight
	}

}
