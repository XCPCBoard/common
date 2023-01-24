package main

import (
	"errors"
	"fmt"
	"github.com/XCPCBoard/common/config"
	"github.com/XCPCBoard/common/logger"
	_ "github.com/XCPCBoard/common/logger"
)

func main() {
	//config.BuildConfig("./config/config.yaml")
	//fmt.Println(config.Conf.Admin.Name)
	//fmt.Println(config.Conf.Secret)
	logger.Logger.Info("hh", 0, "xx")
	logger.Logger.Debug("hh", 0, "xx")
	logger.Logger.Warn("hh", 0, "xx")
	logger.Logger.Error("1", errors.New("12"), 0, fmt.Sprintf("%s : %s", "aa", "aa"))
	//logger.Logger.Panic("1", 0, "vv")
	logger.Logger.Fatal("hh", 0, "xx")

}

func init() {
	config.BuildConfig("./config/config.yaml")
	err := logger.InitLogger()
	if err != nil {
		panic(err)
	}

}
