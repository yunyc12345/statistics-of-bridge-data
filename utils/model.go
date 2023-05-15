package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"yunyc12345/statistics-of-bridge-data/contracts"
)

type Void struct{}

var Member Void

type Chain struct {
	Name              string  `json:"name"`
	Endpoint          string  `json:"endpoint"`
	ZkBridge          string  `json:"zk_bridge"`
	NftBridge         string  `json:"nft_bridge"`
	Tokens            []Token `json:"tokens"`
	Mailer            string  `json:"mailer"`
	Mailbox           string  `json:"mailbox"`
	MailerGreenfield  string  `json:"mailer_greenfield"`
	MailboxGreenfield string  `json:"mailbox_greenfield"`
	AppChainID        uint64  `json:"app_chain_id"`
	ChainID           uint64  `json:"chain_id"`
	StartHeight       uint64  `json:"start_height"`
	EndHeight         uint64  `json:"end_height"`
	WaitNum           uint64  `json:"wait_num"`
	WaitSecond        uint64  `json:"wait_second"`
	Testnet           bool    `json:"testnet"`
}

func (c *Chain) ToMsgNew(startHeight, endHeight uint64) Chain {
	return Chain{
		Name:              c.Name,
		Endpoint:          c.Endpoint,
		ZkBridge:          c.ZkBridge,
		NftBridge:         c.NftBridge,
		Tokens:            c.Tokens,
		Mailer:            c.Mailer,
		Mailbox:           c.Mailbox,
		MailerGreenfield:  c.MailerGreenfield,
		MailboxGreenfield: c.MailboxGreenfield,
		AppChainID:        c.AppChainID,
		ChainID:           c.ChainID,
		StartHeight:       startHeight,
		EndHeight:         endHeight,
		WaitNum:           c.WaitNum,
		WaitSecond:        c.WaitSecond,
		Testnet:           false,
	}
}

func (c *Chain) ToTokenNew(tokens []Token) Chain {
	return Chain{
		Name:              c.Name,
		Endpoint:          c.Endpoint,
		ZkBridge:          c.ZkBridge,
		NftBridge:         c.NftBridge,
		Tokens:            tokens,
		Mailer:            c.Mailer,
		Mailbox:           c.Mailbox,
		MailerGreenfield:  c.MailerGreenfield,
		MailboxGreenfield: c.MailboxGreenfield,
		AppChainID:        c.AppChainID,
		ChainID:           c.ChainID,
		StartHeight:       c.StartHeight,
		EndHeight:         c.EndHeight,
		WaitNum:           c.WaitNum,
		WaitSecond:        c.WaitSecond,
		Testnet:           false,
	}
}

type Server struct {
	Port        string `json:"port"`
	CsvFilePath string `json:"csv_file_path"`
}

type Config struct {
	//TestnetNftbridgeSendersStatistics []Chain `json:"testnet_nftbridge_senders_statistics"`
	Server                            Server  `json:"server"`
	MainnetNftMinterStatistics        []Chain `json:"mainnet_nft_minter_statistics"`
	MainnetNftbridgeSendersStatistics []Chain `json:"mainnet_nftbridge_senders_statistics"`
	MainnetMsgSendersStatistics       []Chain `json:"mainnet_msg_senders_statistics"`
}

type Token struct {
	Name           string `json:"name"`
	Address        string `json:"address"`
	StartHeight    uint64 `json:"start_height"`
	StartTimestamp uint64 `json:"start_timestamp"`

	EndHeight    uint64 `json:"end_height"`
	EndTimestamp uint64 `json:"end_timestamp"`
}

func (t *Token) NewToken(startHeight, endHeight uint64) Token {
	return Token{
		Name:           t.Name,
		Address:        t.Address,
		StartHeight:    startHeight,
		StartTimestamp: 0,
		EndHeight:      endHeight,
		EndTimestamp:   0,
	}
}

type Info struct {
	ChainName      string `json:"chain_name"`
	AppChainId     uint64 `json:"app_chain_id"`
	TokenName      string `json:"token_name"`
	EndHeight      uint64 `json:"end_height"`
	EndTimestamp   uint64 `json:"end_timestamp"`
	StartHeight    uint64 `json:"start_height"`
	StartTimestamp uint64 `json:"start_timestamp"`
}

type ChainData struct {
	Cli       *ethclient.Client
	ZkBridge  *contracts.ZKBridage
	Mailer    *contracts.Mailer
	GfMailer  *contracts.Mailer
	Alpha     *contracts.Nft
	Cobee     *contracts.Nft
	Loyalty   *contracts.Nft
	NftBridge *contracts.NftBridge
	Info      *Chain
}

type RunParams struct {
	Eth RunParam1 `json:"eth"`
	Bsc RunParam1 `json:"bsc"`
}

type RunParam1 struct {
	ChainName string `json:"chain_name"`
	EndHeight uint64 `json:"end_height"`
}

func InitGlobalCliMapAndZKMap(chain *Chain) (*ChainData, error) {
	c, err := ethclient.Dial(chain.Endpoint)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	zkBridges, err := contracts.NewZKBridage(common.HexToAddress(chain.ZkBridge), c)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	mailer, err := contracts.NewMailer(common.HexToAddress(chain.Mailer), c)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	gfMailer, err := contracts.NewMailer(common.HexToAddress(chain.MailerGreenfield), c)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	nftBridge, err := contracts.NewNftBridge(common.HexToAddress(chain.NftBridge), c)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	var alpha *contracts.Nft
	var loyalty *contracts.Nft
	var cobee *contracts.Nft

	for _, token := range chain.Tokens {
		if token.Name == "loyalty" {
			loyalty, err = contracts.NewNft(common.HexToAddress(token.Address), c)
			if err != nil {
				Logger.Errorln(err)
				return nil, err
			}
		}

		if token.Name == "alpha" {
			alpha, err = contracts.NewNft(common.HexToAddress(token.Address), c)
			if err != nil {
				Logger.Errorln(err)
				return nil, err
			}
		}

		if token.Name == "cobee" {
			cobee, err = contracts.NewNft(common.HexToAddress(token.Address), c)
			if err != nil {
				Logger.Errorln(err)
				return nil, err
			}
		}
	}

	cd := ChainData{
		Cli:       c,
		ZkBridge:  zkBridges,
		Mailer:    mailer,
		GfMailer:  gfMailer,
		Alpha:     alpha,
		Cobee:     cobee,
		Loyalty:   loyalty,
		NftBridge: nftBridge,
		Info:      chain,
	}

	return &cd, nil
}
