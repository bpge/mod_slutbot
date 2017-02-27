package slutbot

import (
	"github.com/innovandalism/shodan"
	"strings"
	"github.com/innovandalism/shodan/util"
)

type PetCommand struct{}

func (vc *PetCommand) GetNames() []string {
	return []string{"pet"}
}

func (vc *PetCommand) Invoke(ci *shodan.CommandInvocation) error {
	var target string
	switch {
	case strings.HasPrefix(ci.Arguments[0], "<@") && strings.HasSuffix(ci.Arguments[0], ci.Event.Author.ID + ">"):
		err := ci.Helpers.Reply("You may not pet yourself.")
		if err != nil {
			return util.WrapError(err)
		}
		return nil
	case strings.HasPrefix(ci.Arguments[0], "<@") && strings.HasSuffix(ci.Arguments[0], ">"):
		target = ci.Arguments[0]
	default:
		err := ci.Helpers.Reply("Invalid target.")
		if err != nil {
			return util.WrapError(err)
		}
		return nil
	}
	msg := GeneratePat(target)
	err := ci.Helpers.Reply(msg)
	if err != nil {
		return util.WrapError(err)
	}
	return nil
}