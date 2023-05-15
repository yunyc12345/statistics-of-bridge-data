package task

import (
	"encoding/json"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"os"
	"testing"
	"time"
	"yunyc12345/statistics-of-bridge-data/utils"
)

func Test1(t *testing.T) {
	configPath := "/home/m/go/src/yunyc12345/statistics-of-bridge-data/config.json"
	jsonFile, err := os.Open(configPath)

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

	crontab := cron.New(cron.WithSeconds())

	spec := "50 0 * * * "
	_, err = crontab.AddFunc(spec, func() {
		MsgSenderTask(config, time.Now(), nil)
	})
	if err != nil {
		panic(err)
	}

	_, err = crontab.AddFunc(spec, func() {
		NftMinterTask(config, time.Now(), nil)
	})
	if err != nil {
		panic(err)
	}

	_, err = crontab.AddFunc(spec, func() {
		NftBridgeSenderTask(config, time.Now(), nil)
	})
	if err != nil {
		panic(err)
	}

	crontab.Start()

	select {}
}
