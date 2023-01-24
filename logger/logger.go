// Package logger
// @Author lyf
// @Update lyf 2023.01
package logger

import (
	"context"
	"errors"
	"fmt"
	"github.com/XCPCBoard/common/config"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"runtime"
	"time"
)

//***************************Log 结构体****************************//

// Log 封装以便日后更改，升级
// use slog
type Log struct {
	entity *slog.Logger
}

// getScour 获取出问题的位置
func (l *Log) getScour(skip int) string {
	if skip < 2 {
		l.entity.Error("获取出错代码位置skip错误", errors.New("获取出错代码位置skip错误"), "skip", skip)
		return "源代码位置获取失败，skip错误"
	}
	pc, codePath, codeLine, ok := runtime.Caller(skip)
	if !ok {
		// 不ok，函数栈用尽了
		l.entity.Error("获取出错代码位置函数栈用尽", errors.New("获取出错代码位置函数栈用尽"), "skip", skip)
		return "源代码位置获取失败，函数栈用尽"
	}

	// 拼接文件名与所在行
	code := fmt.Sprintf("%s:%d func name:%s", codePath, codeLine, runtime.FuncForPC(pc).Name())
	return fmt.Sprintf(code)

}

// Error
//
//	@description	logs at LevelError. If err is non-nil, Error appends Any(ErrorKey, err) to the list of attributes.
//	@param	deep	函数栈深度，若调用log的位置是错误发生的位置，则输入0; 否则输入封装的深度
//	@param	param	错误时的相关参数信息，建议用fmt.Sprintf()
func (l *Log) Error(msg string, err error, deep int, param string) {
	l.entity.Error(msg, err, "source", l.getScour(deep+2), "param", param)
}

// Debug
//
//	@description	debug级别
//	@param	deep	函数栈深度，若调用log的位置是错误发生的位置，则输入0; 否则输入封装的深度
//	@param	param	错误时的相关参数信息，建议用fmt.Sprintf()
func (l *Log) Debug(msg string, deep int, param string) {
	l.entity.Debug(msg, "source", l.getScour(deep+2), "param", param)
}

// Info
//
//	@description	info级别错误
//	@param	deep	函数栈深度，若调用log的位置是错误发生的位置，则输入0; 否则输入封装的深度
//	@param	param	错误时的相关参数信息，建议用fmt.Sprintf()
func (l *Log) Info(msg string, deep int, param string) {
	l.entity.Info(msg, "source", l.getScour(deep+2), "param", param)
}

// Warn
//
//	@description	Warn级别错误
//	@param	deep	函数栈深度，若调用log的位置是错误发生的位置，则输入0; 否则输入封装的深度
//	@param	param	错误时的相关参数信息，建议用fmt.Sprintf()
func (l *Log) Warn(msg string, deep int, param string) {
	l.entity.Warn(msg, "source", l.getScour(deep+2), "param", param)
}

// Fatal
//
//	@description	非必要不使用：致命错误，出现错误时程序无法正常运转，输出日志后程序退出(os.Exit(1))
//	@param	deep	函数栈深度，若调用log的位置是错误发生的位置，则输入0; 否则输入封装的深度
//	@param	param	错误时的相关参数信息，建议用fmt.Sprintf()
func (l *Log) Fatal(msg string, deep int, param string) {
	l.entity.Log(LevelFatal, msg, "source", l.getScour(deep+2), "param", param)
	os.Exit(1)
}

// Panic
//
//	@description	panic级别错误，输出日志后调用panic(msg)
//	@param	deep	函数栈深度，若调用log的位置是错误发生的位置，则输入0; 否则输入封装的深度
//	@param	param	错误时的相关参数信息，建议用fmt.Sprintf()
func (l *Log) Panic(msg string, deep int, param string) {
	l.entity.Log(LevelPanic, msg, "source", l.getScour(deep+2), "param", param)
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

		Level: LevelDebug, //需要输出的日志级别，默认为info，测试环境可以写成debug

		//AddSource: true, //为true时，输出调用log的位置信息，但是一旦做了封装就只会输出封装的位置，所以建议为false

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
