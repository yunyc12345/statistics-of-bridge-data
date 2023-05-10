package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"yunyc12345/statistics-of-bridge-data/contracts"
)

type Void struct{}

var Member Void

type Chain struct {
	Name              string   `json:"name"`
	Endpoint          string   `json:"endpoint"`
	ZkBridge          []string `json:"zk_bridge"`
	NftBridge         []string `json:"nft_bridge"`
	NftErc721         []string `json:"nft_erc721"`
	Mailer            []string `json:"mailer"`
	Mailbox           []string `json:"mailbox"`
	MailerGreenfield  []string `json:"mailer_greenfield"`
	MailboxGreenfield []string `json:"mailbox_greenfield"`
	AppChainID        int      `json:"app_chain_id"`
	ChainID           int      `json:"chain_id"`
	StartHeight       int      `json:"start_height"`
	EndHeight         int      `json:"end_height"`
	WaitNum           int      `json:"wait_num"`
	WaitSecond        int      `json:"wait_second"`
	Testnet           bool     `json:"testnet"`
}

type Config struct {
	TestnetNftbridgeSendersStatistics []Chain `json:"testnet_nftbridge_senders_statistics"`
	MainnetNfterc721SendersStatistics []Chain `json:"mainnet_nfterc721_senders_statistics"`
	MainnetNftbridgeSendersStatistics []Chain `json:"mainnet_nftbridge_senders_statistics"`
	MainnetMsgSendersStatistics       []Chain `json:"mainnet_msg_senders_statistics"`
}

type ChainData struct {
	Cli       *ethclient.Client
	ZkBridge  []*contracts.ZKBridage
	Mailer    []*contracts.Mailer
	GfMailer  []*contracts.Mailer
	Nft       []*contracts.Nft
	NftBridge []*contracts.NftBridge
	Info      *Chain
}

func InitGlobalCliMapAndZKMap(chain *Chain) (*ChainData, error) {
	c, err := ethclient.Dial(chain.Endpoint)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	zkBridges := make([]*contracts.ZKBridage, 0)
	for _, s := range chain.ZkBridge {
		zk, err := contracts.NewZKBridage(common.HexToAddress(s), c)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
		zkBridges = append(zkBridges, zk)
	}

	mailers := make([]*contracts.Mailer, 0)
	for _, s := range chain.Mailer {
		mailer, err := contracts.NewMailer(common.HexToAddress(s), c)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
		mailers = append(mailers, mailer)
	}

	gfMailers := make([]*contracts.Mailer, 0)
	for _, s := range chain.MailerGreenfield {
		gfMailer, err := contracts.NewMailer(common.HexToAddress(s), c)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
		gfMailers = append(gfMailers, gfMailer)
	}

	nfts := make([]*contracts.Nft, 0)
	for _, s := range chain.NftErc721 {
		nft, err := contracts.NewNft(common.HexToAddress(s), c)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
		nfts = append(nfts, nft)
	}

	nftBridges := make([]*contracts.NftBridge, 0)
	for _, s := range chain.NftBridge {
		nftBridge, err := contracts.NewNftBridge(common.HexToAddress(s), c)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
		nftBridges = append(nftBridges, nftBridge)
	}

	cd := ChainData{
		Cli:       c,
		ZkBridge:  zkBridges,
		Mailer:    mailers,
		GfMailer:  gfMailers,
		Info:      chain,
		Nft:       nfts,
		NftBridge: nftBridges,
	}

	return &cd, nil
}
