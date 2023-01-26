package main

import (
	"fmt"
	"github.com/XCPCBoard/common/config"
	"github.com/XCPCBoard/common/logger"
	_ "github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/common/mail"
)

func main() {

	fmt.Println("开发中")
}

func init() {
	config.BuildConfig("./config/config.yaml")
	err := logger.InitLogger()
	if err != nil {
		panic(err)
	}
	mail.InitEmail()
}
