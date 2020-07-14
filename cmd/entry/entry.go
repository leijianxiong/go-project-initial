package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-project-initial/configs"
	"os"
	"os/signal"
	"strconv"
)

func main() {
	gin.SetMode(configs.Conf.App.Mode)

	r := gin.New()
	e := engine.Default()

	cfg := NewDefault()

	template.AddComp(chartjs.NewChart())

	// customize a plugin

	//e.AddGenerators(user.Generators)
	//add plugin

	e.AddDisplayFilterXssJsFilter()

	if err := e.AddConfig(*cfg).Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")
	// customize your pages

	//后台欢迎页
	e.HTML("GET", "/admin", datamodel.GetContent)

	go func() {
		if err := endless.ListenAndServe(configs.Conf.App.Listen, r); err != nil {
			log.Infoln("ListenAndServe err:", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	e.MysqlConnection().Close()
}

func NewDefault() *config.Config {
	return &config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host: configs.Conf.Database.Host,
				Port: strconv.FormatInt(int64(configs.Conf.Database.Port), 10),
				User: configs.Conf.Database.Username,
				Pwd: configs.Conf.Database.Password,
				Name: configs.Conf.Database.DBName,
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     config.DriverMysql,
			},
		},
		UrlPrefix: "admin",
		Store: config.Store{
			Path: configs.ProjectDir() + "/uploads",
			Prefix: "uploads",
		},
		Language:           language.CN,
		IndexUrl:           "/",
		Debug:              configs.Conf.App.Mode != gin.ReleaseMode,
		AccessAssetsLogOff: true,
		Animation: config.PageAnimation{
			Type: "fadeInUp",
		},
		ColorScheme: adminlte.ColorschemeSkinBlack,
		// log file absolute path
		InfoLogPath: configs.ProjectDir() + "/logs/" + configs.Conf.App.InfoLogPath,
		AccessLogPath: configs.ProjectDir() + "/logs/" + configs.Conf.App.AccessLogPath,
		ErrorLogPath: configs.ProjectDir() + "/logs/" + configs.Conf.App.ErrorLogPath,
	}
}
