package api

import (
	"fmt"

	"github.com/kataras/iris"
	m "github.com/tingxin/bingo/model"
	"github.com/tingxin/bingo/service/resource/cmd"
	"github.com/tingxin/bingo/service/resource/domain"
	"github.com/tingxin/go-utility/log"
)

// Create used to create new resource on server
func Create(ctx iris.Context) {

	cmd := cmd.NewResourceCmd{}
	err := ctx.ReadJSON(&cmd)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res := m.NewResponse(fmt.Sprintf("数据格式有误: %v", err), nil)
		ctx.JSON(res)
		return
	}

	err = domain.AddResource(cmd)
	if err != nil {
		msg := fmt.Sprintf("failed to add resource due to: \n%v\n", err)
		log.ERROR.Printf(msg)
		ctx.StatusCode(iris.StatusInternalServerError)
		res := &m.Response{
			Msg:    "创建视图失败",
			Data:   nil,
			Level:  3,
			Detail: msg,
		}
		ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	res := &m.Response{
		Msg:  "创建视图成功",
		Data: nil,
	}
	ctx.JSON(res)

}
