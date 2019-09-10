package resource

import (
	"github.com/kataras/iris/core/router"
	"github.com/tingxin/bingo/service/resource/api"
)

func (p *gateway) register(r router.Party) error {
	r.Post("/", api.Create)
	// r.Post("/query/{offset:int}", p.list)
	// r.Post("/query/{offset:int}/{count:int}", p.list)

	r.Get("/health", api.Health)
	return nil
}
