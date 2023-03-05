package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/toyohashi6140/discord-bot/config"
	"github.com/toyohashi6140/discord-bot/models"
	"github.com/toyohashi6140/discord-bot/usecase"
)

func main() {
	startDiscordBot()
}

func startDiscordBot() {
	discord, err := discordgo.New("Bot " + config.DiscordEnv.Token)
	if err != nil {
		panic(fmt.Sprintf("Unauthorized, %s", err.Error()))
	}
	if err := discord.Open(); err != nil {
		panic(fmt.Sprintf("Discord Session Failed, %s", err.Error()))
	}
	discord.AddHandler(chat)

	defer discord.Close()

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot
	fmt.Println("quit")
}

func chat(session *discordgo.Session, event *discordgo.MessageCreate) {
	// botの投稿は無視する
	if config.DiscordEnv.ClientID == event.Author.ID {
		return
	}
	agoniki := &models.Agoniki{}
	if err := usecase.NewCall(agoniki).Response(session, event); err != nil {
		fmt.Println("Message Send Failed ", err)
	}
	if err := usecase.NewLove().Response(session, event); err != nil {
		fmt.Println("Message Send Failed ", err)
	}
	if err := usecase.NewAngry(agoniki).Response(session, event); err != nil {
		fmt.Println("Message Send Failed ", err)
	}
	if err := usecase.NewGirlFriend(agoniki).Response(session, event); err != nil {
		fmt.Println("Message Send Failed ", err)
	}
}
