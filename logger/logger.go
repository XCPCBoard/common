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
// args 需要由数个键值对构成，例如log.Error(msg,err,"name","小明")
func (l *Log) Error(msg string, err error, args ...any) {
	l.entity.Error(msg, err, args...)
}

// Debug debug级别
// args 需要由数个键值对构成，例如log.Debug(msg,"name","小明")
func (l *Log) Debug(msg string, args ...any) {
	l.entity.Debug(msg, args...)
}

// Info info级别
// args 需要由数个键值对构成，例如log.Info(msg,"name","小明")
func (l *Log) Info(msg string, args ...any) {
	l.entity.Info(msg, args...)
}

// Warn Warn级别
// args 需要由数个键值对构成，例如log.Warn(msg,"name","小明")
func (l *Log) Warn(msg string, args ...any) {
	l.entity.Warn(msg, args...)
}

// Fatal 非必要不使用：致命错误，出现错误时程序无法正常运转，输出日志后程序退出(os.Exit(1))
// args 需要由数个键值对构成
func (l *Log) Fatal(msg string, args ...any) {
	l.entity.Log(LevelFatal, msg, args...)
	os.Exit(1)
}

// Panic 输出日志后调用panic(msg)
// args 需要由数个键值对构成
func (l *Log) Panic(msg string, args ...any) {
	l.entity.Log(LevelPanic, msg, args...)
	panic(msg)
}

// Context returns l's context, which may be nil.
func (l *Log) Context() context.Context {
	return l.entity.Context()
}

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
	LevelPanic = slog.Level(12) // 输出日志后panic
	LevelFatal = slog.Level(16) // Fatal 致命错误，出现错误时程序无法正常运转，输出日志后程序退出
)

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
		Level:     LevelDebug, //需要输出的日志级别，默认为info，测试环境可以写成debug
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				// 处理自定义级别
				level := a.Value.Any().(slog.Level)

				//自定义错误等级划分
				switch {
				case level <= LevelDebug:
					a.Value = slog.StringValue("DEBUG")
				case level <= LevelInfo:
					a.Value = slog.StringValue("INFO")
				case level <= LevelWarn:
					a.Value = slog.StringValue("WARN")
				case level <= LevelError:
					a.Value = slog.StringValue("ERROR")
				case level <= LevelPanic:
					a.Value = slog.StringValue("PANIC")
				default:
					a.Value = slog.StringValue("FATAL")
				}
			}
			return a
		},
	}.NewTextHandler(witter)

	//赋值
	Logger = new(Log)
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
