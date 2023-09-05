package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"yunyc12345/statistics-of-bridge-data/contracts"
)

type Void struct{}

var Member Void

type Chain struct {
	Name              string         `json:"name"`
	Endpoint          string         `json:"endpoint"`
	ZkBridge          string         `json:"zk_bridge"`
	NftBridge         string         `json:"nft_bridge"`
	L0NftBridge       string         `json:"l0_nft_bridge"`
	Tokens            []Token        `json:"tokens"`
	Mailer            []string       `json:"mailer"`
	Mailbox           string         `json:"mailbox"`
	MailerGreenfield  []string       `json:"mailer_greenfield"`
	MailboxGreenfield string         `json:"mailbox_greenfield"`
	AppChainID        uint64         `json:"app_chain_id"`
	ChainID           uint64         `json:"chain_id"`
	StartHeight       uint64         `json:"start_height"`
	EndHeight         uint64         `json:"end_height"`
	WaitNum           uint64         `json:"wait_num"`
	WaitSecond        uint64         `json:"wait_second"`
	Testnet           bool           `json:"testnet"`
	Merge             bool           `json:"merge"`
	StatChains        []StatChain    `json:"stat_chains"`
	StatContract      []StatContract `json:"stat_contract"`
}

type StatChain struct {
	Name           string   `json:"name"`
	L0AppChainID   []uint64 `json:"l0_app_chain_id"`
	ZkAppChainId   []uint64 `json:"zk_app_chain_id"`
	StartHeight    uint64   `json:"start_height"`
	StartTimestamp uint64   `json:"start_timestamp"`

	EndHeight    uint64 `json:"end_height"`
	EndTimestamp uint64 `json:"end_timestamp"`

	StatContract []StatContract `json:"stat_contract"`
}

type StatContract struct {
	Name        string `json:"name"`
	Addr        string `json:"addr"`
	StartHeight uint64 `json:"start_height"`
	EndHeight   uint64 `json:"end_height"`
}

func (sc *StatChain) ToNewStatChain(startHeight, endHeight uint64) StatChain {
	return StatChain{
		Name:           sc.Name,
		L0AppChainID:   sc.L0AppChainID,
		ZkAppChainId:   sc.ZkAppChainId,
		StartHeight:    startHeight,
		StartTimestamp: 0,
		EndHeight:      endHeight,
		EndTimestamp:   0,
	}
}

func (sc *StatChain) ToNewStatChain2(startHeight, endHeight uint64, s StatContract) StatChain {
	return StatChain{
		Name:           sc.Name,
		L0AppChainID:   sc.L0AppChainID,
		ZkAppChainId:   sc.ZkAppChainId,
		StartHeight:    startHeight,
		StartTimestamp: 0,
		EndHeight:      endHeight,
		EndTimestamp:   0,
		StatContract:   []StatContract{s},
	}
}

func (c *Chain) ToNewChain(startHeight, endHeight uint64) Chain {
	return Chain{
		Name:              c.Name,
		Endpoint:          c.Endpoint,
		ZkBridge:          c.ZkBridge,
		NftBridge:         c.NftBridge,
		L0NftBridge:       c.L0NftBridge,
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
		Merge:             false,
		StatChains:        c.StatChains,
		StatContract:      c.StatContract,
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
		StatChains:        c.StatChains,
	}
}

func (c *Chain) ToStatChainNew(scs []StatChain) Chain {
	return Chain{
		Name:              c.Name,
		Endpoint:          c.Endpoint,
		ZkBridge:          c.ZkBridge,
		NftBridge:         c.NftBridge,
		L0NftBridge:       c.L0NftBridge,
		Tokens:            c.Tokens,
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
		StatChains:        scs,
	}
}

type Server struct {
	Port        string `json:"port"`
	CsvFilePath string `json:"csv_file_path"`
}

