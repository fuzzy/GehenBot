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
	// 3rd party
	"github.com/aarzilli/golua/lua"
	"github.com/thoj/go-ircevent"
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
	nick     string   // bot nickname
	name     string   // This is the client name (generally nick)
	err      error    // error type, just use it everywhere.
	lua      *lua.State
	handlers []Handler
}

// BotInstance.Connect()

func (b BotInstance) Connect() {
	conn := irc.IRC(b.nick, b.name)

	/* / verbosity handling
	if cfg.Verbose {
		conn.VerboseCallbackHandler = true
	}
	if cfg.Debug {
		conn.Debug = true
	}
	*/

	/******************************
	** Embedded language support **
	******************************/

	// Initialize our lua interpreter
	b.lua = lua.NewState()
	defer b.lua.Close()
	b.lua.OpenLibs()
	// Register our exposed lua functions
	b.lua.Register("register_handler", b.registerHandler)

	Debug(fmt.Sprintf("## %d", len(b.handlers)))
	b.lua.DoFile("/home/mike/Devel/go/gehenbot/plugins/greet.lua")
	Debug(fmt.Sprintf("## %d", len(b.handlers)))

	// connect to the server
	err := conn.Connect(b.address)
	if err != nil {
		Fatal(err.Error())
	}

	// and setup the join handler
	conn.AddCallback("001", func(e *irc.Event) {
		for _, channel := range b.channels {
			conn.Join(channel)
		}
	})

	// set our generic event handler to be the callback for all the event
	// types we give a rats ass about.
	conn.AddCallback("PRIVMSG", b.EventHandler)
	conn.AddCallback("JOIN", b.EventHandler)
	conn.AddCallback("PART", b.EventHandler)
	conn.AddCallback("KICK", b.EventHandler)
	conn.AddCallback("MODE", b.EventHandler)
	conn.AddCallback("QUIT", b.EventHandler)

	conn.Loop()
}

func (b BotInstance) Connected() bool { return true }
