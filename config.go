package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	CommandChar string `json:"command_char"`
	CommandDir  string `json:"command_dir"`
	PluginDir   string `json:"plugin_dir"`
	Bots        []struct {
		Channels []string `json:"channels"`
		Nick     string   `json:"nick"`
		Scripts  []string `json:"scripts"`
		Server   string   `json:"server"`
		User     string   `json:"user"`
	} `json:"bots"`
}

func ReadConfig(jsonCfg string) {
	cfgData, err := ioutil.ReadFile(jsonCfg)
	if err != nil {
		log.Fatalf("Error in io.ReadFile(): %s\n", err)
	}

	err = json.Unmarshal([]byte(cfgData), &cfg)
	if err != nil {
		log.Fatalf("Error in ReadConfig(): %s\n", err)
	}
}
