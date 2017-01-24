package slutbot

import (
	"github.com/innovandalism/shodan"
)

type Module struct{}

var mod = Module{}

func init() {
	shodan.Loader.LoadModule(&mod)
}

func (_ *Module) GetIdentifier() string {
	return "slutbot"
}

func (m *Module) FlagHook() {

}

func (m *Module) Attach(session *shodan.Shodan) {
	session.GetCommandStack().RegisterCommand(&SpankmeCommand{})
	session.GetCommandStack().RegisterCommand(&SpankCommand{})
}
