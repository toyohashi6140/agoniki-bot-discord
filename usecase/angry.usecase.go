package usecase

import (
	"github.com/bwmarrin/discordgo"
	"github.com/toyohashi6140/discord-bot/constants"
	"github.com/toyohashi6140/discord-bot/models"
	"github.com/toyohashi6140/discord-bot/utils"
)

type angry struct {
	agoniki *models.Agoniki
}

func NewAngry(a *models.Agoniki) Responser {
	return &angry{
		agoniki: a,
	}
}

// Response 怒りフラグに応じてメッセージを投稿させる
func (a *angry) Response(session *discordgo.Session, event *discordgo.MessageCreate) error {
	// 怒りカウントをとって一定以上ならめっちゃ怒る。
	a.agoniki.Anger(event.Content)
	// Angry XOR SeriouslyAngryは必ず真になる
	if a.agoniki.Angry {
		texts := a.getAngryText()
		for _, text := range texts {
			if err := utils.SendReply(session, event.ChannelID, text, event.Reference()); err != nil {
				return err
			}
		}
	}
	if a.agoniki.SeriouslyAngry {
		texts := a.getSeriouslyAngryText()
		for _, text := range texts {
			if err := utils.SendReply(session, event.ChannelID, text, event.Reference()); err != nil {
				return err
			}
		}
	}
	if a.agoniki.ALittleAngry {
		texts := a.getALittleAngryText()
		for _, text := range texts {
			if err := utils.SendReply(session, event.ChannelID, text, event.Reference()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *angry) getAngryText() []string {
	i := utils.RandNumber(100)
	if i%2 == 0 {
		return constants.AngryMentionGroup[0]
	} else if i%3 == 0 {
		return constants.AngryMentionGroup[1]
	} else if i%5 == 0 {
		return constants.AngryMentionGroup[2]
	} else if i%7 == 0 {
		return constants.AngryMentionGroup[3]
	}
	return constants.AngryMentionGroup[4]
}

func (a *angry) getSeriouslyAngryText() []string {
	i := utils.RandNumber(100)
	if i%2 == 0 {
		switchYouKnowThatToComeOnYou(constants.SeriouslyAngryMentionGroup[0])
		return constants.SeriouslyAngryMentionGroup[0]
	} else if i%3 == 0 {
		return constants.SeriouslyAngryMentionGroup[1]
	} else if i%5 == 0 {
		return constants.SeriouslyAngryMentionGroup[2]
	} else if i%7 == 0 {
		return constants.SeriouslyAngryMentionGroup[3]
	}
	return constants.SeriouslyAngryMentionGroup[4]
}

func (a *angry) getALittleAngryText() []string {
	i := utils.RandNumber(100)
	if i%2 == 0 {
		return constants.ALittleAngryMentionGroup[0]
	} else if i%3 == 0 {
		return constants.ALittleAngryMentionGroup[1]
	} else if i%5 == 0 {
		return constants.ALittleAngryMentionGroup[2]
	}
	switchYouKnowThatToComeOnYou(constants.ALittleAngryMentionGroup[3])
	return constants.ALittleAngryMentionGroup[3]
}

// 1/2の割合でお前さあ・・・に変更。副作用を起こすため戻り値は返さない
func switchYouKnowThatToComeOnYou(tests []string) {
	i := utils.RandNumber(100)
	if i%2 == 0 {
		tests[0] = constants.ComeOnYou
	} else if tests[0] == constants.ComeOnYou {
		tests[0] = constants.YouKnowThat
	}
}
