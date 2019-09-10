package auth

import (
	"github.com/kataras/iris/core/router"
	"github.com/tingxin/bingo/service/auth/api"
)

func (p *gateway) register(r router.Party) error {

	r.Get("/health", api.Health)
	r.Post("/sign", api.Sign)
	r.Post("/signup", api.SignUp)
	return nil
}
