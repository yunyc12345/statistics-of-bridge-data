package utils

import (
	"encoding/csv"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"os"
	"testing"
	"yunyc12345/statistics-of-bridge-data/contracts"
)

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

func Test3(t *testing.T) {
	//list := &sync.Map{}
	fs, err := os.Open("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv2/stat_msg_sender/2023-05-11/eth mainnet-msg-17233413-2023.05.11.csv")
	if err != nil {
		Logger.Errorln(err)
		panic("")
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			Logger.Errorln(err)
			panic("")
		}
		if err == io.EOF {
			break
		}
		//j, _ := json.Marshal(row[0])
		println(row[0])
	}
}
