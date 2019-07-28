package data

import "github.com/kataras/iris"

func (p *service) query(ctx iris.Context) {
	ctx.StatusCode(200)
}
