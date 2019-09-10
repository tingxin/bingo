package dao

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/tingxin/bingo/service/auth/model"
	"github.com/tingxin/bingo/service/auth/setting"
	g "github.com/tingxin/bingo/setting"
	"github.com/tingxin/go-utility/db/mysql"
)

const (
	sqlSelectUserF = "SELECT email,name, role, password From %s WHERE email = %s LIMIT 1"
)

// GetUserInfoByMail used to
func GetUserInfoByMail(email string) (*model.UserInfo, int, error) {
	command := fmt.Sprintf(sqlSelectUserF, setting.Store, email)
	conn, err := mysql.GetConn(g.MetaDBConnStr)
	if err != nil {
		return nil, iris.StatusInternalServerError, err
	}

	rows, err := mysql.FetchRawWithConn(conn, command)
	if err != nil {
		return nil, iris.StatusInternalServerError, err
	}

	if len(rows) != 1 {
		return nil, iris.StatusNotFound, err
	}

	row := rows[0]

	u := model.UserInfo{}
	u.Name = string(row[1])
	u.Role = string(row[2])
	u.Eamil = email
	u.Password = string(row[3])

	return &u, iris.StatusOK, err

}
