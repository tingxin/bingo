package dao

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/tingxin/bingo/service/auth/setting"
	g "github.com/tingxin/bingo/setting"
	"github.com/tingxin/go-utility/db/mysql"
)

const (
	sqlAddUser = "INSERT INTO %s (email, name, role, password) VALUES(%s, %s,%s, %s)"
)

// AddUser used to
func AddUser(email, name, pass string) (int, error) {
	command := fmt.Sprintf(sqlAddUser, setting.Store, email, name, "normal", pass)

	conn, err := mysql.GetConn(g.MetaDBConnStr)
	if err != nil {
		return iris.StatusInternalServerError, err
	}

	err = mysql.ExecuteWithConn(conn, command)
	if err != nil {
		return iris.StatusInternalServerError, err
	}
	return iris.StatusCreated, nil
}