type Config struct {
	//TestnetNftbridgeSendersStatistics []Chain `json:"testnet_nftbridge_senders_statistics"`
	Server                            Server               `json:"server"`
	MainnetNftMinterStatistics        []Chain              `json:"mainnet_nft_minter_statistics"`
	MainnetNftbridgeSendersStatistics []Chain              `json:"mainnet_nftbridge_senders_statistics"`
	MainnetMsgSendersStatistics       []Chain              `json:"mainnet_msg_senders_statistics"`
	StatDomains                       []StatDomain         `json:"stat_domains"`
	StatReceiverChains                []Chain              `json:"stat_receiver_chains"`
	StatReceiverNft                   []Chain              `json:"stat_receiver_nft"`
	StatClaimNftSender                []StatClaimNftSender `json:"stat_claim_nft_sender"`
}

func (s *StatClaimNftSender) ToNew(srcC, dstC Chain) StatClaimNftSender {
	return StatClaimNftSender{
		Src: srcC,
		Dst: dstC,
	}
}

type StatClaimNftSender struct {
	Src Chain `json:"src"`
	Dst Chain `json:"dst"`
}

type Token struct {
	Name           string `json:"name"`
	Address        string `json:"address"`
	Bridges        string `json:"bridges"`
	StartHeight    uint64 `json:"start_height"`
	StartTimestamp uint64 `json:"start_timestamp"`

	EndHeight    uint64 `json:"end_height"`
	EndTimestamp uint64 `json:"end_timestamp"`
}

func (t *Token) NewToken(startHeight, endHeight uint64) Token {
	return Token{
		Name:           t.Name,
		Address:        t.Address,
		Bridges:        t.Bridges,
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
	Mailer    []*contracts.Mailer
	GfMailer  []*contracts.Mailer
	Alpha     *contracts.Nft
	Cobee     *contracts.Nft
	Loyalty   *contracts.Nft
	NftBridge *contracts.NftBridge
	Info      *Chain
}

type RunParams struct {
	Eth          RunParam1 `json:"eth"`
	Bsc          RunParam1 `json:"bsc"`
	Polygon      RunParam1 `json:"polygon"`
	CoreDao      RunParam1 `json:"core_dao"`
	Celo         RunParam1 `json:"celo"`
	L0Eth        RunParam1 `json:"l0_eth"`
	L0Bsc        RunParam1 `json:"l0_bsc"`
	L0Polygon    RunParam1 `json:"l0_polygon"`
	L0CoreDao    RunParam1 `json:"l0_core_dao"`
	L0Celo       RunParam1 `json:"l0_celo"`
	L0Mantle     RunParam1 `json:"l0_mantle"`
	ComboTestNet RunParam1 `json:"combo_test_net"`
	OpbnbTestNet RunParam1 `json:"opbnb_test_net"`
}

type RunParam1 struct {
	ChainName string `json:"chain_name"`
	EndHeight uint64 `json:"end_height"`
}

type RunParam2 struct {
	ChainName string `json:"chain_name"`
	EndHeight uint64 `json:"end_height"`
	Endpoint  string `json:"endpoint"`
}

type StatDomain struct {
	Url            string `json:"url"`
	Ty             string `json:"ty"`
	StartTimestamp int64  `json:"start_timestamp"`
	EndTimestamp   int64  `json:"end_timestamp"`
}

type StateDomainResp struct {
	Code int `json:"code"`
	Data []struct {
		Id            string `json:"id"`
		DomainName    string `json:"domainName"`
		Address       string `json:"address"`
		SetTime       int64  `json:"setTime"`
		DomainType    string `json:"domainType"`
		SenderAddress string `json:"senderAddress"`
	} `json:"data"`
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
		Mailer:    mailers,
		GfMailer:  gfMailers,
		Alpha:     alpha,
		Cobee:     cobee,
		Loyalty:   loyalty,
		NftBridge: nftBridge,
		Info:      chain,
	}

	return &cd, nil
}
