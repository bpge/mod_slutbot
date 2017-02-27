package slutbot

import (
	"github.com/innovandalism/shodan"
	"github.com/innovandalism/shodan/util"
)

type SpankCommand struct{}

func (vc *SpankCommand) GetNames() []string {
	return []string{"spank"}
}

func (vc *SpankCommand) Invoke(ci *shodan.CommandInvocation) error {
	var target string
	if len(ci.Arguments) > 0 {
		target = ci.Arguments[0]
	} else {
		err := ci.Helpers.Reply("Shodan tired. Shodan sleep. :shodan:")
		if err != nil {
			return util.WrapError(err)
		}
	}
	msg := GenerateSpank(target)
	err := ci.Helpers.Reply(msg)
	if err != nil {
		return util.WrapError(err)
	}
	return nil
}
