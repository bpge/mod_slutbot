package slutbot

import (
	"github.com/innovandalism/shodan"
	"strings"
)

type PetCommand struct{}

func (vc *PetCommand) GetName() string {
	return "pet"
}

func (vc *PetCommand) Invoke(ci *shodan.CommandInvocation) bool {
	var target string
	switch {
	case strings.HasPrefix(ci.Arguments[0], "<@") && strings.HasSuffix(ci.Arguments[0], ci.Event.Author.ID + ">"):
		ci.Session.ChannelMessageSend(ci.Event.ChannelID, "You may not pet yourself.")
		return true
	case strings.HasPrefix(ci.Arguments[0], "<@") && strings.HasSuffix(ci.Arguments[0], ">"):
		target = ci.Arguments[0]
	default:
		ci.Session.ChannelMessageSend(ci.Event.ChannelID, "Invalid target.")
		return true
	}
	msg := GeneratePat(target)
	ci.Session.ChannelMessageSend(ci.Event.ChannelID, msg)
	return true
}