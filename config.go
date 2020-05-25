package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	// MysqlConfig MysqlConfig
	dbConfig = &struct {
		IP       string `json:"ip,omitempty"`
		Port     int    `json:"port,omitempty"`
		Dbname   string `json:"dbnameo,omitempty"`
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}{}
)

const (
	defaultConfigName = "db.yml"
	defaultDir        = "../"
)

func init() {
	var defaultConfigPath string

	_, err := os.Stat(defaultConfigName)
	if err != nil && !os.IsExist(err) {
		defaultConfigPath = filepath.Join(defaultDir, defaultConfigName)
	} else {
		defaultConfigPath = defaultConfigName
	}
	f, err := os.Open(defaultConfigPath)
	if err != nil {
		log.Panic(err)
	}
	viper.SetConfigType("yml")
	err = viper.ReadConfig(f)
	if err != nil {
		log.Panic(err)
	}
	viper.UnmarshalKey("mysql", dbConfig)
	log.Println(dbConfig.IP)
}
