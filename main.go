package main

import (
	"fmt"
	"github.com/XCPCBoard/common/config"
)

func main() {
	config.BuildConfig("./config/config.yaml")
	fmt.Println(config.Conf.Admin.Name)
	fmt.Println(config.Conf.Secret)

}
