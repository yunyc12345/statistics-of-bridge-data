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

func TriggerNftStat(chains []utils.Chain, w *sync.WaitGroup, filePath, dataStr, fileDataStr string, hisList map[string]*sync.Map) {
	fp := filePath + "/" + dataStr + "/stat_nft_minter"
	err := utils.CreateCsvDir(fp)
	if err != nil {
		panic(err)
	}

	w.Add(1)
	defer w.Done()
	utils.Logger.Infof("stat nft minter handle start")
	//list := &sync.Map{}
	wg := &sync.WaitGroup{}

	for _, chain := range chains {
		c := chain
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
					//timt_str, u := handleDemand2(c, t, list)
					StatNftMinterHandler(c, t, list)

					// {chain_name}-{token_name}-{type}-2023/05/11
					name := c.Name + "-" + t.Name + "-" + "stat_minter" + "-" + strconv.FormatUint(t.EndHeight, 10) + "-" + fileDataStr
					ll, err := utils.ToCsv(list, fp, name)
					if err != nil {
						utils.Logger.Errorf("stat nft minter to csv err: %v", err)
					}
					utils.Logger.Infof("chain: %v, token: %v, new nft minter list len: %v", c.Name, t.Name, ll)
				}(token, i)

			}
		}(c)

	}

	time.Sleep(time.Second * 10)
	wg.Wait()

	//err := utils.ToCsv(list, "demand2")
	//if err != nil {
	//	utils.Logger.Errorf("demand2 to csv err: %v", err)
	//}
}

func StatNftMinterHandler(chain utils.Chain, t utils.Token, list *sync.Map) {
	ctx := context.Background()

	cd, err := utils.InitGlobalCliMapAndZKMap(&chain)
	if err != nil {
		utils.Logger.Errorln(err)
		return
	}

	abi, err := contracts.NftMetaData.GetAbi()
	if err != nil {
		utils.Logger.Errorf("chain: %v, get abi error :%v", cd.Info.Name, err)
		return
	}

	event := abi.Events["Transfer"]
	curHeight, nextHeight := t.StartHeight, t.StartHeight

	internal := uint64(3000)
	if chain.Name == "polygon mainnet" {
		internal = uint64(2000)
	}

	//contractAddr := ""
	//if cd.Info.Name == "eth mainnet" {
	//	contractAddr = "0x1e40CD8569F3c91F5101d54AE01a75574a9ccE60"
	//}
	//
	//if cd.Info.Name == "bsc mainnet" {
	//	contractAddr = "0xE09828f0DA805523878Be66EA2a70240d312001e"
	//}

	//utils.Logger.Infof("chain: %v, contract: %v", cd.Info.Name, contractAddr)
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
			Addresses: []common.Address{common.HexToAddress(t.Address)},
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
			from := strings.ToLower(common.BytesToAddress(l.Topics[1].Bytes()).Hex())
			//utils.Logger.Infof("from: %v", from)
			minter := strings.ToLower(common.BytesToAddress(l.Topics[2].Bytes()).Hex())
			//utils.Logger.Infof("minter: %v", minter)
			if from == "0x0000000000000000000000000000000000000000" {
				list.Store(minter, utils.Member)
			}

		}

		time.Sleep(1 * time.Second)

		curHeight = nextHeight
	}

}

func handleDemand2(chain utils.Chain, t utils.Token, list *sync.Map) (string, uint64) {

	cd, err := utils.InitGlobalCliMapAndZKMap(&chain)
	if err != nil {
		utils.Logger.Errorln(err)
		return "", 0
	}

	supply := &big.Int{}
	var nftC *contracts.Nft
	switch t.Name {
	case "loyalty":
		supply, err = cd.Loyalty.TotalSupply(nil)
		if err != nil {
			utils.Logger.Errorf("chain: %v, get nft supply error :%v", cd.Info.Name, err)
			return "", 0
		}
		nftC = cd.Loyalty
		break
	default:
		panic("")
	}

	timeStr := time.Now().Format("2006.01.02 15:04")
	totalNftNum := supply.Uint64()
	utils.Logger.Errorf("chain: %v, loyalty nft supply: %v", cd.Info.Name, totalNftNum)
	internal := uint64(5000)

	count := totalNftNum / internal

	if count >= 1 {
		total := uint64(0)
		wg := &sync.WaitGroup{}
		for i := uint64(0); i < count; i++ {
			wg.Add(1)
			go getAddress(nftC, total, total+internal, list, wg, cd.Info.Name)
			total += internal
			//time.Sleep(2 * time.Second)
		}

		if total < totalNftNum {
			wg.Add(1)
			go getAddress(nftC, total, totalNftNum-total, list, wg, cd.Info.Name)
			//time.Sleep(2 * time.Second)
		}

		time.Sleep(10 * time.Second)
		wg.Wait()
	} else {
		for i := uint64(0); i <= totalNftNum; i++ {
			utils.Logger.Infof("chain: %v, tokenId: %v", cd.Info.Name, i)
			address, err := nftC.OwnerOf(nil, big.NewInt(int64(i)))
			if err != nil {
				utils.Logger.Errorf("chain: %v, get nft tokenId: %v, error :%v", cd.Info.Name, i, err)
				continue
			}

			addr := strings.ToLower(address.Hex())

			list.Store(addr, utils.Member)
			time.Sleep(1 * time.Second)
		}
	}

	return timeStr, totalNftNum
}

func getAddress(nftC *contracts.Nft, s, e uint64, list *sync.Map, wg *sync.WaitGroup, chainName string) {
	//wg.Add(1)
	defer wg.Done()
	for i := s; i < e; i++ {
		utils.Logger.Infof("chain: %v, tokenId: %v", chainName, i)
		address, err := nftC.OwnerOf(nil, big.NewInt(int64(i)))
		if err != nil {
			utils.Logger.Errorf("chain: %v, get nft tokenId: %v, error :%v", chainName, i, err)
		}

		addr := strings.ToLower(address.Hex())

		list.Store(addr, utils.Member)
		time.Sleep(1 * time.Second)
	}
}
