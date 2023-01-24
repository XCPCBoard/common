package main

import (
	"errors"
	"github.com/XCPCBoard/common/config"
	"github.com/XCPCBoard/common/logger"
	_ "github.com/XCPCBoard/common/logger"
)

func main() {
	//config.BuildConfig("./config/config.yaml")
	//fmt.Println(config.Conf.Admin.Name)
	//fmt.Println(config.Conf.Secret)

	logger.Logger.Error("1", errors.New("12"), "state")
	logger.Logger.Error("1", errors.New("12"), "state")

}

func init() {
	config.BuildConfig("./config/config.yaml")
	err := logger.InitLogger()
	if err != nil {
		panic(err)
	}

}
