package slutbot

import (
	"github.com/innovandalism/shodan"
)

type SpankCommand struct{}

func (vc *SpankCommand) GetName() string {
	return "spank"
}

func (vc *SpankCommand) Invoke(ci *shodan.CommandInvocation) bool {
	var target string
	if len(ci.Arguments) > 0 {
		target = ci.Arguments[0]
	} else {
		ci.Session.ChannelMessageSend(ci.Event.ChannelID, "Shodan tired. Shodan sleep. :shodan:")
		return true
	}
	msg := GenerateSpank(target)
	ci.Session.ChannelMessageSend(ci.Event.ChannelID, msg)
	return true
}
