package main

import (
	"os"
	"log"
	"encoding/json"
)

type Config struct {
	CmdChar  string
	Nick     string
	User     string
	Servers  map[string][]string
}

func ReadConfig(jsonCfg string) Config {
	cfgFile, _ := os.Open(jsonCfg)
	decoder    := json.NewDecoder(cfgFile)
	cfgData    := Config{}
	err        := decoder.Decode(&cfgData)
	
	if err != nil { log.Fatalf("Error in ReadConfig(): %s\n", err) }

	return cfgData
}
