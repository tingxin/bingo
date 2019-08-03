package auth

import (
	"github.com/kataras/iris"
)

// health used to check service health
func (p *service) health(ctx iris.Context) {
	ctx.StatusCode(200)
}
