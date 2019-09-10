package dao

import (
	"github.com/tingxin/bingo/service/resource/model"
	g "github.com/tingxin/bingo/setting"
	gosqler "github.com/tingxin/go-sqler"
	"github.com/tingxin/go-utility/db/mysql"
)

// SaveResourceFields used to save resource fields
func SaveResourceFields(resourceID int64, fields []*model.FieldM) error {
	conn, err := mysql.GetConn(g.MetaDBConnStr)
	if err != nil {
		return err
	}
	addFields := gosqler.NewInsert("fields")
	addFields.AddColumns("view_id", "title", "name", "desc", "indicator_type", "group", "order", "selected", "create_time", "update_time")
	for _, data := range fields {
		addFields.AddValues(resourceID, data.Title, data.Name, data.Desc, data.IndicatorType, data.Group, data.Order, data.Selected, data.CreateTime, data.UpdateTime)
	}
	sql := addFields.String()

	err = mysql.ExecuteWithConn(conn, sql)
	if err != nil {
		return err
	}

	return nil
}
