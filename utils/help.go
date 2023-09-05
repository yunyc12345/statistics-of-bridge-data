package utils

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func JudgHaveHisData(dirPath string) (bool, error) {

	// stat msg 2023-05-11
	statMsgDirPath := dirPath + "/2023-05-11/stat_msg_sender"
	b1, err := DirIsEmpty(statMsgDirPath)
	if err != nil {
		Logger.Errorln(err)
		return false, err
	}

	// stat nft mint 2023-05-11
	statNftMinterDirPath := dirPath + "/2023-05-11/stat_nft_minter"
	b2, err := DirIsEmpty(statNftMinterDirPath)
	if err != nil {
		Logger.Errorln(err)
		return false, err
	}

	// stat nftbridge sender 2023-05-11
	statNftBridgeSenderPath := dirPath + "/2023-05-11/stat_nftbridge_sender"
	b3, err := DirIsEmpty(statNftBridgeSenderPath)
	if err != nil {
		Logger.Errorln(err)
		return false, err
	}

	if !b1 && !b2 && !b3 {
		return true, nil
	} else {
		return false, nil
	}
}

func CreateCsvDir(dirPath string) error {
	b, err := PathExists(dirPath)
	if err != nil {
		Logger.Errorln(err)
		return err
	}

	if !b {
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			Logger.Errorln(err)
			return err
		}
	}

	return nil
}

func CreateCsvFile(filePath string) error {
	f1, err := os.Create(filePath)
	if err != nil {
		Logger.Errorln(err)
		return err
	}
	defer f1.Close()
	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func DirIsEmpty(path string) (bool, error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		Logger.Errorln(err)
		return true, err
	}
	if len(dir) == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func GetChainCurHeight(chain Chain) (uint64, error) {
	c, err := ethclient.Dial(chain.Endpoint)
	defer c.Close()
	if err != nil {
		Logger.Errorln(err)
		panic(err)
	}

	height, err := c.BlockNumber(context.Background())
	for err != nil {
		Logger.Errorf("chain: %v,err: %v", chain.Name, err)
		time.Sleep(10 * time.Second)
		height, err = c.BlockNumber(context.Background())
	}
	return height, nil
}

func CreateMsgCsvFileName(chainName, fileDataStr string, endHeight uint64) string {
	name := chainName + "-" + "msg" + "-" + strconv.FormatUint(endHeight, 10) + "-" + fileDataStr
	return name
}

func CreateReceiverNftCsvFileName(chainName, tokenName, fileDataStr string, endHeight uint64) string {
	name := chainName + "-" + tokenName + "-" + "stat_receiver" + "-" + strconv.FormatUint(endHeight, 10) + "-" + fileDataStr
	return name
}

func CreateClaimNftSenderCsvName(key, fileDataStr string, srcHeight, dstHeight uint64) string {
	name := key + "-" + "stat_claim_nft_sender" + "-" + strconv.FormatUint(srcHeight, 10) + "-" + strconv.FormatUint(dstHeight, 10) + "-" + fileDataStr
	return name
}

func CreateNftCsvFileName(chainName, fileDataStr, tokenName string, endHeight uint64) string {
	name := chainName + "-" + tokenName + "-" + "stat_minter" + "-" + strconv.FormatUint(endHeight, 10) + "-" + fileDataStr
	return name
}

func CreateNftBridgeCsvFileName(chainName, fileDataStr, tokenName string, endHeight uint64) string {
	name := chainName + "-" + tokenName + "-" + "stat_sender" + "-" + strconv.FormatUint(endHeight, 10) + "-" + fileDataStr
	return name
}

func CreateNftBridgeForReceiverCsvFileName(chainName, fileDataStr, receiverChainName string, endHeight uint64, tokenName string) string {
	name := chainName + "-" + receiverChainName + "-" + tokenName + "-" + "stat_sender" + "-" + strconv.FormatUint(endHeight, 10) + "-" + fileDataStr
	return name
}

func CreateNftBridgeForReceiversCsvFileName(chainName, fileDataStr, receiverChainAndTokenName string, endHeight uint64) string {

	name := chainName + "-" + receiverChainAndTokenName + "-" + "stat_sender" + "-" + strconv.FormatUint(endHeight, 10) + "-" + fileDataStr
	return name
}

func CreateDomainCsvFileName(domainTy, fileDataStr string, endTimestamp int64) string {
	name := domainTy + "-" + "stat_domain" + "-" + strconv.FormatInt(endTimestamp, 10) + "-" + fileDataStr
	return name
}

func CreateCsvPath(csvFileName, csvRootDirPath, dirDataStr, statDirStr string) string {
	filePath := csvRootDirPath + "/" + dirDataStr + "/" + statDirStr + "/" + csvFileName
	return filePath
}

func GetYesterdayCsvData(filePath string) (*sync.Map, int, error) {
	list := &sync.Map{}
	l := 0
	fs, err := os.Open(filePath + ".csv")
	if err != nil {
		Logger.Errorln(err)
		return list, 0, err
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			Logger.Errorln(err)
			return list, 0, err
		}
		if err == io.EOF {
			break
		}
		list.Store(row[0], Member)
		l++
	}
	return list, l, nil
}

func ToCsv(list *sync.Map, filePath, name string) (int, error) {
	l := 0
	f, err := os.Create(filePath + "/" + name + ".csv")
	if err != nil {
		Logger.Errorln(err)
		return 0, err
	}
	defer f.Close()

	_, err = f.WriteString("\xEF\xBB\xBF")
	if err != nil {
		Logger.Errorln(err)
		return 0, err
	}

	w := csv.NewWriter(f) //创建一个新的写入文件流

	var data = make([][]string, 0)

	list.Range(func(key, value interface{}) bool {
		k := key.(string)
		data = append(data, []string{k})
		l++
		return true
	})

	err = w.WriteAll(data)
	if err != nil {
		Logger.Errorln(err)
		return 0, err
	}
	w.Flush()
	Logger.Infof("%v to csv over", name)
	return l, nil
}

