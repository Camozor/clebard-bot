package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var dogCommandList = []string{"woof", "wouf", "ouaf", "whouaf", "wouaf"}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("DISCORD_TOKEN")

	fmt.Println(token)

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dg.AddHandler(handleUserMessage)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err = dg.Close()
	if err != nil {
		return
	}
}

func handleUserMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if slices.Contains(dogCommandList, m.Content) {
		fmt.Println("Contient")
	} else {
		fmt.Println("Contient pas")
	}
}
