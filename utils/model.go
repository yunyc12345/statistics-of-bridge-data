package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"yunyc12345/statistics-of-bridge-data/contracts"
)

type Void struct{}

var Member Void

type Chain struct {
	Name              string `json:"name"`
	Endpoint          string `json:"endpoint"`
	ZkBridge          string `json:"zk_bridge"`
	NftBridge         string `json:"nft_bridge"`
	NftErc721         string `json:"nft_erc721"`
	Mailer            string `json:"mailer"`
	Mailbox           string `json:"mailbox"`
	MailerGreenfield  string `json:"mailer_greenfield"`
	MailboxGreenfield string `json:"mailbox_greenfield"`
	AppChainID        int    `json:"app_chain_id"`
	ChainID           int    `json:"chain_id"`
	StartHeight       int    `json:"start_height"`
	EndHeight         int    `json:"end_height"`
	WaitNum           int    `json:"wait_num"`
	WaitSecond        int    `json:"wait_second"`
	Testnet           bool   `json:"testnet"`
}

type Config struct {
	TestnetNftbridgeSendersStatistics []Chain `json:"testnet_nftbridge_senders_statistics"`
	MainnetNfterc721SendersStatistics []Chain `json:"mainnet_nfterc721_senders_statistics"`
	MainnetNftbridgeSendersStatistics []Chain `json:"mainnet_nftbridge_senders_statistics"`
	MainnetMsgSendersStatistics       []Chain `json:"mainnet_msg_senders_statistics"`
}

type ChainData struct {
	Cli       *ethclient.Client
	ZkBridge  *contracts.ZKBridage
	Mailer    *contracts.Mailer
	GfMailer  *contracts.Mailer
	Nft       *contracts.Nft
	NftBridge *contracts.NftBridge
	Info      *Chain
}

func InitGlobalCliMapAndZKMap(chain *Chain) (*ChainData, error) {
	c, err := ethclient.Dial(chain.Endpoint)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	zk, err := contracts.NewZKBridage(common.HexToAddress(chain.ZkBridge), c)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}
	var mailer *contracts.Mailer
	if chain.Mailer != "" {
		mailer, err = contracts.NewMailer(common.HexToAddress(chain.Mailer), c)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
	}

	var gfMailer *contracts.Mailer
	if chain.MailerGreenfield != "" {
		gfMailer, err = contracts.NewMailer(common.HexToAddress(chain.MailerGreenfield), c)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
	}

	var nft *contracts.Nft
	if chain.NftErc721 != "" {
		nft, err = contracts.NewNft(common.HexToAddress(chain.NftErc721), c)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
	}
	var nftBridge *contracts.NftBridge
	if chain.NftBridge != "" {
		nftBridge, err = contracts.NewNftBridge(common.HexToAddress(chain.NftBridge), c)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
	}

	cd := ChainData{
		Cli:       c,
		ZkBridge:  zk,
		Mailer:    mailer,
		GfMailer:  gfMailer,
		Info:      chain,
		Nft:       nft,
		NftBridge: nftBridge,
	}

	return &cd, nil
}
