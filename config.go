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
)

type Config struct {
	Daemonize   bool   `json:"daemonize"`
	Verbose     bool   `json:"verbose"`
	Debug       bool   `json:"debug"`
	CommandChar string `json:"command_char"`
	PluginDir   string `json:"plugin_dir"`
	Bots        []struct {
		Channels []string `json:"channels"`
		Nick     string   `json:"nick"`
		Scripts  []string `json:"scripts"`
		Server   string   `json:"server"`
		User     string   `json:"user"`
	} `json:"bots"`
}

func ReadConfig(jsonCfg string) Config {
	var cfg Config

	cfgData, err := ioutil.ReadFile(jsonCfg)
	if err != nil {
		Fatal(fmt.Sprintf("Error in io.ReadFile(): %s\n", err))
	}

	err = json.Unmarshal([]byte(cfgData), &cfg)
	if err != nil {
		Fatal(fmt.Sprintf("Error in ReadConfig(): %s\n", err))
	}

	return cfg
}
