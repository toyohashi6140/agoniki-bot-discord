package usecase

import (
	"github.com/bwmarrin/discordgo"
	"github.com/toyohashi6140/discord-bot/constants"
	"github.com/toyohashi6140/discord-bot/models"
	"github.com/toyohashi6140/discord-bot/utils"
)

type girlfriend struct {
	agoniki *models.Agoniki
}

func NewGirlFriend(a *models.Agoniki) Responser {
	return &girlfriend{
		agoniki: a,
	}
}

func (g *girlfriend) Response(session *discordgo.Session, event *discordgo.MessageCreate) error {
	emotion := g.agoniki.NegativeOrPositive(event.Content)
	var texts []string
	if emotion == models.Positive {
		texts = g.getPositiveText()
	} else if emotion == models.Negative {
		texts = g.getNegativeText()
	}
	for _, text := range texts {
		if err := utils.SendMessage(session, event.ChannelID, text); err != nil {
			return err
		}
	}
	return nil
}

func (g *girlfriend) getNegativeText() []string {
	i := utils.RandNumber(100)
	if i%3 == 0 {
		return constants.GirlFriendNegativeMentionGroup[0]
	} else if i%2 == 0 {
		return constants.GirlFriendNegativeMentionGroup[1]
	}
	return constants.GirlFriendNegativeMentionGroup[2]
}

func (g *girlfriend) getPositiveText() []string {
	i := utils.RandNumber(100)
	if i%2 == 0 {
		return constants.GirlFriendPositiveMentionGroup[0]
	}
	return constants.GirlFriendPositiveMentionGroup[1]
}
