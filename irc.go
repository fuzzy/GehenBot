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
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type IrcMessage struct {
	Host    string
	Numeric string
	Target  string
	Args    []string
}

type IrcCallback struct {
	Event    string
	Callback func(*IrcMessage)
}

type IrcClient struct {
	Host      string
	Port      int
	Nick      string
	Name      string
	Input     chan string
	Socket    net.Conn
	Ssl       bool
	SslSock   *tls.Conn
	SslConfig *tls.Config
	SslVerify bool
	callbacks []IrcCallback
	error
}

func (i *IrcClient) Connect() {
	i.Input = make(chan string)

	if i.Host == "" {
		fmt.Println("ERROR: No value for Host attribute.")
		os.Exit(1)
	}
	if i.Port == 0 {
		i.Port = 6667
	} // this is a sane default

	if i.Nick == "" {
		i.Nick = "gicl"
	}
	if i.Name == "" {
		i.Name = i.Nick
	}

	i.Socket, i.error = net.Dial("tcp", fmt.Sprintf("%s:%d", i.Host, i.Port))
	if i.error != nil {
		fmt.Println(i.error.Error())
		Fatal(i.error.Error())
	}

	fmt.Println("test")
	if i.Ssl {
		fmt.Println("ssl shit started")
		if i.SslVerify {
			fmt.Printf("verifying the ssl shit")
			i.SslConfig = &tls.Config{ServerName: i.Host}
		} else {
			fmt.Println("not verifying the ssl shit")
			i.SslConfig = &tls.Config{ServerName: i.Host, InsecureSkipVerify: true}
		}
		i.SslSock = tls.Client(i.Socket, i.SslConfig)
		if i.error = i.SslSock.Handshake(); i.error != nil {
			fmt.Println(i.error.Error())
			Fatal(i.error.Error())
		}
	}

	if i.error != nil {
		fmt.Printf("ERROR: %s\n", i.error)
		os.Exit(1)
	} else {
		go i.pingLoop()
		go i.writeLoop()

		i.AddEventHandler("PING", func(e *IrcMessage) { i.SendRaw(fmt.Sprintf("PONG %d", time.Now().UnixNano())) })

		var buffer *bufio.Reader
		if !i.Ssl {
			buffer = bufio.NewReader(i.Socket)
		} else if i.Ssl {
			buffer = bufio.NewReader(i.SslSock)
		}

		for {

			if !i.Ssl {
				i.Socket.SetReadDeadline(time.Now().Add(58 + (15 * time.Minute)))
			} else if i.Ssl {
				i.SslSock.SetReadDeadline(time.Now().Add(58 + (15 * time.Minute)))
			}

			str, err := buffer.ReadString('\n')
			fmt.Println(str)
			var zero time.Time
			if !i.Ssl {
				i.Socket.SetReadDeadline(zero)
			} else if i.Ssl {
				i.SslSock.SetReadDeadline(zero)
			}

			if len(str) > 0 {
				Log(strings.TrimRight(str, "\n"))

				// parse the line into a IrcMessage object
				var event IrcMessage
				var bits []string
				var args_t string

				bits = strings.Split(strings.TrimRight(str, "\r\n"), " ")
				args_t = strings.Join(bits[3:], " ")
				event.Host = bits[0][1:]
				event.Numeric = bits[1]
				event.Target = bits[2]

				if len(args_t) > 0 {
					event.Args = strings.Split(args_t[1:], " ")
				} else {
					event.Args = strings.Split(" ", " ")
				}
				fmt.Println("before")
				fmt.Println(str)
				if strings.Contains(strings.TrimRight(str, "\n"), "*** Found your hostname") {
					fmt.Println("during")
					i.SendRaw(fmt.Sprintf("NICK %s", i.Nick))
					i.SendRaw(fmt.Sprintf("USER %s 0.0.0.0 0.0.0.0 :%s", i.Nick, i.Nick))
				}
				fmt.Println("after")
				// Now let the event handler do it's work
				i.eventHandler(event)

			}
			if err != nil {
				fmt.Println(err.Error())
				break
			}
		}
	}
}

func (i *IrcClient) SendRaw(msg string) {
	i.Input <- msg + "\r\n"
}

func (i *IrcClient) AddEventHandler(event string, callback func(*IrcMessage)) {
	i.callbacks = append(i.callbacks, IrcCallback{event, callback})
}

/******************************
** IrcClient private methods **
******************************/

func (i *IrcClient) eventHandler(event IrcMessage) {
	for _, handler := range i.callbacks {
		if event.Numeric == handler.Event {
			handler.Callback(&event)
		}
	}
}

// this is really only handy for keeping us connected (hopefully)
func (i *IrcClient) pingLoop() {
	for {
		i.SendRaw(fmt.Sprintf("PING %d", time.Now().UnixNano()))
		time.Sleep(1 * time.Minute)
	}
}

// this is what handles sending all the messages, I hope to move this
func (i *IrcClient) writeLoop() {
	for {
		b, ok := <-i.Input
		if !ok || b == "" || i.Socket == nil || i.SslSock == nil {
			return
		} else {
			if i.Socket != nil && !i.Ssl {
				i.Socket.SetWriteDeadline(time.Now().Add(1 * time.Minute))
				_, err := i.Socket.Write([]byte(b))
				if err != nil {
					fmt.Printf("ERROR: %s\n", err.Error())
					return
				}
				var zero time.Time
				i.Socket.SetWriteDeadline(zero)
			} else if i.SslSock != nil && i.Ssl {
				i.SslSock.SetWriteDeadline(time.Now().Add(1 * time.Minute))
				_, err := i.SslSock.Write([]byte(b))
				if err != nil {
					fmt.Printf("ERROR: %s\n", err.Error())
					return
				}
			}
		}
	}
}
