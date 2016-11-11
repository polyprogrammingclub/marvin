package awake

import (
	"net/url"
	"time"

	"github.com/riking/homeapi/marvin"
)

func init() {
	marvin.RegisterModule(NewAwakeModule)
}

const Identifier = "core"

type AwakeModule struct {
	team   marvin.Team
	quit   chan struct{}
	ticker *time.Ticker
}

func NewAwakeModule(t marvin.Team) marvin.Module {
	mod := &AwakeModule{team: t}
	return mod
}

func (mod *AwakeModule) Identifier() marvin.ModuleID {
	return Identifier
}

func (mod *AwakeModule) Load(t marvin.Team) {
}

func (mod *AwakeModule) Enable(t marvin.Team) {
	mod.ticker = time.NewTicker(20 * time.Minute)
	mod.quit = make(chan struct{})
	go mod.onTick()
}

func (mod *AwakeModule) Disable(t marvin.Team) {
	mod.ticker.Stop()
	close(mod.quit)

}

func (mod *AwakeModule) onTick() {
	for {
		select {
		case <-mod.ticker.C:
			mod.team.SlackAPIPost("users.setActive", url.Values{})
		case <-mod.quit:
			return
		}
	}
}
