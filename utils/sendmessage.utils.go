package utils

import "github.com/bwmarrin/discordgo"

func SendMessage(session *discordgo.Session, channelID, msg string) error {
	if _, err := session.ChannelMessageSend(channelID, msg); err != nil {
		return err
	}
	return nil
}

func SendReply(session *discordgo.Session, channelID, msg string, ref *discordgo.MessageReference) error {
	if _, err := session.ChannelMessageSendReply(channelID, msg, ref); err != nil {
		return err
	}
	return nil
}
