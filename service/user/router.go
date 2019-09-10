package user

import (
	"github.com/kataras/iris/core/router"
	"github.com/tingxin/bingo/service/user/api"
)

func (p *gateway) register(r router.Party) error {
	r.Get("/health", api.Health)
	return nil
}
