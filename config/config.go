package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

//func init() {
//	// 读取
//	f, err := os.Open(ConfigPath)
//	defer f.Close()
//	if err != nil {
//		log.Errorf("Init config Error %v", err)
//		return
//	}
//	// 解构
//	err = yaml.NewDecoder(f).Decode(&Conf)
//	if err != nil {
//		log.Errorf("Decode Conf Error %v", err)
//		panic(err)
//	}
//}

// BuildConfig 根据configPath打开config.yaml
func BuildConfig(configPath string) {
	// 读取
	f, err := os.Open(configPath)
	defer f.Close()
	if err != nil {
		log.Errorf("Init config Error %v", err)
		return
	}
	// 解构
	err = yaml.NewDecoder(f).Decode(&Conf)
	if err != nil {
		log.Errorf("Decode Conf Error %v", err)
		panic(err)
	}
}

var (
	Conf       = Config{}
	ConfigPath = "./config/config.yaml"
)

type Storage struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Base     string `yaml:"base"`
}
type SuperAdmin struct {
	Name     string `yaml:"name"`
	PassWord string `yaml:"password"`
	Email    string `yaml:"email"`
}
type LogConfig struct {
	DirPath    string `yaml:"dir-path"`
	NameFormat string `yaml:"name-format"`
}
type EmailConfig struct {
	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	AuthorizationCode string `yaml:"authorization-code"`
	ToolEmail         string `yaml:"tool-email"`
}
type Config struct {
	Storages map[string]Storage `yaml:"storages"`
	Admin    SuperAdmin         `yaml:"admin"`
	Secret   string             `yaml:"secret"`
	Log      LogConfig          `yaml:"log"`
	Mail     EmailConfig        `yaml:"mail"`
}
