package api

import (
	"github.com/kataras/iris"
)

// Health used to check service health
func Health(ctx iris.Context) {
	ctx.StatusCode(200)
}
