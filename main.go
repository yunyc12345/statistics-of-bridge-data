package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
	"yunyc12345/statistics-of-bridge-data/handle"
	"yunyc12345/statistics-of-bridge-data/utils"
)

func main() {

	utils.InitLogger(&utils.LogConf{
		Level: "info",
		Path:  "log/app.log",
	})
	cfp := os.Getenv("CONFIG_FILE_PATH")

	fmt.Printf("config path: %v\n", cfp)

	jsonFile, err := os.Open(cfp)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config utils.Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}

	err = utils.CreateCsvDir(config.Server.CsvFilePath)
	if err != nil {
		panic(err)
	}
	//j, _ := json.Marshal(config)
	//utils.Logger.Info(string(j))

	w := &sync.WaitGroup{}

	//go handle.TestnetNftbridgeSendersStatistics(config.TestnetNftbridgeSendersStatistics, w)

	go handle.MainnetNftMinterStatistics(config.MainnetNftMinterStatistics, w, config.Server.CsvFilePath)

	//go handle.MainnetNftbridgeSendersStatistics(config.MainnetNftbridgeSendersStatistics, w, config.Server.CsvFilePath)

	//go handle.MainnetMsgSendersStatistics(config.MainnetMsgSendersStatistics, w, config.Server.CsvFilePath)

	time.Sleep(time.Second * 20)
	w.Wait()
	utils.Logger.Info("server over")
}
