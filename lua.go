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
	// 3rd party
	"github.com/aarzilli/golua/lua"
)

// exported functions given to lua

func (b *BotInstance) registerHandler(L *lua.State) int {
	ev := L.ToString(1)
	cb := L.ToString(2)
	handler := Handler{event: ev, callback: cb}
	b.handlers = append(b.handlers, handler)
	return 1
}

// actions

func (b *BotInstance) luaSay(L *lua.State) int {
	tg := L.ToString(1)
	tx := L.ToString(2)
	b.conn.Privmsg(tg, tx)
	return 1
}

func (b *BotInstance) luaJoin(L *lua.State) int {
	ch := L.ToString(1)
	b.conn.Join(ch)
	return 1
}

func (b *BotInstance) luaPart(L *lua.State) int {
	ch := L.ToString(1)
	b.conn.Part(ch)
	return 1
}
