package utils

import "os"

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
