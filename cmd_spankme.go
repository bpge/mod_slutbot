package slutbot

import (
	"github.com/innovandalism/shodan"
	"github.com/innovandalism/shodan/util"
)

type SpankmeCommand struct{}

func (vc *SpankmeCommand) GetNames() []string {
	return []string{"spankme"}
}

func (vc *SpankmeCommand) Invoke(ci *shodan.CommandInvocation) error {
	msg := GenerateSpank("<@" + ci.Event.Author.ID + ">")
	err := ci.Helpers.Reply(msg)
	if err != nil {
		return util.WrapError(err)
	}
	return nil
}
