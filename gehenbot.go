package main

import (
	"fmt"
	"log"
	"time"
)

// globals that make me ill, move to configs asap
var (
	bots         []BotInstance
	cfg          Config
	GehenVersion = "0.2.0"
	daemon       = false
	verbose      = false
)

// main program entry

func main() {

	// if we have been directed to daemonize, then we need to do so
	if daemon {
		log.Println("We should be daemonizing")
	}

	// if we are supposed to be verbose, we need to pass that off to the BotInstance objects
	if verbose {
		log.Println("We should be logging")
	}

	// get our config data
	ReadConfig("./gehenbot.json")

	// print our little banner
	fmt.Printf("\nGehenBot v%s by Mike 'Fuzzy' Partin <fuzzy@fumanchu.org>\n\n", GehenVersion)

	// set our config data
	for _, b := range cfg.Bots {
		bot := BotInstance{}
		bot.address = b.Server
		bot.channels = b.Channels
		bot.nick = b.Nick
		if b.User == "" {
			bot.name = b.Nick
		} else {
			bot.name = b.User
		}
		// fire off the connection and event handlers
		go bot.Connect()
		// and push it onto the stack
		bots = append(bots, bot)
	}

	for {
		for _, b := range bots {
			if b.Connected() {
				time.Sleep(100000)
			}
		}
	}
}
