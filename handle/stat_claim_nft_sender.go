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

func TriggerClaimNftSenderStat(scnss []utils.StatClaimNftSender, w *sync.WaitGroup, filePath, dataStr, fileDataStr string, hisList map[string]*sync.Map) {
	fp := filePath + "/" + dataStr + "/stat_claim_nft_sender"
	err := utils.CreateCsvDir(fp)
	if err != nil {
		panic(err)
	}

	w.Add(1)
	defer w.Done()
	utils.Logger.Infof("stat claim nft sender handle start")

	wg := &sync.WaitGroup{}
	for _, scns := range scnss {
		wg.Add(1)
		go func(c utils.StatClaimNftSender) {
			defer wg.Done()
			list := &sync.Map{}

			srcTokenName := ""
			for _, token := range c.Src.Tokens {
				srcTokenName += token.Name + "|"
			}
			if srcTokenName == "" {
				srcTokenName = "null|"
			}

			dstTokenName := ""
			for _, token := range c.Dst.Tokens {
				dstTokenName += token.Name + "|"
			}
			if dstTokenName == "" {
				dstTokenName = "null|"
			}

			key := c.Src.Name + "-" + c.Dst.Name + "-" + srcTokenName + "-" + dstTokenName
			if hisList != nil && len(hisList) != 0 {
				list = hisList[key]
			}

			ClaimNftNftHandler(c, list)

			name := key + "-" + "stat_claim_nft_sender" + "-" + strconv.FormatUint(c.Src.EndHeight, 10) + "-" + strconv.FormatUint(c.Dst.EndHeight, 10) + "-" + fileDataStr
			ll, err := utils.ToCsv(list, fp, name)
			if err != nil {
				utils.Logger.Errorf("stat nft bridge receiverNft key: %v, sender to csv err: %v", key, err)
			}
			utils.Logger.Infof("src chain: %v, dst chain: %v, new nft bridge list len: %v", c.Src.Name, c.Dst.Name, ll)

		}(scns)
	}
}

type u struct {
	Src     bool
	SrcHash string
	Dst     bool
	DstHash string
	Sender  string
}

