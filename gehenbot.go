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
	"fmt"
	"log"
	"os"
)

var (
	cfg          Config
	GehenVersion = "0.2.0"
	GehenAuthor  = "Mike 'Fuzzy' Partin <fuzzy@fumanchu.org>"
)

func Log(line string) {
	if cfg.Verbose {
		log.Println(fmt.Sprintf("INFO: %s", line))
		fmt.Println(fmt.Sprintf("\033[1;32mINFO\033[0m: %s", line))
	}
}

func Debug(line string) {
	if cfg.Debug {
		log.Println(fmt.Sprintf("DEBUG: %s", line))
	}
}

func Fatal(line string) {
	log.Fatalln(fmt.Sprintf("FATAL: %s", line))
	os.Exit(1)
}

// main program entry

func main() {
	var bots []BotInstance // this shouldn't really need a note

	// get our config data
	cfg = ReadConfig("./gehenbot.json")

	// if we have been directed to daemonize, then we need to do so
	if cfg.Daemonize {
		Log("We should be daemonizing")
	}

	mlog_fn := fmt.Sprintf("%s/gehenbot-debug.log", cfg.LogDir)
	mlog, err := os.OpenFile(mlog_fn, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Fatal(fmt.Sprintf("Error opening %s/gehenbot-debug.log", cfg.LogDir))
	}
	defer mlog.Close()
	log.SetOutput(mlog)

	// set our config data
	idx := 1
	for _, b := range cfg.Bots {
		bot := BotInstance{}
		bot.address = b.Server
		bot.channels = b.Channels
		bot.scripts = b.Scripts
		bot.nick = b.Nick
		if b.User == "" {
			bot.name = b.Nick
		} else {
			bot.name = b.User
		}
		// fire off the connection and event handlers
		bots = append(bots, bot)
		if len(cfg.Bots) == idx {
			bot.Connect()
		} else {
			go bot.Connect()
			idx += 1
		}
	}

}
