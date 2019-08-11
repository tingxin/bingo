package auth

import (
	"os"

	"github.com/kataras/iris"
	"github.com/tingxin/bingo/service/auth/normal"
)

var (
	secret []byte
)

func init() {
	key := os.Getenv("Secret")
	if key == "" {
		panic("failed to get env Secret")
	}
	secret = []byte(key)
}

func (p *service) sign(ctx iris.Context) {
	// authKind := os.Getenv("AuthKind")
	var method func(iris.Context, []byte)
	method = normal.Sign
	method(ctx, secret)
}
