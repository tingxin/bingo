package resource

import (
	"fmt"

	"github.com/kataras/iris"
	m "github.com/tingxin/bingo/model"
	"github.com/tingxin/bingo/service/resource/common"
	"github.com/tingxin/bingo/service/resource/model"
	g "github.com/tingxin/bingo/setting"
	gosqler "github.com/tingxin/go-sqler"
	"github.com/tingxin/go-utility/db/mysql"
	"github.com/tingxin/go-utility/log"
)

var (
	addResource gosqler.Insert
	addFields   gosqler.Insert
)

func init() {
	addResource = gosqler.NewInsert("resource")
	addFields = gosqler.NewInsert("fields")
	addFields.AddColumns("view_id", "title", "name", "desc", "indicator_type", "group", "order", "selected", "create_time", "update_time")

}

func (p *service) create(ctx iris.Context) {
	defer func() {
		if err := recover(); err != nil {
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
		}
	}()

	resourceEntity := &model.ResourceM{}
	err := ctx.ReadJSON(resourceEntity)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res := m.NewResponse(fmt.Sprintf("数据格式有误: %v", err), nil)
		ctx.JSON(res)
		return
	}

	id := common.NewIntUID()

	doneSaveResource := make(chan error)
	doneSaveFields := make(chan error)

	go saveResource(id, resourceEntity, doneSaveResource)
	go saveFields(id, resourceEntity.Fields, doneSaveFields)

	err = <-doneSaveFields

	if err != nil {
		panic(err)
	}

	err = <-doneSaveResource
	if err != nil {
		panic(err)
	}

}

func saveResource(resourceID int64, data *model.ResourceM, done chan<- error) {

	conn, err := mysql.GetConn(g.MetaDBConnStr)
	if err != nil {
		done <- err
		return
	}

	addResource.ClearValues()
	addResource.AddValues(resourceID, data.Name, data.Desc, data.Creator, data.Editor, data.CreateTime, data.UpdateTime, data.Visible, data.VisibleTime, data.Order, data.Kind)
	sql := addResource.String()

	err = mysql.ExecuteWithConn(conn, sql)
	if err != nil {
		done <- err
		return
	}

	close(done)
}

func saveFields(resourceID int64, fields []*model.FieldM, done chan<- error) {
	conn, err := mysql.GetConn(g.MetaDBConnStr)
	if err != nil {
		done <- err
		return
	}

	addFields.ClearValues()
	for _, data := range fields {
		addFields.AddValues(resourceID, data.Title, data.Name, data.Desc, data.IndicatorType, data.Group, data.Order, data.Selected, data.CreateTime, data.UpdateTime)
	}
	sql := addFields.String()

	err = mysql.ExecuteWithConn(conn, sql)
	if err != nil {
		done <- err
		return
	}

	close(done)
}
