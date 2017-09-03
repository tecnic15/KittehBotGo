package BotCommands

import (
	"github.com/NamedKitten/KittehBotGo/util"
	"gopkg.in/bwmarrin/discordgo.develop"
	"time"
	"github.com/go-errors/errors"
)

func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx *commands.Context) (error) {
	start := time.Now()
	message, err := s.ChannelMessageSend(m.ChannelID, "...")
	elapsed := time.Since(start)
	if err != nil {
		return errors.Wrap(err, 1)
	}
	s.ChannelMessageEdit(m.ChannelID, message.ID, "Ping! Time taken to send message: `"+elapsed.String()+"`.")
	return nil
}