func ClaimNftNftHandler(scns utils.StatClaimNftSender, list *sync.Map) {
	ctx := context.Background()

	data := make(map[string]*u)
	//senderData := make(map[string]*u)
	//receicerData := make(map[string]*u)

	srcCd, err := utils.InitGlobalCliMapAndZKMap(&scns.Src)
	if err != nil {
		utils.Logger.Errorln(err)
		return
	}

	dstCd, err := utils.InitGlobalCliMapAndZKMap(&scns.Dst)
	if err != nil {
		utils.Logger.Errorln(err)
		return
	}

	srcTargetTokenM := make(map[string]bool)
	for _, token := range scns.Src.Tokens {
		srcTargetTokenM[strings.ToLower(token.Address)] = true
	}

	srcChainName := srcCd.Info.Name
	dstChainName := dstCd.Info.Name

	abi, err := contracts.NftBridgeMetaData.GetAbi()
	if err != nil {
		utils.Logger.Errorf("src chain: %v, dst chain: %v, get abi error :%v", srcChainName, dstChainName, err)
		return
	}

	utils.Logger.Infof("dstBridgeContract: %v", scns.Dst.NftBridge)
	utils.Logger.Infof("srcBridgeContract: %v", scns.Src.NftBridge)

	// 0xe11d2ca26838f15acb41450029a785bb3d6f909b7f622ebf9c45524ded76f411
	srcEvent := abi.Events["TransferNFT"]
	utils.Logger.Infof("event: %v", srcEvent.ID)

	srcCurHeight, srcNextHeight := scns.Src.StartHeight, scns.Src.StartHeight
	utils.Logger.Infof("srcEvent: %v", srcEvent.ID.Hex())
	internal := uint64(5000)
	srcBridgeContract := scns.Src.NftBridge

	isOver := false
	for srcCurHeight < scns.Src.EndHeight {
		if srcCurHeight+internal >= scns.Src.EndHeight {
			srcNextHeight = scns.Src.EndHeight
			isOver = true
		} else {
			srcNextHeight += internal
		}

		utils.Logger.Infof("src chain: %v, dst chain: %v, height: %v ~ %v", srcChainName, dstChainName, srcCurHeight, srcNextHeight)

		query := ethereum.FilterQuery{
			BlockHash: nil,
			FromBlock: big.NewInt(int64(srcCurHeight)), // The error will occur if logs cannot be pulled from the latest block, with a difference of 32.
			ToBlock:   big.NewInt(int64(srcNextHeight)),
			Addresses: []common.Address{common.HexToAddress(srcBridgeContract)},
			Topics:    [][]common.Hash{{srcEvent.ID}},
		}

		try := 0
		logs, err := srcCd.Cli.FilterLogs(ctx, query)
		for err != nil && try <= 10 {
			try++
			utils.Logger.Errorf("src chain: %v, dst chain: %v, try: %v, get logs error :%v", srcChainName, dstChainName, try, err)
			time.Sleep(5 * time.Second)
			logs, err = srcCd.Cli.FilterLogs(ctx, query)
		}

		if err != nil || try > 10 {
			utils.Logger.Errorf("src chain: %v, dst chain: %v, try: %v, get logs error :%v", srcChainName, dstChainName, try, err)
		}

		utils.Logger.Infof("src chain: %v, dst chain: %v, logs len: %v", srcChainName, dstChainName, len(logs))

		for _, l := range logs {
			seq := l.Topics[1].Big()
			eventData, err := abi.Unpack(srcEvent.Name, l.Data)
			if err != nil {
				utils.Logger.Errorf("rc chain: %v, dst chain: %v, err: %v", srcChainName, dstChainName, err)
				continue
			}
			srcContractAddr := strings.ToLower(eventData[0].(common.Address).Hex())
			if _, ok := srcTargetTokenM[srcContractAddr]; !ok {
				continue
			}
			sender := strings.ToLower(eventData[3].(common.Address).Hex())
			srcAppId := srcCd.Info.AppChainID
			dstAppId := eventData[2].(uint16)

			srcAppIdStr := strconv.FormatUint(uint64(srcAppId), 10)
			dstAppIdStr := strconv.FormatUint(uint64(dstAppId), 10)

			k := srcAppIdStr + "-" + dstAppIdStr + "-" + seq.String()

			utils.Logger.Infof("sender k: %v, hash: %v, height: %v", k, l.TxHash.Hex(), l.BlockNumber)
			data[k] = &u{
				Src:     true,
				SrcHash: l.TxHash.Hex(),
				Dst:     false,
				Sender:  sender,
			}
			list.Store(sender, utils.Member)
		}

		time.Sleep(1 * time.Second)

		srcCurHeight = srcNextHeight
		if isOver {
			break
		} else {
			continue
		}
	}

	//utils.Logger.Infof("src chain: %v, dst chain: %v, src handle over", srcChainName, dstChainName)
	//utils.Logger.Infof("data len: %v", len(data))
	//// 0x32aae95950c2e1f2c1a419165ba01c63c49604db10ee1b95d9960c0f5b9b9fa8
	//dstEvent := abi.Events["ReceiveNFT"]
	//dstCurHeight, dstNextHeight := scns.Dst.StartHeight, scns.Dst.StartHeight
	//utils.Logger.Infof("dstCurHeight: %v, dstNextHeight: %v", dstCurHeight, dstNextHeight)
	//utils.Logger.Infof("dstEvent: %v", dstEvent.ID.Hex())
	//
	//dstBridgeContract := scns.Dst.NftBridge
	//matchCount := 0
	//
	//isOver = false
	//for dstCurHeight < scns.Dst.EndHeight {
	//	if dstCurHeight+internal >= scns.Dst.EndHeight {
	//		dstNextHeight = scns.Dst.EndHeight
	//		isOver = true
	//	} else {
	//		dstNextHeight += internal
	//	}
	//
	//	utils.Logger.Infof("src chain: %v, dst chain: %v, height: %v ~ %v", srcChainName, dstChainName, dstCurHeight, dstNextHeight)
	//
	//	query := ethereum.FilterQuery{
	//		BlockHash: nil,
	//		FromBlock: big.NewInt(int64(dstCurHeight)), // The error will occur if logs cannot be pulled from the latest block, with a difference of 32.
	//		ToBlock:   big.NewInt(int64(dstNextHeight)),
	//		Addresses: []common.Address{common.HexToAddress(dstBridgeContract)},
	//		Topics:    [][]common.Hash{{dstEvent.ID}},
	//	}
	//
	//	try := 0
	//	logs, err := dstCd.Cli.FilterLogs(ctx, query)
	//	for err != nil && try <= 10 {
	//		try++
	//		utils.Logger.Errorf("src chain: %v, dst chain: %v, try: %v, get logs error :%v", srcChainName, dstChainName, try, err)
	//		time.Sleep(5 * time.Second)
	//		logs, err = dstCd.Cli.FilterLogs(ctx, query)
	//	}
	//
	//	if err != nil || try > 10 {
	//		utils.Logger.Errorf("src chain: %v, dst chain: %v, try: %v, get logs error :%v", srcChainName, dstChainName, try, err)
	//	}
	//
	//	utils.Logger.Infof("src chain: %v, dst chain: %v, logs len: %v", srcChainName, dstChainName, len(logs))
	//
	//	for _, l := range logs {
	//		seq := l.Topics[1].Big()
	//		eventData, err := abi.Unpack(dstEvent.Name, l.Data)
	//		if err != nil {
	//			utils.Logger.Errorf("rc chain: %v, dst chain: %v, err: %v", srcChainName, dstChainName, err)
	//			continue
	//		}
	//		srcAppId := eventData[4].(uint16)
	//		dstAppId := dstCd.Info.AppChainID
	//
	//		srcAppIdStr := strconv.FormatUint(uint64(srcAppId), 10)
	//		dstAppIdStr := strconv.FormatUint(dstAppId, 10)
	//		k := srcAppIdStr + "-" + dstAppIdStr + "-" + seq.String()
	//		utils.Logger.Infof("receiver k: %v, hash: %v, height: %v", k, l.TxHash.Hex(), l.BlockNumber)
	//
	//		//lock.Lock()
	//		if s, ok := data[k]; ok {
	//			s.Dst = true
	//			s.DstHash = l.TxHash.Hex()
	//		}
	//	}
	//
	//	time.Sleep(1 * time.Second)
	//
	//	dstCurHeight = dstNextHeight
	//	if isOver {
	//		break
	//	} else {
	//		continue
	//	}
	//}
	//
	//for _, u := range data {
	//	if u.Src && u.Dst {
	//		list.Store(u.Sender, utils.Member)
	//	}
	//}

	//utils.Logger.Infof("data len: %v, match count: %v", len(data), matchCount)

}
