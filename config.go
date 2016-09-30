/*
  Copyright (c) 2014, Mike 'Fuzzy' Partin <fuzzy@fumanchu.org>
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:

  1. Redistributions of source code must retain the above copyright notice, this
     list of conditions and the following disclaimer.

  2. Redistributions in binary form must reproduce the above copyright notice,
     this list of conditions and the following disclaimer in the documentation
     and/or other materials provided with the distribution.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
  FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
  DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
  SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
  CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
  OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
  OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main

import (
	// stdlib
	"encoding/json"
	"fmt"
	"io/ioutil"

	// 3rd party
	"github.com/go-yaml/yaml"
)

type Config struct {
	Daemonize bool   `json:"daemonize" yaml:"daemonize"`
	Verbose   bool   `json:"verbose" yaml:"verbose"`
	Debug     bool   `json:"debug" yaml:"debug"`
	LogDir    string `json:"log_dir" yaml:"log_dir"`
	PluginDir string `json:"plugin_dir" yaml:"plugin_dir"`
	Networks  []struct {
		Name        string   `json:"name" yaml:"name"`
		CommandChar string   `json:"command_char" yaml:"command_char"`
		Channels    []string `json:"channels" yaml:"channels"`
		Nick        string   `json:"nick" yaml:"nick"`
		Scripts     []string `json:"scripts" yaml:"scripts"`
		Server      string   `json:"server" yaml:"server"`
		Port        int      `json:"port" yaml:"port"`
		User        string   `json:"user" yaml:"user"`
		Ssl         bool     `json:"ssl" yaml:"ssl"`
		SslVerify   bool     `json:"ssl_verify" yaml:"ssl_verify"`
	} `json:"networks"`
}

func ReadConfig(jsonCfg string) Config {
	var cfg Config

	cfgData, err := ioutil.ReadFile(jsonCfg)
	if err != nil {
		Fatal(fmt.Sprintf("Error in io.ReadFile(): %s\n", err))
	}

	if jsonCfg[len(jsonCfg)-4:] == ".yml" || jsonCfg[len(jsonCfg)-4:] == "yaml" {
		err = yaml.Unmarshal(cfgData, &cfg)
	} else if jsonCfg[len(jsonCfg)-4:] == ".jsn" || jsonCfg[len(jsonCfg)-4:] == "json" {
		err = json.Unmarshal([]byte(cfgData), &cfg)
	}

	if err != nil {
		Fatal(fmt.Sprintf("Error in ReadConfig(): %s\n", err))
	}

	return cfg
}
