package utils

import (
	"encoding/csv"
	"os"
	"sync"
)

func ToCsv(list *sync.Map, filePath, name string) error {

	f, err := os.Create(filePath + "/" + name + ".csv")
	if err != nil {
		Logger.Errorln(err)
		return err
	}
	defer f.Close()

	_, err = f.WriteString("\xEF\xBB\xBF")
	if err != nil {
		Logger.Errorln(err)
		return err
	}

	w := csv.NewWriter(f) //创建一个新的写入文件流

	var data = make([][]string, 0)

	list.Range(func(key, value interface{}) bool {
		k := key.(string)
		data = append(data, []string{k})
		return true
	})

	err = w.WriteAll(data)
	if err != nil {
		Logger.Errorln(err)
		return err
	}
	w.Flush()
	Logger.Infof("%v to csv over", name)
	return nil
}
