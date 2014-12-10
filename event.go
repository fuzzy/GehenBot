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
	"strings"
	// 3rd party
	"github.com/husio/go-irc"
)

type IrcEvent struct {
	ircCommand string
	fromNick   string
	fromHost   string
	command    string
	params     []string
	args       []string
}

// ParseEvent(message irc.Message)

func ParseEvent(message irc.Message) IrcEvent {
	var cmnd, fromNick, fromHost string
	var args []string
	var evt IrcEvent

	Debug(fmt.Sprintf("Entered into ParseEvent(): %s\n", message))

	// find out if we start with our command character,
	if len(message.Trailing()) > 0 {
		if message.Trailing()[0] == cfg.CommandChar[0] {
			cmnd = strings.Split(message.Trailing(), " ")[0][1:]
			args = strings.Split(message.Trailing(), " ")[1:]
			fromNick = strings.Split(message.Prefix(), "!")[0]
			fromHost = strings.Split(message.Prefix(), "!")[1]

			evt.ircCommand = message.Command()
			evt.fromNick = fromNick
			evt.fromHost = fromHost
			evt.command = cmnd
			evt.params = message.Params()
			evt.args = args
		}
	}
	return evt
}
