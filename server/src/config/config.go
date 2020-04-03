package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"sqlbar/server/src/logs"
)

type (
	// JFile struct
	JFile struct {
		JConfig `json:"config"`
	}

	// JConfig struct
	JConfig struct {
		JServer `json:"server"`
		JDb     `json:"db"`
	}

	// JServer struct
	JServer struct {
		Host     string `json:"host"`
		Name     string `json:"name"`
		Port     int    `json:"port"`
		LogLevel string `json:"log_level"`
	}

	// JDb struct
	JDb struct {
		Name     string `json:"name"`
		User     string `json:"user"`
		Password string `json:"password"`
	}
)

var (
	// ServerName var
	ServerName string
	// ServerVersion var
	ServerVersion string
	// ServerHost var
	ServerHost string
	// ServerPort var
	ServerPort int
	// ServerLogLevel var
	ServerLogLevel string
	// DbName var
	DbName string
	// DbUser var
	DbUser string
	// DbPassword var
	DbPassword string
)

func init() {
	logs.Log.PushFuncName("config", "config", "init")
	defer logs.Log.PopFuncName()

	// config -----------------------------------------------------------------
	file, err := os.Open("config.json")
	if err != nil {
		logs.Log.Error("os.Open('config.json')", err)
		return
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		logs.Log.Error("ioutil.ReadAll(file)", err)
		return
	}

	var jsonConfig JFile
	json.Unmarshal(byteValue, &jsonConfig)

	ServerName = jsonConfig.JConfig.JServer.Name
	ServerHost = jsonConfig.JConfig.JServer.Host
	ServerPort = jsonConfig.JConfig.JServer.Port
	ServerLogLevel = jsonConfig.JConfig.JServer.LogLevel
	DbName = jsonConfig.JConfig.JDb.Name
	DbUser = jsonConfig.JConfig.JDb.User
	DbPassword = jsonConfig.JConfig.JDb.Password

	// info -------------------------------------------------------------------
	file, err = os.Open("info.json")
	if err != nil {
		logs.Log.Error("os.Open('info.json')", err)
		return
	}
	defer file.Close()

	byteValueInfo, err := ioutil.ReadAll(file)
	if err != nil {
		logs.Log.Error("ioutil.ReadAll(file)", err)
		return
	}

	var jsonInfo InfoFile
	json.Unmarshal(byteValueInfo, &jsonInfo)

	ServerVersion = jsonInfo.InfoConfig.InfoServer.Version
	// ------------------------------------------------------------------------

	logs.Log.Info("CONFIG LOADED params:",
		"server.name:", ServerName,
		"server.version:", ServerVersion,
		"server.host:", ServerHost,
		"server.port:", ServerPort,
		"server.log_level:", ServerLogLevel,
		"db.name:", DbName,
		"db.user:", DbUser)
}
