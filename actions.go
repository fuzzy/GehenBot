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

/////////////////
// IRC Actions //
/////////////////

// BotInstance.Nick(nick string)

func (b BotInstance) Nick(nick string) { b.conn.Send("NICK %s", nick) }

// BotInstance.User(name string)

func (b BotInstance) User(name string) { b.conn.Send("USER %s * * :...", b.name) }

// BotInstance.Join(chanName string)

func (b BotInstance) Join(chanName string) { b.conn.Send("JOIN %s", chanName) }

// BotInstance.Say(sendTo string, message string)

func (b BotInstance) Say(sndTo string, msg string) { b.conn.Send("PRIVMSG %s :%s\n", sndTo, msg) }

// BotInstance.Pong(msg string)

func (b BotInstance) Pong(msg string) { b.conn.Send("PONG %s", msg) }
