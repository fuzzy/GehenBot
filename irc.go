package main

import (
	// Stdlib
	"os"
	"log"
	"fmt"
	"os/exec"
	"strings"
	// 3rd party
	"github.com/husio/go-irc"
)

//////////////////
// IrcEvent Object
//////////////////

type IrcEvent struct {
	ircCommand string
	fromNick string
	fromHost string
	command string
	params []string
	args []string
}

// ParseEvent(message irc.Message)

func ParseEvent(message irc.Message) IrcEvent {
	var cmnd string

	if message.Trailing()[0] == cfg.CmdChar[0] {
		cmnd = strings.Split(message.Trailing(), " ")[0][1:]
	} else {
		cmnd = ""
	}

	evt := IrcEvent{
		ircCommand: message.Command(),
		fromNick: strings.Split(message.Prefix(), "!")[0],
		fromHost: strings.Split(message.Prefix(), "!")[1],
		command: cmnd,
		params: message.Params(),
		args: strings.Split(message.Trailing(), " ")[1:],
	}
	return evt
}


/////////////////////
// BotInstance Object
/////////////////////

type BotInstance struct {
	address   string        // irc server address
	channels  []string      // list of channels to join
	nick      string        // bot nickname
	name      string        // This is the client name (generally nick)
	conn      irc.Client    // base connection type, useless until
                          // BotInstance.Connect() is called.
	err       error         // error type, just use it everywhere.
	evChan    chan IrcEvent // event channel.
}

// BotInstance.Connect()

func (b BotInstance) Connect() {
	// If we have not been given an address error
	if b.address != "" {
		// if we don't have a nick, set the default
		if b.nick == "" { b.nick = "GehenBot"  }
		// likewise with our name
		if b.name == "" { b.name = b.nick      }

		// at this point, if we don't have a channel it isn't a showstopper
		// but it doesn't make much sense either.
		b.conn, b.err = irc.Connect(b.address)
		if b.err != nil { log.Fatalf("Error in BotInstance.Connect()->Connect(): %s\n", b.err)      }
	} else { log.Fatalln("Bot instance was misconfigured when Connect() was called.") }

	// well we made it
	b.Nick(b.nick)
	b.User(b.name)

	// now join our channels
	for _,ch := range b.channels {
		b.Join(ch)
	}

	for {
		message, err := b.conn.ReadMessage()
		if err != nil { log.Fatalf("Error in BotInstance.Connect()->ReadMessage(): %s\n", err) }

		if message.Command() == "PING" { b.Pong(message.Trailing()) }
		if message.Command() == "PRIVMSG" {
			log.Println(message)
			event := ParseEvent(message)
			if event.command == "help" {
				// only caring about one command for help
				if len(strings.Split(message.Trailing(), " ")) > 1 {
					b.ShowCommands(event, strings.Split(message.Trailing(), " ")[1])
				} else {
					b.ShowCommands(event, "")
				}
			} else {		
				b.Command(event)
			}			
		}
	}
		
}

func (b BotInstance) Connected() bool { return true }

func (b BotInstance) ShowCommands(event IrcEvent, cmd string) bool {
	dir, err := os.Open("/home/fuzzy/gehenbot/commands/")
	if err != nil { return false }
	defer dir.Close()

	fileInfo, err := dir.Readdir(-1)
	if err != nil { return false }

	if cmd == "" {
		var cmnds []string
		for _, fi := range fileInfo {
			if fi.Mode() & 0111 != 0 {
				if fi.Mode() & os.ModeSymlink != os.ModeSymlink {
					cmnds = append(cmnds, strings.Split(fi.Name(), ".")[0]) 
				}
			}
		}
		if event.params[0][0] != '#' { 
			b.Say(event.fromNick, fmt.Sprintf("Commands: %s", strings.Join(cmnds, ", "))) 
		} else {
			b.Say(event.params[0], fmt.Sprintf("Commands: %s", strings.Join(cmnds, ", ")))
		}
	} else {
		for _, fi := range fileInfo {
			if strings.Split(fi.Name(), ".")[0] == cmd {
				// and now we know the filename matches our command
				cmnd := fmt.Sprintf("/home/fuzzy/gehenbot/commands/%s", fi.Name()) 
				out, err := exec.Command(cmnd, "--usage").Output()
				if err != nil { return false }
				if event.params[0][0] != '#' { 
					b.Say(event.fromNick, fmt.Sprintf("%s", out)) 
				} else {
					b.Say(event.params[0], fmt.Sprintf("%s", out))
				}
			}
		}
	}

	return true
}

func (b BotInstance) Command(event IrcEvent) bool {
	dir, err := os.Open("/home/fuzzy/gehenbot/commands/")
	if err != nil { return false }
	defer dir.Close()

	fileInfo, err := dir.Readdir(-1)
	if err != nil { return false }

	for _, fi := range fileInfo {
		if fi.Mode() & 0111 != 0 {
			// We now have an executable filename
			if strings.Split(fi.Name(), ".")[0] == event.command {
				// and now we know the filename matches our command
				cmd := fmt.Sprintf("/home/fuzzy/gehenbot/commands/%s", fi.Name()) 
				out, err := exec.Command(cmd, event.fromNick, event.fromHost, strings.Join(event.args, " ")).Output()
				if err != nil { return false }
				if event.params[0][0] != '#' { 
					b.Say(event.fromNick, fmt.Sprintf("%s", out)) 
				} else {
					b.Say(event.params[0], fmt.Sprintf("%s", out))
				}
			}
		}
	}

	return true
}

// BotInstance.Nick(nick string)

func (b BotInstance) Nick(nick string) { b.conn.Send("NICK %s", nick)            }

// BotInstance.User(name string)

func (b BotInstance) User(name string) { b.conn.Send("USER %s * * :...", b.name) }

// BotInstance.Join(chanName string)

func (b BotInstance) Join(chanName string) { b.conn.Send("JOIN %s", chanName)    }

// BotInstance.Say(sendTo string, message string)

func (b BotInstance) Say(sndTo string, msg string) { b.conn.Send("PRIVMSG %s :%s\n", sndTo, msg) }

// BotInstance.Pong(msg string)

func (b BotInstance) Pong(msg string) { b.conn.Send("PONG %s", msg) }
