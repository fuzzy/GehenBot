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
	"github.com/husio/go-irc"
)

/////////////////////
// BotInstance Object
/////////////////////

type CallBack func(event IrcEvent) int

type Plugin struct {
	PluginType string
	Callback   CallBack
}

type BotInstance struct {
	address  string        // irc server address
	channels []string      // list of channels to join
	nick     string        // bot nickname
	name     string        // This is the client name (generally nick)
	conn     irc.Client    // base connection type,
	err      error         // error type, just use it everywhere.
	evChan   chan IrcEvent // event channel.
	plugins  []Plugin
}

// BotInstance.Connect()

func (b BotInstance) Connect() {
	// If we have not been given an address error
	if b.address != "" {
		// if we don't have a nick, set the default
		if b.nick == "" {
			b.nick = "GehenBot"
		}
		// likewise with our name
		if b.name == "" {
			b.name = b.nick
		}

		// at this point, if we don't have a channel it isn't a showstopper
		// but it doesn't make much sense either.
		b.conn, b.err = irc.Connect(b.address)
		Log(b.address)
		if b.err != nil {
			Fatal(fmt.Sprintf("Error in BotInstance.Connect()->Connect(): %s\n", b.err))
		}
	} else {
		Fatal("Bot instance was misconfigured when Connect() was called.")
	}

	// well we made it
	b.Nick(b.nick)
	b.User(b.name)

	// now join our channels
	for _, ch := range b.channels {
		b.Join(ch)
	}

	// setup our lua integration
	// L := b.InitLua()
	// defer L.Close()

	for {
		message, err := b.conn.ReadMessage()
		event := ParseEvent(message)

		if err != nil {
			Fatal(fmt.Sprintf("Error in BotInstance.Connect()->ReadMessage(): %s\n", err))
		}

		if message.Command() == "PING" {
			b.Pong(message.Trailing())
		}

		if message.Command() == "PRIVMSG" {
			Log(fmt.Sprintf("%s %s!%s %s %s %s\n", event.ircCommand, event.fromNick, event.fromHost, event.params, event.command, event.args))
		}
	}

}

func (b BotInstance) Connected() bool { return true }
