package main

import (
	"fmt"
	"irohbot/handlers"
	"irohbot/utils"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var dg *discordgo.Session

func init() {
	utils.ReadProperties()
}

func init() {
	var err error
	dg, err = discordgo.New("Bot " + utils.GetProperty("bot.token"))

	if err != nil {
		log.Fatal("Error creating Discord session,", err)
	}
}

func main() {
	var err error

	dg.AddHandler(handlers.MessageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
