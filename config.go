/**
 * This is NOT a freeware, use is subject to license terms
 *
 * path   go-push/config
 * date   2018/9/7 10:14
 * author chenjingxiu
 */
package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"path/filepath"
)

type GlobalConfig struct {
	Database DB     `json:"database"`
	Redis    RD     `json:"redis"`
	RabbitMQ MQ     `json:"rabbitMQ"`
	Mongodb  Mog    `json:"mongodb"`
	Push     Secret `json:"push"`
	Api      Inter  `json:"api"`
}

type DB struct {
	Default DBItem `json:"default"`
}

type DBItem struct {
	Driver  string `json:"driver"`
	Dsn     string `json:"dsn"`
	Prefix  string `json:"prefix"`
	MaxOpen int    `json:"maxOpen"`
	MaxIdle int    `json:"maxIdle"`
	ShowSQl bool   `json:"showSQL"`
}

type RD struct {
	Default RDItem `json:"default"`
	Master  RDItem `json:"master"`
	Slave   RDItem `json:"slave"`
}

type RDItem struct {
	Url string `json:"url"`
}

type MQ struct {
	Default MQItem `json:"default"`
}

type MQItem struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	RoutingKey  string `json:"routingKey"`
	Dsn         string `json:"dsn"`
	VirtualHost string `json:"virtualHost"`
}

type Secret struct {
	HuaWei SecretItem `json:"HW"`
	Mi     SecretItem `json:"MI"`
}

type SecretItem struct {
	Name       string `json:"name"`
	AppId      string `json:"appId"`
	AppSecret  string `json:"appSecret"`
	Icon       string `json:"icon"`
	IntentHead string `json:"intentHead"`
	IntentTail string `json:"intentTail"`
	Package    string `json:"package"`
}

type Mog struct {
	Default MogItem `json:"default"`
}

type MogItem struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
	Dsn      string `json:"dsn"`
}

type Inter struct {
	Url map[string]string `json:"url"`
}

var (
	GConfig    *GlobalConfig
	configLock = new(sync.Mutex)
)

func init() {
	fileName, err := filepath.Abs("config.json")
	if err != nil {
		panic(err)
	}
	err = loadConfig(fileName)
	if err != nil {
		panic(err)
	}
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT)
	go func() {
		for {
			<-s
			loadConfig(fileName)
			fmt.Println("reload the config file")
		}
	}()
}

func GetConfig() *GlobalConfig {
	configLock.Lock()
	defer configLock.Unlock()
	return GConfig
}

func loadConfig(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read config file error: ", err)
		return err
	}
	config := new(GlobalConfig)
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Unmarshal the config file to struct error: ", err)
		return err
	}
	configLock.Lock()
	GConfig = config
	configLock.Unlock()
	return nil
}
