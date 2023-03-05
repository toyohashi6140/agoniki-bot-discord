package usecase

import "github.com/bwmarrin/discordgo"

type Responser interface {
	Response(*discordgo.Session, *discordgo.MessageCreate) error
}
