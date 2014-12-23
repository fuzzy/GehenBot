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

import "github.com/qur/gopy/lib"

func (b *BotInstance) initPython() {
	py.Initialize()

	methods := []py.Method{
		{"register", b.pyRegisterHandler, ""},
		{"say", b.pySay, ""},
		{"action", b.pyAction, ""},
		{"join", b.pyJoin, ""},
		{"part", b.pyPart, ""},
		{"quit", b.pyQuit, ""},
		{"nick", b.pyNick, ""},
		{"whois", b.pyWhois, ""},
		{"mode", b.pyMode, ""},
		{"myNick", b.pyMyNick, ""},
		{"myName", b.pyMyName, ""},
		{"scriptDir", b.pyScriptDir, ""},
		{"commandChar", b.pyCommandChar, ""},
	}

	_, err := py.InitModule("bot", methods)
	if err != nil {
		Fatal(err.Error())
	}
}

func (b *BotInstance) pyRegisterHandler(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}

func (b *BotInstance) pySay(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}

func (b *BotInstance) pyAction(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}

func (b *BotInstance) pyJoin(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}

func (b *BotInstance) pyPart(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}

func (b *BotInstance) pyQuit(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}

func (b *BotInstance) pyNick(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}

func (b *BotInstance) pyWhois(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}

func (b *BotInstance) pyMode(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}

func (b *BotInstance) pyMyName(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}

func (b *BotInstance) pyMyNick(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}

func (b *BotInstance) pyCommandChar(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}

func (b *BotInstance) pyScriptDir(args *py.Tuple) error {
	var target string
	var message string
	err := py.ParseTuple(args, "ss", &target, &message)
	if err != nil {
		Fatal(err.Error())
	}
	b.Say(target, message)
	return nil
}
