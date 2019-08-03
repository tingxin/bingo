package normal

import (
	"crypto/sha256"
	"fmt"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	m "github.com/tingxin/bingo/model"
	"github.com/tingxin/bingo/service/auth/model"
	"github.com/tingxin/bingo/service/auth/setting"
	g "github.com/tingxin/bingo/setting"
	"github.com/tingxin/go-utility/db/mysql"
	"github.com/tingxin/go-utility/log"
)

const (
	sqlSelectUserF = "SELECT email,name, role, password From %s WHERE email = %s LIMIT 1"
	sqlAddUser     = "INSERT INTO %s (email, name, role, password) VALUES(%s, %s,%s, %s)s"
)

// Sign by email and password
func Sign(ctx iris.Context, secret []byte) {
	user := &model.User{}
	err := ctx.ReadForm(user)
	if err != nil {

		ctx.StatusCode(iris.StatusBadRequest)
		res := m.NewResponse("登录密码和账号不能为空", nil)
		ctx.JSON(res)
		return
	}

	command := fmt.Sprintf(sqlSelectUserF, setting.Store, user.Eamil)

	conn, err := mysql.GetConn(g.MetaDBConnStr)
	if err != nil {
		log.ERROR.Printf("failed to connect db due to: \n%v\n", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		res := m.NewResponse("无法连接认证服务器数据库", nil)
		ctx.JSON(res)
		return
	}

	rows, err := mysql.FetchRawWithConn(conn, command)
	if err != nil {
		log.ERROR.Printf("failed to query: \n%s\n due to: \n%v\n", command, err)
		ctx.StatusCode(iris.StatusInternalServerError)
		res := m.NewResponse("执行用户查询失败", nil)
		ctx.JSON(res)
		return
	}

	if len(rows) != 1 {
		ctx.StatusCode(iris.StatusNotFound)
		res := m.NewResponse("该用户不存在", nil)
		ctx.JSON(res)
		return
	}

	row := rows[0]
	email := string(row[0])
	username := string(row[1])
	role := string(row[2])
	password := string(row[3])

	sha := sha256.New()
	sha.Write([]byte(user.Password))

	passHash := fmt.Sprintf("%x", sha.Sum(nil))

	if passHash != password {
		ctx.StatusCode(iris.StatusUnauthorized)
		res := m.NewResponse(fmt.Sprintf("账号:%s 密码错误", email), nil)
		ctx.JSON(res)
		return
	}

	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Eamil,
		"name":  username,
		"role":  role,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(secret)
	ctx.StatusCode(iris.StatusOK)
	res := m.NewResponse("登录成功", tokenString)
	ctx.JSON(res)
	return
}

// SignUp by email and password
func SignUp(ctx iris.Context, secret []byte) {
	user := &model.UserInfo{}
	err := ctx.ReadForm(user)
	if err != nil {

		ctx.StatusCode(iris.StatusBadRequest)
		res := m.NewResponse("登录密码和账号不能为空", nil)
		ctx.JSON(res)
		return
	}

	command := fmt.Sprintf(sqlAddUser, setting.Store, user.Eamil, user.Name, "normal", user.Password)

	conn, err := mysql.GetConn(g.MetaDBConnStr)
	if err != nil {
		log.ERROR.Printf("failed to connect db due to: \n%v\n", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		res := m.NewResponse("无法连接认证服务器数据库", nil)
		ctx.JSON(res)
		return
	}

	err = mysql.ExecuteWithConn(conn, command)
	if err != nil {
		log.ERROR.Printf("failed to query: \n%s\n due to: \n%v\n", command, err)
		ctx.StatusCode(iris.StatusInternalServerError)
		res := m.NewResponse("添加用户失败", nil)
		ctx.JSON(res)
		return
	}

	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Eamil,
		"name":  user.Name,
		"role":  "normal",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(secret)
	ctx.StatusCode(iris.StatusOK)
	res := m.NewResponse("登录成功", tokenString)
	ctx.JSON(res)
	return
}
