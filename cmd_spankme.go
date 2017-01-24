package slutbot

import (
	"github.com/innovandalism/shodan"
)

type SpankmeCommand struct{}

func (vc *SpankmeCommand) GetName() string {
	return "spankme"
}

func (vc *SpankmeCommand) Invoke(ci *shodan.CommandInvocation) bool {
	msg := GenerateSpank("<@" + ci.Event.Author.ID + ">")
	ci.Session.ChannelMessageSend(ci.Event.ChannelID, msg)
	return true
}
