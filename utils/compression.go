package utils

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ZipStatIsExist(dirDataStr, csvRootDirPath string) (bool, error) {
	zipFilePath := csvRootDirPath + "/" + dirDataStr + ".zip"
	_, err := os.Stat(zipFilePath)
	if err == nil {
		Logger.Infof("zip file exist: %v", zipFilePath)
		return true, nil
	}
	if os.IsNotExist(err) {
		Logger.Infof("zip file not exist: %v", zipFilePath)
		return false, nil
	} else {
		Logger.Errorln(err)
	}

	return false, nil
}
func ZipStat(dirDataStr, csvRootDirPath string) error {
	sourceFilePath := csvRootDirPath + "/" + dirDataStr
	exists, err := PathExists(sourceFilePath)
	if err != nil {
		Logger.Errorln(err)
		return err
	}

	if !exists {
		Logger.Infof("source file not exists: %v", sourceFilePath)
		return errors.New("source file not exists")
	}

	b, err := ZipStatIsExist(dirDataStr, csvRootDirPath)
	if err != nil {
		Logger.Errorln(err)
		return err
	}

	if b {
		return nil
	}

	zipFilePath := csvRootDirPath + "/" + dirDataStr + ".zip"
	// 创建：zip文件
	zipfile, _ := os.Create(zipFilePath)
	defer zipfile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(sourceFilePath, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == sourceFilePath {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, sourceFilePath+`/`)

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})

	return nil
}
