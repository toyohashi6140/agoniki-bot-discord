package usecase

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/toyohashi6140/discord-bot/constants"
	"github.com/toyohashi6140/discord-bot/utils"
)

type love struct{}

func NewLove() Responser {
	return &love{}
}

func (l *love) Response(session *discordgo.Session, event *discordgo.MessageCreate) error {
	for _, love := range constants.Loves {
		if strings.Contains(event.Content, love) {
			if err := utils.SendMessage(session, event.ChannelID, constants.CanNotFallInLoveUnlessYouAreBroken); err != nil {
				return err
			}
		}
	}
	return nil
}
