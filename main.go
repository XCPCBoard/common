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

	logger.Log.Error("1", errors.New("12"), "state", 500)
}

func init() {
	config.BuildConfig("./config/config.yaml")
	log, err := logger.InitLogger()
	if err != nil {
		panic(err)
	}
	logger.Log = log
}
