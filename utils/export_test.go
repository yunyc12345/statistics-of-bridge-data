package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
	"testing"
	"yunyc12345/statistics-of-bridge-data/contracts"
)

func TestExport(t *testing.T) {

	InitLogger(&LogConf{
		Level: "info",
		Path:  "log/app.log",
	})

	list := &sync.Map{}

	list.Store("0xd9249E3D614dAa3E2e94aEAEe5F21f94E4AdeC9a", Member)
	list.Store("0x35B0AD80dF8b29AF65CAc60Df30D5604D37FEc29", Member)
	list.Store("0x35B0AD80dF8b29AF65CAc60Df30D5604D37FEc29", Member)
	list.Store("0xE09828f0DA805523878Be66EA2a70240d312001e", Member)

	err := ToCsv(list, "mm")
	if err != nil {
		panic("")
	}
}

func Test2(t *testing.T) {
	c, _ := ethclient.Dial("https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e")
	nft, _ := contracts.NewNft(common.HexToAddress("0xC5ae0d15593316e0cC905840eD2dE83E2DD4EA9E"), c)
	supply, _ := nft.TotalSupply(nil)
	println(supply.Uint64())

	c2, _ := ethclient.Dial("https://eth-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e")
	nft2, _ := contracts.NewNft(common.HexToAddress("0x0154494A8f48C54772642055B6669196d9749dED"), c2)
	supply2, _ := nft2.TotalSupply(nil)
	println(supply2.Uint64())

}
