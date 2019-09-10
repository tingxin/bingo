package dao

import (
	"github.com/tingxin/bingo/service/resource/model"
	g "github.com/tingxin/bingo/setting"
	gosqler "github.com/tingxin/go-sqler"
	"github.com/tingxin/go-utility/db/mysql"
)

// SaveResource used to save resource
func SaveResource(data *model.ResourceM) error {

	conn, err := mysql.GetConn(g.MetaDBConnStr)
	if err != nil {
		return err
	}

	addResource := gosqler.NewInsert("resource")
	addResource.AddValues(data.ID, data.Name, data.Desc, data.Creator, data.Editor, data.CreateTime, data.UpdateTime, data.Visible, data.VisibleTime, data.Order, data.Kind)
	sql := addResource.String()

	err = mysql.ExecuteWithConn(conn, sql)
	if err != nil {
		return err
	}

	return nil
}
