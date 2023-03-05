package usecase

import (
	"github.com/bwmarrin/discordgo"
	"github.com/toyohashi6140/discord-bot/constants"
	"github.com/toyohashi6140/discord-bot/models"
	"github.com/toyohashi6140/discord-bot/utils"
)

type call struct {
	agoniki *models.Agoniki
}

func NewCall(a *models.Agoniki) Responser {
	return &call{
		agoniki: a,
	}
}

func (c *call) Response(session *discordgo.Session, event *discordgo.MessageCreate) error {
	if isCall := c.agoniki.Call(event.Content); isCall {
		texts := c.getText()
		for _, text := range texts {
			if err := utils.SendReply(session, event.ChannelID, text, event.Reference()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *call) getText() []string {
	i := utils.RandNumber(100)
	if i%10 == 0 {
		return constants.CallMentionGroup[2]
	} else if i%2 == 0 {
		return constants.CallMentionGroup[0]
	}
	return constants.CallMentionGroup[1]
}
