/*
 * @Author: wtf
 * @Date: 2020-08-25 14:05:49
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-25 14:35:41
 * @Description: plase write Description
 */
package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)


var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func GetLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func GetLogFileFullPath() string {
	prefixPath := GetLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch  {
		case os.IsNotExist(err):
			mkDir()
		case os.IsPermission(err):
			log.Fatalf("Permission: %v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to openfile :%v", err)
	}
	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	path := fmt.Sprintf("%s/%s", dir, GetLogFilePath())
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}


