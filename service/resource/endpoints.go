package resource

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/kataras/iris"
	"github.com/tingxin/bingo/setting"
	"github.com/tingxin/go-utility/db/mysql"
	"github.com/tingxin/go-utility/log"
)

// list used to check service health
func (p *service) list(ctx iris.Context) {
	ctx.StatusCode(200)
}

// health used to check service health
func (p *service) health(ctx iris.Context) {
	ctx.StatusCode(200)
}

// health used to check service health
func (p *service) prepare() {
	conn, err := mysql.GetConn(setting.MetaDBConnStr)
	if err != nil {
		panic(err)
	}
	rootFolder, _ := os.Getwd()
	prerequisites := fmt.Sprintf("%s/service/resource/assets/prerequisite", rootFolder)
	files, err := ioutil.ReadDir(prerequisites)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileName := file.Name()
		filePath := fmt.Sprintf("%s/%s", prerequisites, fileName)
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		sql := string(data)
		log.INFO.Printf("Run\n %s \n", sql)
		mysql.ExecuteWithConn(conn, sql)
		log.INFO.Printf("Done %s \n", fileName)
		time.Sleep(time.Second)
	}
}
