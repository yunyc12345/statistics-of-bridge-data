package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"sync"
	"time"
	"yunyc12345/statistics-of-bridge-data/contracts"
	"yunyc12345/statistics-of-bridge-data/handle"
	"yunyc12345/statistics-of-bridge-data/utils"
)

func eth_alpha_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "eth mainnet",
		Endpoint:          "https://eth-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       0,
		EndHeight:         0,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "alpha",
		Address:        "0x0154494A8f48C54772642055B6669196d9749dED",
		StartHeight:    16950603,
		StartTimestamp: 0,
		EndHeight:      17259152,
		EndTimestamp:   0,
	}
	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "eth mainnet-alpha-stat_minter-17259152-2023.05.15")
	if err != nil {
		utils.Logger.Errorln(err)
	}

	return
}

func eth_zkLightClient_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "eth mainnet",
		Endpoint:          "https://eth-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       0,
		EndHeight:         0,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "zkLightClient",
		Address:        "0x4D8910d4AbA36f856D5C1aD9f43d667E54E0A964",
		StartHeight:    17384358,
		StartTimestamp: 0,
		EndHeight:      17389364,
		EndTimestamp:   0,
	}
	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "eth mainnet-zkLightClient-stat_minter-17259152-2023.06.02")
	if err != nil {
		utils.Logger.Errorln(err)
	}

	return
}

func bsc_alpha_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       0,
		EndHeight:         0,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "alpha",
		Address:        "0xC5ae0d15593316e0cC905840eD2dE83E2DD4EA9E",
		StartHeight:    26956110,
		StartTimestamp: 0,
		EndHeight:      28207389,
		EndTimestamp:   0,
	}
	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-alpha-stat_minter-28207389-2023.05.15")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func bsc_zkLightClient_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       0,
		EndHeight:         0,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "zkLightClient",
		Address:        "0xD2cCC9EE7Ea2ccd154c727A46D475ddA49E99852",
		StartHeight:    28714086,
		StartTimestamp: 0,
		EndHeight:      28734406,
		EndTimestamp:   0,
	}
	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-zkLightClient-stat_minter-28207389-2023.06.02")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func bsc_gffaucet_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       0,
		EndHeight:         0,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "gffaucet",
		Address:        "0x13D23d867e73aF912Adf5d5bd47915261eFa28F2",
		StartHeight:    28858087,
		StartTimestamp: 0,
		EndHeight:      28964277,
		EndTimestamp:   0,
	}
	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-gffaucet-stat_minter-28964277-2023.06.10")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func bsc_faucet_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       0,
		EndHeight:         0,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "faucet",
		Address:        "0x9885C17Dd44c00C37B98F510cdff099EfF437dcE",
		StartHeight:    29028650,
		StartTimestamp: 0,
		EndHeight:      29033185,
		EndTimestamp:   0,
	}
	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-faucet-stat_minter-29033185-2023.06.12")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func bsc_luban_loyalty_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       0,
		EndHeight:         0,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "luban_loyalty",
		Address:        "0x9885C17Dd44c00C37B98F510cdff099EfF437dcE",
		StartHeight:    29028650,
		StartTimestamp: 0,
		EndHeight:      29165411,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-luban_loyalty-stat_minter-29165411-2023.06.17")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func polygon_alpha_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "polygon mainnet",
		Endpoint:          "https://polygon-rpc.com",
		ZkBridge:          "",
		NftBridge:         "0x2E953a70C37E8CB4553DAe1F5760128237c8820D",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "alpha",
		Address:        "0x9d5d479a84f3358e8e27afe056494bd2da239acd",
		StartHeight:    43122427,
		StartTimestamp: 0,
		EndHeight:      43346764,
		EndTimestamp:   0,
	}
	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "polygon mainnet-alpha-stat_minter-43346764-2023.05.31")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func coredao_alpha_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "core dao",
		Endpoint:          "https://rpc.coredao.org",
		ZkBridge:          "",
		NftBridge:         "0xE09828f0DA805523878Be66EA2a70240d312001e",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         4631608,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "alpha",
		Address:        "0x61DFDbcC65DaF1F60fB1DbE703D84940dA28526c",
		StartHeight:    5420650,
		StartTimestamp: 0,
		EndHeight:      5737400,
		EndTimestamp:   0,
	}
	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-07-17/stat_nft_minter/core dao-alpha-stat_minter-5420650-2023.07.17")
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}
	handle.StatNftMinterHandler(c, t, list)
	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "core dao-alpha-stat_minter-5737400-2023.07.28")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func coredao_opbnb_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "core dao",
		Endpoint:          "https://rpc.coredao.org",
		ZkBridge:          "",
		NftBridge:         "0xE09828f0DA805523878Be66EA2a70240d312001e",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         4631608,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "opbnb",
		Address:        "0x0f83DA622E36Ee42cfeB222257E1baF20E16a491",
		StartHeight:    4651650,
		StartTimestamp: 0,
		EndHeight:      4672120,
		EndTimestamp:   0,
	}
	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "core dao-opbnb-stat_minter-4672120-2023.06.21")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func bsc_opbnb_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       0,
		EndHeight:         0,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "opbnb",
		Address:        "0x9c614a8E5a23725214024d2C3633BE30D44806f9",
		StartHeight:    29259397,
		StartTimestamp: 0,
		EndHeight:      29280184,
		EndTimestamp:   0,
	}
	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-opbnb-stat_minter-29280184-2023.06.21")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func polygon_opbnb_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "polygon mainnet",
		Endpoint:          "https://polygon-rpc.com",
		ZkBridge:          "",
		NftBridge:         "0x2E953a70C37E8CB4553DAe1F5760128237c8820D",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "opbnb",
		Address:        "0xfeb105763753e9d26DfD4aae1Ed368aa7cC18260",
		StartHeight:    44120096,
		StartTimestamp: 0,
		EndHeight:      44147776,
		EndTimestamp:   0,
	}
	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "polygon mainnet-opbnb-stat_minter-44147776-2023.06.21")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func eth_opbnb_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "eth mainnet",
		Endpoint:          "https://eth-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       0,
		EndHeight:         0,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "opbnb",
		Address:        "0xcB6cA8eE67E90115dbCe8b0a990443bF7015ceD3",
		StartHeight:    17519216,
		StartTimestamp: 0,
		EndHeight:      17524328,
		EndTimestamp:   0,
	}
	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "eth mainnet-opbnb-stat_minter-17524328-2023.06.21")
	if err != nil {
		utils.Logger.Errorln(err)
	}

	return
}