func ErrorIsNotExist(e error) bool {
	if strings.Contains(e.Error(), "no such file or directory") {
		return true
	} else {
		return false
	}
}

func GetPrevChainEndHeight(csvRootDirPath, dirDataStr string) (map[string]uint64, error) {
	m := make(map[string]uint64)
	filePath := csvRootDirPath + "/" + dirDataStr + "/stat_msg_sender"
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	for _, file := range files {
		mates := strings.Split(file.Name(), "-")
		height, err := strconv.ParseUint(mates[2], 10, 64)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
		m[mates[0]] = height
	}
	return m, nil
}

func GetPrevDomainEndTimestamp(csvRootDirPath, dirDataStr string) (map[string]int64, error) {
	m := make(map[string]int64)
	filePath := csvRootDirPath + "/" + dirDataStr + "/stat_domain"
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	for _, file := range files {
		mates := strings.Split(file.Name(), "-")
		height, err := strconv.ParseInt(mates[2], 10, 64)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
		m[mates[0]] = height
	}
	return m, nil
}

func GetPrevChainTokensEndHeight(csvRootDirPath, dirDataStr string) (map[string]uint64, error) {
	m := make(map[string]uint64)
	filePath := csvRootDirPath + "/" + dirDataStr + "/stat_nftbridge_sender"
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	for _, file := range files {
		mates := strings.Split(file.Name(), "-")
		height, err := strconv.ParseUint(mates[3], 10, 64)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
		m[mates[0]+"-"+mates[1]] = height
	}
	return m, nil
}

type DoubleChainHeight struct {
	SrcHeight uint64 `json:"src_height"`
	DstHeight uint64 `json:"dst_height"`
}

func GetPrevChainStatClaimNftSenderEndHeight(csvRootDirPath, dirDataStr string) (map[string]DoubleChainHeight, error) {
	m := make(map[string]DoubleChainHeight)
	filePath := csvRootDirPath + "/" + dirDataStr + "/stat_claim_nft_sender"
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	for _, file := range files {
		mates := strings.Split(file.Name(), "-")
		srcHeight, err := strconv.ParseUint(mates[5], 10, 64)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
		dstHeight, err := strconv.ParseUint(mates[6], 10, 64)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}

		key := mates[0] + "-" + mates[1] + "-" + mates[2] + "-" + mates[3]

		m[key] = DoubleChainHeight{
			SrcHeight: srcHeight,
			DstHeight: dstHeight,
		}
	}
	return m, nil
}

func GetPrevChainReceiverNftEndHeight(csvRootDirPath, dirDataStr string) (map[string]uint64, error) {
	m := make(map[string]uint64)
	filePath := csvRootDirPath + "/" + dirDataStr + "/stat_nftbridge_receiver_nft"
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	for _, file := range files {
		mates := strings.Split(file.Name(), "-")
		height, err := strconv.ParseUint(mates[3], 10, 64)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
		m[mates[0]+"-"+mates[1]] = height
	}
	return m, nil
}

func GetPrevStatNftEndHeight(csvRootDirPath, dirDataStr string) (map[string]uint64, error) {
	m := make(map[string]uint64)
	filePath := csvRootDirPath + "/" + dirDataStr + "/stat_nft_minter"
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	for _, file := range files {
		mates := strings.Split(file.Name(), "-")
		height, err := strconv.ParseUint(mates[3], 10, 64)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
		m[mates[0]+"-"+mates[1]] = height
	}
	return m, nil
}

func GerPrevChainReceiverEndHeight(csvRootDirPath, dirDataStr string) (map[string]uint64, error) {
	m := make(map[string]uint64)
	filePath := csvRootDirPath + "/" + dirDataStr + "/stat_nftbridge_receiver_chain_sender"
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		Logger.Errorln(err)
		return nil, err
	}

	for _, file := range files {
		Logger.Infof("file: %v", file.Name())
		mates := strings.Split(file.Name(), "-")
		height, err := strconv.ParseUint(mates[4], 10, 64)
		if err != nil {
			Logger.Errorln(err)
			return nil, err
		}
		m[mates[0]+"-"+mates[1]+"-"+mates[2]] = height
	}
	return m, nil
}

func ToInfoJson(filePath string, infos []*Info) error {
	f, err := os.Create(filePath + "/" + "info" + ".json")
	if err != nil {
		Logger.Errorln(err)
		return err
	}
	defer f.Close()

	j, err := json.Marshal(infos)
	if err != nil {
		Logger.Errorln(err)
		return err
	}
	_, err = f.Write(j)
	if err != nil {
		Logger.Errorln(err)
		return err
	}
	return nil
}

func JudgeWhetherItIsNextMonth(csvRootDirPath, dirDataStr string) (bool, error) {
	//m := make(map[string]uint64)
	filePath := csvRootDirPath + "/" + dirDataStr + "/stat_msg_sender"
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		Logger.Errorln(err)
		return false, err
	}
	prevMonth := ""
	nowMonth := time.Now().Month().String()

	for _, file := range files {
		mates := strings.Split(file.Name(), "-")
		prevDate := strings.Split(mates[3], ".")
		prevMonth = prevDate[1]
		break
	}
	if prevMonth == nowMonth {
		Logger.Infof("diff month, prevMonth: %v, nowMonth: %v", prevMonth, nowMonth)
		return false, nil
	} else {
		Logger.Infof("same month, prevMonth: %v, nowMonth: %v", prevMonth, nowMonth)
		return true, nil
	}
}
