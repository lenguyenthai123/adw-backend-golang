package utils

import (
	"backend-golang/pkgs/log"
	"os"
	"path/filepath"
)

func GetFileName(filePath string) string {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.JsonLogger.Error(err.Error())
		panic(err)
	}

	return fileInfo.Name()
}

func GetDirectoryPath(filePath string) string {
	dirPath := filepath.Dir(filePath)
	return dirPath
}
