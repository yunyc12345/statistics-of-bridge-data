package handle

//func MainnetNfterc721SendersStatistics(chains []utils.Chain, w *sync.WaitGroup) {
//	w.Add(1)
//	defer w.Done()
//	utils.Logger.Infof("demand2 handle start")
//	list := &sync.Map{}
//	wg := sync.WaitGroup{}
//	for _, chain := range chains {
//		c := chain
//		go handleDemand2(c, list, &wg)
//	}
//	time.Sleep(time.Second * 10)
//	wg.Wait()
//
//	err := utils.ToCsv(list, "demand2")
//	if err != nil {
//		utils.Logger.Errorf("demand2 to csv err: %v", err)
//	}
//}
//
//func handleDemand2(chain utils.Chain, list *sync.Map, wg *sync.WaitGroup) {
//	wg.Add(1)
//	defer wg.Done()
//
//	cd, err := utils.InitGlobalCliMapAndZKMap(&chain)
//	if err != nil {
//		utils.Logger.Errorln(err)
//		return
//	}
//
//	supply, err := cd.Nft.TotalSupply(nil)
//	if err != nil {
//		utils.Logger.Errorf("chain: %v, get nft supply error :%v", cd.Info.Name, err)
//		return
//	}
//	totalNftNum := supply.Int64()
//	internal := int64(10000)
//
//	count := totalNftNum / internal
//
//	if count >= 1 {
//		total := int64(0)
//		for i := int64(0); i < count; i++ {
//			go getAddress(cd, total, total+internal, list, wg)
//			total += internal
//			//time.Sleep(2 * time.Second)
//		}
//
//		if total < totalNftNum {
//			go getAddress(cd, total, totalNftNum-total, list, wg)
//			//time.Sleep(2 * time.Second)
//		}
//
//		time.Sleep(10 * time.Second)
//
//	} else {
//		for i := int64(0); i <= totalNftNum; i++ {
//			utils.Logger.Infof("chain: %v, tokenId: %v", cd.Info.Name, i)
//			address, err := cd.Nft.OwnerOf(nil, big.NewInt(i))
//			if err != nil {
//				utils.Logger.Errorf("chain: %v, get nft tokenId: %v, error :%v", cd.Info.Name, i, err)
//				continue
//			}
//
//			addr := strings.ToLower(address.Hex())
//
//			list.Store(addr, utils.Member)
//			time.Sleep(1 * time.Second)
//		}
//	}
//}
//
//func getAddress(cd *utils.ChainData, s, e int64, list *sync.Map, wg *sync.WaitGroup) {
//	wg.Add(1)
//	defer wg.Done()
//	for i := s; i < e; i++ {
//		utils.Logger.Infof("chain: %v, tokenId: %v", cd.Info.Name, i)
//		address, err := cd.Nft.OwnerOf(nil, big.NewInt(i))
//		if err != nil {
//			utils.Logger.Errorf("chain: %v, get nft tokenId: %v, error :%v", cd.Info.Name, i, err)
//		}
//
//		addr := strings.ToLower(address.Hex())
//
//		list.Store(addr, utils.Member)
//		time.Sleep(1 * time.Second)
//	}
//}
