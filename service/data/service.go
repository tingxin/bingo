package data

import (
	"fmt"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/tingxin/bingo/common/meta"
	"github.com/tingxin/bingo/model"
	s "github.com/tingxin/bingo/service"
)

type service struct {
	s.Server
	domain string
	port   int16
}

// New used to create a new data service
func New() s.Server {
	instance := service{}
	instance.domain = "data"
	instance.port = 5025
	return &instance
}

func (p *service) Run() error {
	meta.Run()

	api := iris.Default()
	api.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(model.NewResponse(iris.StatusNotFound, "404 没有找到你想要的资源！", nil))
	})
	api.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.WriteString("网络异常，请重试！")
	})

	//"github.com/iris-contrib/middleware/cors"
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	// api.Use(irisyaag.New())
	appVersion := "/v1"
	v := api.Party(appVersion, crs).AllowMethods(iris.MethodOptions)
	{
		p.register(v)
	}

	address := fmt.Sprintf("0.0.0.0:%d", p.port)
	api.Run(iris.Addr(address))
	return nil
}

func (p *service) Stop() error {
	return nil

}

func (p *service) register(r router.Party) error {
	r.Post("/query", p.query)
	r.Post("/query/{offset:int}", p.query)
	r.Post("/query/{offset:int}/{count:int}", p.query)

	r.Get("/health", p.health)
	return nil
}

// health used to check service health
func (p *service) health(ctx iris.Context) {
	ctx.StatusCode(200)
}
