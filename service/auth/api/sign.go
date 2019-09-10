package api

import (
	"fmt"

	"github.com/kataras/iris"
	m "github.com/tingxin/bingo/model"
	"github.com/tingxin/bingo/service/auth/cmd"
	"github.com/tingxin/bingo/service/auth/domain/jwt"
)

// Sign used to check service health
func Sign(ctx iris.Context) {
	cmd := cmd.SignCmd{}
	err := ctx.ReadForm(&cmd)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res := m.NewResponse("登录密码和账号不能为空", nil)
		ctx.JSON(res)
		return
	}

	authEntity := jwt.NewEntity()
	data, code, err := authEntity.Sign(cmd)

	if err != nil {
		ctx.StatusCode(code)
		res := m.NewResponse(fmt.Sprintf("登录失败：%v", err), nil)
		ctx.JSON(res)
		return
	}

	ctx.StatusCode(code)
	res := m.NewResponse("登录成功", data)
	ctx.JSON(res)
}
