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
	// Stdlib
	"fmt"
	"strconv"
	"strings"
	// 3rd party
	"github.com/aarzilli/golua/lua"
)

/////////////////////
// BotInstance Object
/////////////////////

type Handler struct {
	event    string
	callback string
}

type BotInstance struct {
	address  string   // irc server address
	channels []string // list of channels to join
	scripts  []string
	nick     string // bot nickname
	name     string // This is the client name (generally nick)
	err      error  // error type, just use it everywhere.
	lua      *lua.State
	handlers []Handler
	conn     IrcClient
}

// BotInstance.Connect()

func (b *BotInstance) Connect() {
	tmp := strings.Split(b.address, ":")
	b.conn.Host = tmp[0]
	b.conn.Port, _ = strconv.Atoi(tmp[1])
	b.conn.Nick = b.nick
	b.conn.Name = b.name

	/******************************
	** Embedded language support **
	******************************/

	/************************
	** Setup lua scripting **
	************************/

	// Initialize our lua interpreter
	b.lua = lua.NewState()
	defer b.lua.Close()
	b.lua.OpenLibs()

	// setup the integration
	b.initLua()

	for _, ch := range b.scripts {
		b.lua.DoFile(fmt.Sprintf("%s/%s", cfg.PluginDir, ch))
	}

	// and setup the join handler
	b.conn.AddEventHandler("001", func(e *IrcMessage) {
		for _, channel := range b.channels {
			b.conn.SendRaw(fmt.Sprintf("JOIN %s", channel))
		}
	})

	// set our generic event handler to be the callback for all the event
	// types we give a rats ass about.
	b.conn.AddEventHandler("PRIVMSG", b.EventHandler)
	b.conn.AddEventHandler("JOIN", b.EventHandler)
	b.conn.AddEventHandler("PART", b.EventHandler)
	b.conn.AddEventHandler("KICK", b.EventHandler)
	b.conn.AddEventHandler("MODE", b.EventHandler)
	b.conn.AddEventHandler("QUIT", b.EventHandler)

	b.conn.Connect()
}
