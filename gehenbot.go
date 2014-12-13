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
	"time"
)

var (
	cfg          Config
	GehenVersion = "0.2.0"
	GehenAuthor  = "Mike 'Fuzzy' Partin <fuzzy@fumanchu.org>"
)

func Log(line string) {
	if cfg.Verbose {
		log.Println(fmt.Sprintf("INFO: %s", line))
	}
}

func Debug(line string) {
	if cfg.Debug {
		log.Println(fmt.Sprintf("DEBUG: %s", line))
	}
}

func Fatal(line string) { log.Fatalln(line) }

// main program entry

func main() {
	var bots []BotInstance // this shouldn't really need a note

	// get our config data
	cfg = ReadConfig("./gehenbot.json")

	// if we have been directed to daemonize, then we need to do so
	if cfg.Daemonize {
		Log("We should be daemonizing")
	}

	// set our config data
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
		go bot.Connect()
		// and push it onto the stack
		bots = append(bots, bot)
	}

	for {
		time.Sleep(100000)
	}
}
