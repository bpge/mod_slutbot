package slutbot

import (
	"github.com/innovandalism/shodan"
	"github.com/innovandalism/shodan/util"
)

type SpankmeCommand struct{}

func (vc *SpankmeCommand) GetNames() []string {
	return []string{"spankme"}
}

func (vc *SpankmeCommand) Invoke(ci *shodan.CommandInvocation) bool {
	msg := GenerateSpank("<@" + ci.Event.Author.ID + ">")
	_, err := ci.Session.ChannelMessageSend(ci.Event.ChannelID, msg)
	if err != nil {
		util.ReportThreadError(false, err)
	}
	return true
}
