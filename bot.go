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
)

/////////////////////
// BotInstance Object
/////////////////////

type Handler struct {
	event    string
	callback string
}

type BotInstance struct {
	address  string // irc server address
	port     int
	channels []string // list of channels to join
	scripts  []string
	nick     string // bot nickname
	name     string // This is the client name (generally nick)
	err      error  // error type, just use it everywhere.
	lua      *lua.State
	handlers []Handler
	conn     IrcClient
	cmdchar  string // command character
}

// BotInstance.Connect()

func (b *BotInstance) Connect() {
	b.conn.Host = b.address
	b.conn.Port = b.port
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
		// TODO:
		// This neds to take into account that we now have plugins
		// that can be presented in one of any number of languages
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

/*******************
** Public methods **
*******************/

func (b *BotInstance) Say(target string, message string) {
	b.conn.SendRaw(fmt.Sprintf("PRIVMSG %s :%s", target, message))
}

func (b *BotInstance) Action(target string, message string) {
	b.conn.SendRaw(fmt.Sprintf("PRIVMSG %s :\001ACTION %s\001", target, message))
}

func (b *BotInstance) Join(target string) {
	b.conn.SendRaw(fmt.Sprintf("JOIN %s", target))
}

func (b *BotInstance) Part(target string) {
	b.conn.SendRaw(fmt.Sprintf("PART %s", target))
}

func (b *BotInstance) Quit() {
	b.conn.SendRaw("QUIT")
}

func (b *BotInstance) Nick(nick string) {
	b.conn.SendRaw(fmt.Sprintf("NICK %s", nick))
}

func (b *BotInstance) Whois(nick string) {
	b.conn.SendRaw(fmt.Sprintf("WHOIS %s", nick))
}

func (b *BotInstance) Mode(target string, mode string) {
	b.conn.SendRaw(fmt.Sprintf("MODE %s %s", target, mode))
}
