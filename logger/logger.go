package logger

import (
	"fmt"
	"github.com/XCPCBoard/common/config"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"time"
)

var Log *slog.Logger

func InitLogger() (*slog.Logger, error) {

	//init file
	dirPath := config.Conf.Log.DirPath
	if !fileExists(dirPath) {
		if err := os.Mkdir(dirPath, 0777); err != nil {
			return nil, err
		}
	}
	path := dirPath + "/" + config.Conf.Log.NameFormat
	if !fileExists(path) {
		file, err := os.Create(path)
		defer file.Close()
		if err != nil {
			fmt.Println("mkdir log err!")
			return nil, err
		}
	}

	witter, err := getWriter(path)
	if err != nil {
		return nil, err
	}

	//Text 类型
	handle := slog.HandlerOptions{
		AddSource: true,
	}.NewTextHandler(witter)
	log := slog.New(handle)

	return log, nil
}

func getWriter(path string) (io.Writer, error) {
	// 保存60天内的日志，每24小时(整点)分割一次日志
	return rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Hour*24*60),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

}

// fileExists 查看文件/文件夹是否存在
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
