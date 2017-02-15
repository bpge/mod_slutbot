package slutbot

import (
	"github.com/innovandalism/shodan"
	"github.com/innovandalism/shodan/util"
)

type SpankCommand struct{}

func (vc *SpankCommand) GetNames() []string {
	return []string{"spank"}
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
	_, err := ci.Session.ChannelMessageSend(ci.Event.ChannelID, msg)
	if err != nil {
		util.ReportThreadError(false, err)
	}
	return true
}