func polygon_zkLightClient_mint_stat(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "polygon mainnet",
		Endpoint:          "https://polygon-rpc.com",
		ZkBridge:          "",
		NftBridge:         "0x2E953a70C37E8CB4553DAe1F5760128237c8820D",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "zkLightClient",
		Address:        "0x6b0C248679F493481411a0A14cd5FC2DBBe8Ab02",
		StartHeight:    43398055,
		StartTimestamp: 0,
		EndHeight:      43425856,
		EndTimestamp:   0,
	}
	list := &sync.Map{}
	handle.StatNftMinterHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "polygon mainnet-zkLightClient-stat_minter-43346764-2023.06.02")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func bsc_stat_msg(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            []string{"0x5c5979832A60c17bB06676fa906bEdD1A013e18c", "0xfD3F4D96378072db0862a6F76CC258C2B7ea36cc"},
		Mailbox:           "",
		MailerGreenfield:  []string{"0x044332f4b34fD5639482dE41cD1d767F7304fA00", "0xf0295A8caD5f287CC52b6F5994fE4aa6FB6e8D4B"},
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       28705724,
		EndHeight:         29136683,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}
	list := &sync.Map{}
	handle.StatMsgSenderHandler(c, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-msg-29136683-2023.06.16")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func eth_stat_msg(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "eth mainnet",
		Endpoint:          "https://eth-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            []string{"0x1Fbc29bD726257EfE5555D47e71030a32426cEbA", "0xbf782f6Eb61d617f7f72B978329339Aa728D21DB"},
		Mailbox:           "",
		MailerGreenfield:  []string{"0xB88fE84F922Caf34184EfEbD4146088f5bE0200B", "0x7Ea2c4d09AB095EEbdE5890a08869167bC762Ec3"},
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       17382265,
		EndHeight:         17488698,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}
	list := &sync.Map{}
	handle.StatMsgSenderHandler(c, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "eth mainnet-msg-17488698-2023.06.16")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func polygon_stat_msg(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "polygon mainnet",
		Endpoint:          "https://polygon-rpc.com",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            []string{"0x81f0d785F82630Bbe6639B7eB07EAB21E2992293", "0xdB6fb08DD8Ce406DA8Ff53FAe65Bd374e3d68681"},
		Mailbox:           "",
		MailerGreenfield:  []string{"0x3fFb989765868CA9f75c90F94A7A17AaB2a981D7", "0x8163A9B0901F63c27471b4d051b7250ECDDD362d"},
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       43386770,
		EndHeight:         43956704,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}
	list := &sync.Map{}
	handle.StatMsgSenderHandler(c, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "polygon mainnet-msg-43956704-2023.06.16")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func coredao_stat_msg(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "core dao",
		Endpoint:          "https://rpc.coredao.org",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            []string{"0x373d2c6dAbb65F1F85a7bC058664b205da0B0824"},
		Mailbox:           "",
		MailerGreenfield:  []string{"0xa0c6ED9Dd8431c40C74BCeEAdA5DaC996dB51f48"},
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       4931307,
		EndHeight:         4960083,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}
	list := &sync.Map{}
	handle.StatMsgSenderHandler(c, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "core dao-msg-4931307-2023.06.30")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func celo_stat_msg(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "celo",
		Endpoint:          "https://rpc.ankr.com/celo",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            []string{"0x468Bd5792543b47aaeF12C16E668E7Fd3Fe05872"},
		Mailbox:           "",
		MailerGreenfield:  []string{"0xEae5c76B250206140a27CA0B75DEB28D7D7EAd9B"},
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       20082986,
		EndHeight:         20100265,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}
	list := &sync.Map{}
	handle.StatMsgSenderHandler(c, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "celo-msg-20082986-2023.06.30")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func bsc_stat_msg_and_hist(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            []string{"0x5c5979832A60c17bB06676fa906bEdD1A013e18c", "0xfD3F4D96378072db0862a6F76CC258C2B7ea36cc"},
		Mailbox:           "",
		MailerGreenfield:  []string{"0x044332f4b34fD5639482dE41cD1d767F7304fA00", "0xf0295A8caD5f287CC52b6F5994fE4aa6FB6e8D4B"},
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       29538700,
		EndHeight:         29567407,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}
	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-06-30/stat_msg_sender/bsc mainnet-msg-29538700-2023.06.30")
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}
	handle.StatMsgSenderHandler(c, list)
	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-msg-29567407-2023.07.01")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func eth_stat_msg_and_hist(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "eth mainnet",
		Endpoint:          "https://eth-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            []string{"0x1Fbc29bD726257EfE5555D47e71030a32426cEbA", "0xbf782f6Eb61d617f7f72B978329339Aa728D21DB"},
		Mailbox:           "",
		MailerGreenfield:  []string{"0xB88fE84F922Caf34184EfEbD4146088f5bE0200B", "0x7Ea2c4d09AB095EEbdE5890a08869167bC762Ec3"},
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       17588385,
		EndHeight:         17595509,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}
	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-06-30/stat_msg_sender/eth mainnet-msg-17588385-2023.06.30")
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}
	handle.StatMsgSenderHandler(c, list)
	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "eth mainnet-msg-17595509-2023.07.01")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func polygon_stat_msg_and_hist(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "polygon mainnet",
		Endpoint:          "https://polygon-rpc.com",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            []string{"0x81f0d785F82630Bbe6639B7eB07EAB21E2992293", "0xdB6fb08DD8Ce406DA8Ff53FAe65Bd374e3d68681"},
		Mailbox:           "",
		MailerGreenfield:  []string{"0x3fFb989765868CA9f75c90F94A7A17AaB2a981D7", "0x8163A9B0901F63c27471b4d051b7250ECDDD362d"},
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       44495581,
		EndHeight:         44535081,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}
	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-06-30/stat_msg_sender/polygon mainnet-msg-44495581-2023.06.30")
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}
	handle.StatMsgSenderHandler(c, list)
	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "polygon mainnet-msg-44535081-2023.07.01")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func coredao_stat_msg_and_hist(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "core dao",
		Endpoint:          "https://rpc.coredao.org",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            []string{"0x373d2c6dAbb65F1F85a7bC058664b205da0B0824"},
		Mailbox:           "",
		MailerGreenfield:  []string{"0xa0c6ED9Dd8431c40C74BCeEAdA5DaC996dB51f48"},
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       4931307,
		EndHeight:         4960083,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}
	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-06-30/stat_msg_sender/core dao-msg-4931307-2023.06.30")
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}
	handle.StatMsgSenderHandler(c, list)
	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "core dao-msg-4960083-2023.07.01")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func celo_stat_msg_and_hist(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "celo",
		Endpoint:          "https://rpc.ankr.com/celo",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            []string{"0x468Bd5792543b47aaeF12C16E668E7Fd3Fe05872"},
		Mailbox:           "",
		MailerGreenfield:  []string{"0xEae5c76B250206140a27CA0B75DEB28D7D7EAd9B"},
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       20082986,
		EndHeight:         20100265,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}
	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-06-30/stat_msg_sender/celo-msg-20082986-2023.06.30")
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}
	handle.StatMsgSenderHandler(c, list)
	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "celo-msg-20100265-2023.07.01")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func eth_stat_nftbridge_sender(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "eth mainnet",
		Endpoint:          "https://eth-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "0x1e40CD8569F3c91F5101d54AE01a75574a9ccE60",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       16950603,
		EndHeight:         17162285,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}
	list := &sync.Map{}
	handle.StatNftBridgeHandler(c, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "eth mainnet-stat_sender-17162285-2023.05.15")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func bsc_stat_nftbridge_sender(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "0xE09828f0DA805523878Be66EA2a70240d312001e",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}
	list := &sync.Map{}
	handle.StatNftBridgeHandler(c, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-stat_sender-27815520-2023.05.15")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func polygon_stat_nftbridge_sender_zklightclient(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "polygon mainnet",
		Endpoint:          "https://polygon-rpc.com",
		ZkBridge:          "",
		NftBridge:         "0x2E953a70C37E8CB4553DAe1F5760128237c8820D",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "zklightclient",
		Address:        "0x6b0C248679F493481411a0A14cd5FC2DBBe8Ab02",
		StartHeight:    43398055,
		StartTimestamp: 0,
		EndHeight:      43537371,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "polygon mainnet-zklightclient-stat_sender-43537371-2023.05.26")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func bsc_stat_nftbridge_sender_zklightclient(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "0xE09828f0DA805523878Be66EA2a70240d312001e",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "zklightclient",
		Address:        "0xD2cCC9EE7Ea2ccd154c727A46D475ddA49E99852",
		StartHeight:    28714086,
		StartTimestamp: 0,
		EndHeight:      28820664,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-zklightclient-stat_sender-28820664-2023.05.26")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func bsc_stat_nftbridge_sender_sbt(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://rpc.ankr.com/bsc/9b4c185910e116c4c899f349e1616730dbe94ec45dee6c5ffaa7e1f1a64c168d",
		ZkBridge:          "",
		NftBridge:         "0xE09828f0DA805523878Be66EA2a70240d312001e",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "sbt",
		Address:        "0x8FC516dCdcC1f25F9c1352fDBdc8F3b4e164e596",
		Bridges:        "0x4d4b02D4d4188A1d0Cf3D8290e9481321B94d864",
		StartHeight:    30717202,
		StartTimestamp: 0,
		EndHeight:      30744585,
		EndTimestamp:   0,
	}

	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-08-10/stat_nftbridge_sender/bsc mainnet-sbt-stat_sender-30717202-2023.08.10")
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-sbt-stat_sender-30744585-2023.08.11")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func eth_stat_nftbridge_sender_zklightclient(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "eth mainnet",
		Endpoint:          "https://eth-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "0x1e40CD8569F3c91F5101d54AE01a75574a9ccE60",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       16950603,
		EndHeight:         17162285,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "zklightclient",
		Address:        "0x4D8910d4AbA36f856D5C1aD9f43d667E54E0A964",
		StartHeight:    17384358,
		StartTimestamp: 0,
		EndHeight:      17410645,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "eth mainnet-zklightclient-stat_sender-17410645-2023.05.26")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_polygon_stat_nftbridge_sender_zklightclient(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "polygon mainnet",
		Endpoint:          "https://polygon-rpc.com",
		ZkBridge:          "",
		NftBridge:         "0xFFdF4Fe05899C4BdB1676e958FA9F21c19ECB9D5",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       44712756,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "zklightclient",
		Address:        "0x6b0C248679F493481411a0A14cd5FC2DBBe8Ab02",
		StartHeight:    44712756,
		StartTimestamp: 0,
		EndHeight:      45047749,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 polygon mainnet-zklightclient-stat_sender-45047749-2023.07.14")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_bsc_stat_nftbridge_sender_zklightclient(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "0x3668c325501322CEB5a624E95b9E16A019cDEBe8",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "zklightclient",
		Address:        "0xD2cCC9EE7Ea2ccd154c727A46D475ddA49E99852",
		StartHeight:    29697674,
		StartTimestamp: 0,
		EndHeight:      29940653,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 bsc mainnet-zklightclient-stat_sender-29940653-2023.07.14")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_eth_stat_nftbridge_sender_zklightclient(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "eth mainnet",
		Endpoint:          "https://eth-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "0xaD4294BCfDEa6CEb2A4158a38A1a84a6C1A04052",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       16950603,
		EndHeight:         17162285,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "zklightclient",
		Address:        "0x4D8910d4AbA36f856D5C1aD9f43d667E54E0A964",
		StartHeight:    17627819,
		StartTimestamp: 0,
		EndHeight:      17688001,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 eth mainnet-zklightclient-stat_sender-17688001-2023.07.14")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_coredao_stat_nftbridge_sender_alpha(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "core dao",
		Endpoint:          "https://rpc.coredao.org",
		ZkBridge:          "",
		NftBridge:         "0x3701c5897710F16F1f75c6EaE258bf11Ee189a5d",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        116,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "alpha",
		Address:        "0x61DFDbcC65DaF1F60fB1DbE703D84940dA28526c",
		StartHeight:    5046483,
		StartTimestamp: 0,
		EndHeight:      5334348,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 core dao-alpha-stat_sender-5334348-2023.07.14")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_eth_stat_nftbridge_sender_alpha(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "eth mainnet",
		Endpoint:          "https://eth-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "0xaD4294BCfDEa6CEb2A4158a38A1a84a6C1A04052",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       16950603,
		EndHeight:         17162285,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "alpha",
		Address:        "0x0154494A8f48C54772642055B6669196d9749dED",
		StartHeight:    17627819,
		StartTimestamp: 0,
		EndHeight:      17688001,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 eth mainnet-alpha-stat_sender-5334348-2023.07.14")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_polygon_stat_nftbridge_sender_alpha(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "polygon mainnet",
		Endpoint:          "https://polygon-rpc.com",
		ZkBridge:          "",
		NftBridge:         "0xFFdF4Fe05899C4BdB1676e958FA9F21c19ECB9D5",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       44712756,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "alpha",
		Address:        "0x9d5d479a84f3358e8e27afe056494bd2da239acd",
		StartHeight:    44712756,
		StartTimestamp: 0,
		EndHeight:      45047749,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 polygon mainnet-alpha-stat_sender-45047749-2023.07.14")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_bsc_stat_nftbridge_sender_alpha(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "0x3668c325501322CEB5a624E95b9E16A019cDEBe8",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "alpha",
		Address:        "0xC5ae0d15593316e0cC905840eD2dE83E2DD4EA9E",
		StartHeight:    29697674,
		StartTimestamp: 0,
		EndHeight:      29940653,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 bsc mainnet-alpha-stat_sender-29940653-2023.07.14")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_bsc_stat_nftbridge_sender_lubun_loyalty(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "0x3668c325501322CEB5a624E95b9E16A019cDEBe8",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "luban_loyalty",
		Address:        "0x9885C17Dd44c00C37B98F510cdff099EfF437dcE",
		StartHeight:    29697674,
		StartTimestamp: 0,
		EndHeight:      29940653,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 bsc mainnet-luban_loyalty-stat_sender-29940653-2023.07.14")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func bsc_stat_nftbridge_sender_lubun_loyalty(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "0xE09828f0DA805523878Be66EA2a70240d312001e",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "luban_loyalty",
		Address:        "0x9885C17Dd44c00C37B98F510cdff099EfF437dcE",
		StartHeight:    29028650,
		StartTimestamp: 0,
		EndHeight:      29165411,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-luban_loyalty-stat_sender-29165411-2023.06.17")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func coredao_stat_nftbridge_sender_alpha(wg *sync.WaitGroup) {
	defer wg.Done()
	c := utils.Chain{
		Name:              "core dao",
		Endpoint:          "https://rpc.coredao.org",
		ZkBridge:          "",
		NftBridge:         "0x5c5979832A60c17bB06676fa906bEdD1A013e18c",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        116,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	t := utils.Token{
		Name:           "alpha",
		Address:        "0x61DFDbcC65DaF1F60fB1DbE703D84940dA28526c",
		StartHeight:    4619130,
		StartTimestamp: 0,
		EndHeight:      4732291,
		EndTimestamp:   0,
	}

	list := &sync.Map{}
	handle.StatNftBridgeForTokenHandler(c, t, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "core dao-alpha-stat_sender-4732291-2023.06.23")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func stat_cobee() {
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

		fmt.Printf("chain: bsc, height: %v ~ %v\n", curHeight, nextHeight)

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
			fmt.Printf("chain: bsc, try: %v, get logs error :%v\n", try, err)
			time.Sleep(5 * time.Second)
			logs, err = c.FilterLogs(ctx, query)
		}

		if err != nil || try > 10 {
			fmt.Printf("chain: bsc, try: %v, get logs error :%v\n", try, err)
		}

		fmt.Printf("chain: bsc, logs len: %v\n", len(logs))

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

func bsc_stat_nftbridge_sender_for_receiver_opbnb_chain(wg *sync.WaitGroup) {
	defer wg.Done()

	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "0xE09828f0DA805523878Be66EA2a70240d312001e",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         27815520,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	sc := utils.StatChain{
		Name:           "opbnb",
		ZkAppChainId:   []uint64{116},
		StartHeight:    29259397,
		StartTimestamp: 0,
		EndHeight:      29280814,
		EndTimestamp:   0,
		StatContract: []utils.StatContract{
			{
				Name:        "opbnb",
				Addr:        "0x9c614a8E5a23725214024d2C3633BE30D44806f9",
				StartHeight: 0,
				EndHeight:   0,
			},
		},
	}

	list := &sync.Map{}
	handle.StatNftBridgeForReceiverHandler(c, list, sc)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "bsc mainnet-linea testnet-stat_sender-29280814-2023.05.24")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func eth_stat_nftbridge_sender_for_receiver_opbnb_chain(wg *sync.WaitGroup) {
	defer wg.Done()

	c := utils.Chain{
		Name:              "eth mainnet",
		Endpoint:          "https://eth-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		NftBridge:         "",
		Tokens:            nil,
		Mailer:            []string{"0x1Fbc29bD726257EfE5555D47e71030a32426cEbA"},
		Mailbox:           "",
		MailerGreenfield:  []string{"0xB88fE84F922Caf34184EfEbD4146088f5bE0200B"},
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       17519216,
		EndHeight:         17524328,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	sc := utils.StatChain{
		Name:           "opbnb",
		ZkAppChainId:   []uint64{116},
		StartHeight:    17519216,
		StartTimestamp: 0,
		EndHeight:      17524328,
		EndTimestamp:   0,
		StatContract: []utils.StatContract{
			{
				Name:        "opbnb",
				Addr:        "0xcB6cA8eE67E90115dbCe8b0a990443bF7015ceD3",
				StartHeight: 0,
				EndHeight:   0,
			},
		},
	}

	list := &sync.Map{}
	handle.StatNftBridgeForReceiverHandler(c, list, sc)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "eth mainnet-opbnb-opbnb-stat_sender-17524328-2023.06.21")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func polygon_stat_nftbridge_sender_for_receiver_opbnb_chain(wg *sync.WaitGroup) {
	defer wg.Done()

	c := utils.Chain{
		Name:              "polygon mainnet",
		Endpoint:          "https://polygon-rpc.com",
		ZkBridge:          "",
		NftBridge:         "0x2E953a70C37E8CB4553DAe1F5760128237c8820D",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       43386770,
		EndHeight:         43956704,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	sc := utils.StatChain{
		Name:           "opbnb",
		ZkAppChainId:   []uint64{116},
		StartHeight:    44120096,
		StartTimestamp: 0,
		EndHeight:      44300663,
		EndTimestamp:   0,
		StatContract: []utils.StatContract{
			{
				Name:        "opbnb",
				Addr:        "0xfeb105763753e9d26DfD4aae1Ed368aa7cC18260",
				StartHeight: 0,
				EndHeight:   0,
			},
		},
	}

	list := &sync.Map{}
	handle.StatNftBridgeForReceiverHandler(c, list, sc)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "polygon mainnet-opbnb-opbnb-stat_sender-44300663-2023.06.25")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func coredao_stat_nftbridge_sender_for_receiver_opbnb_chain(wg *sync.WaitGroup) {
	defer wg.Done()

	c := utils.Chain{
		Name:              "core dao",
		Endpoint:          "https://rpc.coredao.org",
		ZkBridge:          "",
		NftBridge:         "0x5c5979832A60c17bB06676fa906bEdD1A013e18c",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         4631608,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	sc := utils.StatChain{
		Name:           "opbnb",
		ZkAppChainId:   []uint64{116},
		StartHeight:    4651650,
		StartTimestamp: 0,
		EndHeight:      4787313,
		EndTimestamp:   0,
		StatContract: []utils.StatContract{
			{
				Name:        "opbnb",
				Addr:        "0x0f83DA622E36Ee42cfeB222257E1baF20E16a491",
				StartHeight: 0,
				EndHeight:   0,
			},
		},
	}

	list := &sync.Map{}
	handle.StatNftBridgeForReceiverHandler(c, list, sc)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "core dao-opbnb-opbnb-stat_sender-4787313-2023.06.25")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_bsc_stat_nftbridge_sender_for_receiver_opbnb_chain(wg *sync.WaitGroup) {
	defer wg.Done()

	c := utils.Chain{
		Name:              "bsc mainnet",
		Endpoint:          "https://bsc-mainnet.nodereal.io/v1/29c57e58ee374baca2e9ad15a5c8273e",
		ZkBridge:          "",
		L0NftBridge:       "0x3668c325501322CEB5a624E95b9E16A019cDEBe8",
		NftBridge:         "0xE09828f0DA805523878Be66EA2a70240d312001e",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         29768341,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	scc := utils.StatContract{
		Name:        "onft",
		Addr:        "0x87a218Ae43C136B3148A45EA1A282517794002c8",
		StartHeight: 0,
		EndHeight:   0,
	}

	sc := utils.StatChain{
		Name:           "core dao|polygon mainnet|combo testnet|celo|opbnb",
		L0AppChainID:   []uint64{153, 109, 125},
		ZkAppChainId:   []uint64{114, 116},
		StartHeight:    30285259,
		StartTimestamp: 0,
		EndHeight:      30314000,
		EndTimestamp:   0,
		StatContract: []utils.StatContract{
			scc,
		},
	}

	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-07-26/stat_nftbridge_receiver_chain_sender/l0 bsc mainnet-core dao|polygon mainnet|combo testnet|celo-onft-stat_sender-30285259-2023.07.26")
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}
	handle.StatNftBridgeForReceiverAndContractHandler(c, list, sc, scc)
	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 bsc mainnet-core dao|polygon mainnet|combo testnet|celo|opbnb-onft-stat_sender-30314000-2023.07.27")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_polygon_stat_nftbridge_sender_for_receiver_opbnb_chain(wg *sync.WaitGroup) {
	defer wg.Done()

	c := utils.Chain{
		Name:              "polygon mainnet",
		Endpoint:          "https://rpc.ankr.com/polygon/9b4c185910e116c4c899f349e1616730dbe94ec45dee6c5ffaa7e1f1a64c168d",
		ZkBridge:          "",
		L0NftBridge:       "0xFFdF4Fe05899C4BdB1676e958FA9F21c19ECB9D5",
		NftBridge:         "0x2E953a70C37E8CB4553DAe1F5760128237c8820D",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       43386770,
		EndHeight:         44809953,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	scc := utils.StatContract{
		Name:        "onft",
		Addr:        "0x141A1fb33683C304DA7C3fe6fC6a49B5C0c2dC42",
		StartHeight: 0,
		EndHeight:   0,
	}

	sc := utils.StatChain{
		Name:           "core dao|bsc mainnet|combo testnet|celo|opbnb",
		L0AppChainID:   []uint64{153, 102, 125},
		ZkAppChainId:   []uint64{114, 116},
		StartHeight:    45519365,
		StartTimestamp: 0,
		EndHeight:      45558766,
		EndTimestamp:   0,
		StatContract: []utils.StatContract{
			scc,
		},
	}

	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-07-26/stat_nftbridge_receiver_chain_sender/l0 polygon mainnet-core dao|bsc mainnet|combo testnet|celo-onft-stat_sender-45519365-2023.07.26")
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}
	handle.StatNftBridgeForReceiverAndContractHandler(c, list, sc, scc)
	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 polygon mainnet-core dao|bsc mainnet|combo testnet|celo|opbnb-onft-stat_sender-45558766-2023.07.27")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_coredao_stat_nftbridge_sender_for_receiver_opbnb_chain(wg *sync.WaitGroup) {
	defer wg.Done()

	c := utils.Chain{
		Name:              "core dao",
		Endpoint:          "https://rpc.coredao.org",
		ZkBridge:          "",
		L0NftBridge:       "0x3701c5897710F16F1f75c6EaE258bf11Ee189a5d",
		NftBridge:         "0x5c5979832A60c17bB06676fa906bEdD1A013e18c",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         4631608,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	scc := utils.StatContract{
		Name:        "onft",
		Addr:        "0x36e5121618c0Af89E81AcD33B6b8F5CF5cDD961a",
		StartHeight: 0,
		EndHeight:   0,
	}

	sc := utils.StatChain{
		Name:           "polygon mainnet|bsc mainnet|combo testnet|celo|opbnb",
		L0AppChainID:   []uint64{109, 102, 125},
		ZkAppChainId:   []uint64{114, 116},
		StartHeight:    5679799,
		StartTimestamp: 0,
		EndHeight:      5708599,
		EndTimestamp:   0,
		StatContract: []utils.StatContract{
			scc,
		},
	}

	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-07-26/stat_nftbridge_receiver_chain_sender/l0 core dao-polygon mainnet|bsc mainnet|combo testnet|celo-onft-stat_sender-5679799-2023.07.26")
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}
	handle.StatNftBridgeForReceiverAndContractHandler(c, list, sc, scc)
	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 core dao-polygon mainnet|bsc mainnet|combo testnet|celo|opbnb-onft-stat_sender-5708599-2023.07.27")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_celo_stat_nftbridge_sender_for_receiver_opbnb_chain(wg *sync.WaitGroup) {
	defer wg.Done()

	c := utils.Chain{
		Name:              "celo",
		Endpoint:          "https://rpc.ankr.com/celo/9b4c185910e116c4c899f349e1616730dbe94ec45dee6c5ffaa7e1f1a64c168d",
		ZkBridge:          "",
		L0NftBridge:       "0xe47b0a5F2444F9B360Bd18b744B8D511CfBF98c6",
		NftBridge:         "0x24339b7f8d303527C8681382AbD4Ec299757aF63",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       26956110,
		EndHeight:         4631608,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
	}

	scc := utils.StatContract{
		Name:        "onft",
		Addr:        "0xb404e5233aB7E426213998C025f05EaBaBD41Da6",
		StartHeight: 0,
		EndHeight:   0,
	}

	sc := utils.StatChain{
		Name:           "bsc mainnet|core dao|polygon mainnet|combo testnet|opbnb",
		L0AppChainID:   []uint64{102, 153, 109},
		ZkAppChainId:   []uint64{114, 116},
		StartHeight:    20532247,
		StartTimestamp: 0,
		EndHeight:      20549527,
		EndTimestamp:   0,
		StatContract: []utils.StatContract{
			scc,
		},
	}

	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-07-26/stat_nftbridge_receiver_chain_sender/l0 celo-bsc mainnet|core dao|polygon mainnet|combo testnet-onft-stat_sender-20532247-2023.07.26")
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}
	//list := &sync.Map{}
	handle.StatNftBridgeForReceiverAndContractHandler(c, list, sc, scc)
	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 celo-bsc mainnet|core dao|polygon mainnet|combo testnet|opbnb-onft-stat_sender-20549527-2023.07.27")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func l0_mantle_stat_nftbridge_receiver_nft(wg *sync.WaitGroup) {
	defer wg.Done()

	c := utils.Chain{
		Name:              "l0 mantle",
		Endpoint:          "https://rpc.mantle.xyz",
		ZkBridge:          "",
		NftBridge:         "",
		L0NftBridge:       "0x5E6D2a7aA45277cA3037FebA4a30CbB8f8caD3B9",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        0,
		ChainID:           0,
		StartHeight:       506368,
		EndHeight:         751856,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
		Merge:             false,
		StatChains:        nil,
		StatContract: []utils.StatContract{
			{
				Name:        "bsc onft",
				Addr:        "0x03dDC4b60D6bBf399A8397d73462060FdFb83476",
				StartHeight: 0,
				EndHeight:   0,
			},
			{
				Name:        "polygon onft",
				Addr:        "0xCf995797cB2E65Cc290e084f0127B1e8Ebc692c8",
				StartHeight: 0,
				EndHeight:   0,
			},
			{
				Name:        "core dao onft",
				Addr:        "0xE09828f0DA805523878Be66EA2a70240d312001e",
				StartHeight: 0,
				EndHeight:   0,
			},
			{
				Name:        "celo onft",
				Addr:        "0xA98163227B85CcC765295Ce5C18E8aAD663De147",
				StartHeight: 0,
				EndHeight:   0,
			},
		},
	}
	list := &sync.Map{}
	handle.StatNftBridgeForReceiverNftHandler(c, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "l0 mantle-bsc onft|polygon onft|core dao onft|celo onft|-stat_receiver-751856-2023.08.05")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func combo_to_opbnb_claim_origin(wg *sync.WaitGroup) {
	defer wg.Done()

	srcC := utils.Chain{
		Name:        "combo testnet",
		Endpoint:    "https://test-rpc.combonetwork.io",
		ZkBridge:    "",
		NftBridge:   "0x8E69018871C8d30788Dcaf8ee2c97D8F8ecE5d8D",
		L0NftBridge: "",
		Tokens: []utils.Token{
			{
				Name:           "combo nft",
				Address:        "0x20Cb10B8f601d4B2C62962BB938554F3824e24f3",
				Bridges:        "0x8E69018871C8d30788Dcaf8ee2c97D8F8ecE5d8D",
				StartHeight:    0,
				StartTimestamp: 0,
				EndHeight:      0,
				EndTimestamp:   0,
			},
		},
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        114,
		ChainID:           91715,
		StartHeight:       8383384,
		EndHeight:         10688712,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
		Merge:             false,
		StatChains:        nil,
		StatContract:      nil,
	}

	dstC := utils.Chain{
		Name:              "opbnb testnet",
		Endpoint:          "https://opbnb-testnet-rpc.bnbchain.org",
		ZkBridge:          "",
		NftBridge:         "0xbA09FDF988d4113460dbdD96FEfD33C8400E4E0D",
		L0NftBridge:       "",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        116,
		ChainID:           5611,
		StartHeight:       3634365,
		EndHeight:         5783895,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
		Merge:             false,
		StatChains:        nil,
		StatContract:      nil,
	}

	scsn := utils.StatClaimNftSender{
		Src: srcC,
		Dst: dstC,
	}
	//list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-08-21/stat_claim_nft_sender/combo testnet-opbnb testnet-combo nft|-null|-stat_claim_nft_sender-10602312-5710458-2023.08.21")
	//if err != nil {
	//	utils.Logger.Errorln(err)
	//	panic(err)
	//}
	list := &sync.Map{}
	handle.ClaimNftNftHandler(scsn, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "combo testnet-opbnb testnet-combo nft|-null|-stat_claim_nft_sender-10688712-5783895-2023.08.22")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func combo_to_opbnb_claim2(wg *sync.WaitGroup) {
	defer wg.Done()

	srcC := utils.Chain{
		Name:        "combo testnet",
		Endpoint:    "https://test-rpc.combonetwork.io",
		ZkBridge:    "",
		NftBridge:   "0x8E69018871C8d30788Dcaf8ee2c97D8F8ecE5d8D",
		L0NftBridge: "",
		Tokens: []utils.Token{
			{
				Name:           "combo nft",
				Address:        "0x20Cb10B8f601d4B2C62962BB938554F3824e24f3",
				Bridges:        "0x8E69018871C8d30788Dcaf8ee2c97D8F8ecE5d8D",
				StartHeight:    0,
				StartTimestamp: 0,
				EndHeight:      0,
				EndTimestamp:   0,
			},
		},
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        114,
		ChainID:           91715,
		StartHeight:       9651911,
		EndHeight:         10956289,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
		Merge:             false,
		StatChains:        nil,
		StatContract:      nil,
	}

	dstC := utils.Chain{
		Name:              "opbnb testnet",
		Endpoint:          "https://opbnb-testnet.nodereal.io/v1/0daebfa60722447594b98bafcacf73bc",
		ZkBridge:          "",
		NftBridge:         "0xbA09FDF988d4113460dbdD96FEfD33C8400E4E0D",
		L0NftBridge:       "",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        116,
		ChainID:           5611,
		StartHeight:       4747094,
		EndHeight:         6051494,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
		Merge:             false,
		StatChains:        nil,
		StatContract:      nil,
	}

	scsn := utils.StatClaimNftSender{
		Src: srcC,
		Dst: dstC,
	}
	//list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-08-22/stat_claim_nft_sender/combo testnet-opbnb testnet-combo nft|-null|-stat_claim_nft_sender-10688712-5783895-2023.08.22")
	//if err != nil {
	//	utils.Logger.Errorln(err)
	//	panic(err)
	//}
	list := &sync.Map{}
	handle.ClaimNftNftHandler(scsn, list)
	_, err := utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "combo testnet-opbnb testnet-combo nft|-null|-stat_claim_nft_sender-10956289-6051494-2023.08.25")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}

func combo_to_opbnb_claim3(wg *sync.WaitGroup) {
	defer wg.Done()

	srcC := utils.Chain{
		Name:        "combo testnet",
		Endpoint:    "https://test-rpc.combonetwork.io",
		ZkBridge:    "",
		NftBridge:   "0x8E69018871C8d30788Dcaf8ee2c97D8F8ecE5d8D",
		L0NftBridge: "",
		Tokens: []utils.Token{
			{
				Name:           "combo nft",
				Address:        "0x20Cb10B8f601d4B2C62962BB938554F3824e24f3",
				Bridges:        "0x8E69018871C8d30788Dcaf8ee2c97D8F8ecE5d8D",
				StartHeight:    0,
				StartTimestamp: 0,
				EndHeight:      0,
				EndTimestamp:   0,
			},
		},
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        114,
		ChainID:           91715,
		StartHeight:       9712344,
		EndHeight:         9712344,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
		Merge:             false,
		StatChains:        nil,
		StatContract:      nil,
	}

	dstC := utils.Chain{
		Name:              "opbnb testnet",
		Endpoint:          "https://opbnb-testnet.nodereal.io/v1/0daebfa60722447594b98bafcacf73bc",
		ZkBridge:          "",
		NftBridge:         "0xbA09FDF988d4113460dbdD96FEfD33C8400E4E0D",
		L0NftBridge:       "",
		Tokens:            nil,
		Mailer:            nil,
		Mailbox:           "",
		MailerGreenfield:  nil,
		MailboxGreenfield: "",
		AppChainID:        116,
		ChainID:           5611,
		StartHeight:       6561494,
		EndHeight:         6647895,
		WaitNum:           0,
		WaitSecond:        0,
		Testnet:           false,
		Merge:             false,
		StatChains:        nil,
		StatContract:      nil,
	}

	scsn := utils.StatClaimNftSender{
		Src: srcC,
		Dst: dstC,
	}
	list, _, err := utils.GetYesterdayCsvData("/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv3/2023-08-31/stat_claim_nft_sender/combo testnet-opbnb testnet-combo nft|-null|-stat_claim_nft_sender-11466311-6561494-2023.08.31")
	if err != nil {
		utils.Logger.Errorln(err)
		panic(err)
	}
	//list := &sync.Map{}
	handle.ClaimNftNftHandler(scsn, list)
	_, err = utils.ToCsv(list, "/home/m/go/src/yunyc12345/statistics-of-bridge-data/csv/", "combo testnet-opbnb testnet-combo nft|-null|-stat_claim_nft_sender-11552712-6647895-2023.09.01")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return
}
