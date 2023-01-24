package logger

import (
	"context"
	"fmt"
	"github.com/XCPCBoard/common/config"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"time"
)

//***************************Log 结构体****************************//

// Log 封装以便日后更改，升级
// use slog
type Log struct {
	entity *slog.Logger
}

// Error logs at LevelError. If err is non-nil, Error appends Any(ErrorKey, err) to the list of attributes.
func (l *Log) Error(msg string, err error, args ...any) {
	l.entity.Error(msg, err, args)
}

func (l *Log) Debug(msg string, args ...any) {
	l.entity.Debug(msg, args)
}

func (l *Log) Info(msg string, args ...any) {
	l.entity.Info(msg, args)
}
func (l *Log) Warn(msg string, args ...any) {
	l.entity.Warn(msg, args)
}

// Context returns l's context, which may be nil.
func (l *Log) Context() context.Context {
	return l.entity.Context()
}

// Log 自定义级别: 目前超过8都是error，低于-4都是debug
//
//	LevelDebug Level = -4
//	LevelInfo  Level = 0
//	LevelWarn  Level = 4
//	LevelError Level = 8
func (l *Log) Log(level int, msg string, args ...any) {
	le := slog.Level(level)
	l.entity.Log(le, msg, args)
}

//***************************Log 初始化****************************//

var Logger *Log

func InitLogger() error {

	//init file
	dirPath := config.Conf.Log.DirPath
	if !fileExists(dirPath) {
		if err := os.Mkdir(dirPath, 0777); err != nil {
			return err
		}
	}
	path := dirPath + "/" + config.Conf.Log.NameFormat
	if !fileExists(path) {
		file, err := os.Create(path)
		defer file.Close()
		if err != nil {
			fmt.Println("mkdir log err!")
			return err
		}
	}

	witter, err := getWriter(path)
	if err != nil {
		return err
	}

	//Text 类型
	handle := slog.HandlerOptions{
		AddSource: true,
	}.NewTextHandler(witter)
	
	//赋值
	Logger.entity = slog.New(handle)
	return nil
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
