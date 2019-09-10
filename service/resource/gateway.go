package resource

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/tingxin/bingo/middleware/auth"
	"github.com/tingxin/bingo/model"
	s "github.com/tingxin/bingo/service"
	"github.com/tingxin/bingo/setting"
	"github.com/tingxin/go-utility/db/mysql"
	"github.com/tingxin/go-utility/log"
)

type gateway struct {
	s.Server
	domain string
	port   int16
}

// New used to create a new data service
func New() s.Server {
	instance := gateway{}
	instance.domain = "resource"
	instance.port = 5027
	return &instance
}

func (p *gateway) Run() error {
	p.prepare()
	api := iris.Default()
	api.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(model.NewResponse("404 没有找到你想要的资源！", nil))
	})
	api.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.WriteString("网络异常，请重试！")
	})

	//"github.com/iris-contrib/middleware/cors"
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	// api.Use(irisyaag.New())
	appVersion := "/v1"
	v := api.Party(appVersion, crs).AllowMethods(iris.MethodOptions)
	{
		p.register(v)
	}
	api.Use(auth.Run)
	address := fmt.Sprintf("0.0.0.0:%d", p.port)
	api.Run(iris.Addr(address))
	return nil
}

func (p *gateway) Stop() error {
	return nil
}

// health used to check service health
func (p *gateway) prepare() {
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
