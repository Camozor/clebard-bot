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
		handleDogImageCommand(s, m)
	}
}

func handleDogImageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	//response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer response.Body.Close()

	var clebardPhoto = "https://images.dog.ceo/breeds/dhole/n02115913_564.jpg"
	var messageEmbed = discordgo.MessageEmbed{Image: &discordgo.MessageEmbedImage{URL: clebardPhoto}}

	var message = discordgo.MessageSend{Embeds: []*discordgo.MessageEmbed{&messageEmbed}}
	_, err := s.ChannelMessageSendComplex(m.ChannelID, &message)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Could not send image to server for channel %s", m.ChannelID)
		return
	}

	//if response.StatusCode == 200 {
	//	var message = discordgo.MessageSend{Embeds: []*discordgo.MessageEmbed{&messageEmbed}}
	//	_, err := s.ChannelMessageSendComplex(m.ChannelID, &message)
	//	if err != nil {
	//		fmt.Println(err)
	//		fmt.Printf("Could not send image to server for channel %s", m.ChannelID)
	//		return
	//	}
	//}
}
