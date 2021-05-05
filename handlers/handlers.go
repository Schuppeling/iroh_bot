package handlers

import (
	"irohbot/quotes"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var commands []string = []string{
	"!wisewords: Gets a random inspirational quote",
	"!leaves: Play a sad song",
	"!seasons: Play a party song"}

func buildEmbedFields() []*discordgo.MessageEmbedField {
	var fields []*discordgo.MessageEmbedField

	for _, command := range commands {
		part := strings.Split(command, ": ")
		field := *&discordgo.MessageEmbedField{
			Name:  part[0],
			Value: part[1],
		}

		fields = append(fields, &field)
	}

	return fields
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "!iroh":
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:  "Commands",
			Color:  0xf28a49,
			Fields: buildEmbedFields(),
		})
		break
	case "!wisewords":
		s.ChannelMessageSend(m.ChannelID, quotes.GetRandomQuote())
		break
	case "!leaves":
		s.ChannelMessageSend(m.ChannelID, "https://www.youtube.com/watch?v=f56Cbjwwv-E")
		break
	case "!seasons":
		s.ChannelMessageSend(m.ChannelID, "https://www.youtube.com/watch?v=YMwftR69o9w")
		break
	}
}
