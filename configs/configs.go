package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func init() {
	new(sync.Once).Do(func() {
		if err := Parse(); err != nil {
			panic(err)
		}
	})
}

type Config struct {
	App *App
	Database *Database
	Redis *Redis
	Log *Log
	Mail *Mail
}

type App struct {
	Name string
	Mode string
	Listen string
	CacheExpire time.Duration
	InfoLogPath string
	AccessLogPath string
	ErrorLogPath string
	StackCollectNum int
}

type Database struct {
	Host string
	Port int
	Username string
	Password string
	DBName string
	Charset string
}

type Redis struct {
	Host string
	Port int
	DB int
	Password string
}

type Log struct {
	RotationCount int
	RotationTime time.Duration
	Level string
}

type Mail struct {
	Host string
	Port int
	Username string
	Password string
}

var Conf *Config

func Default() *Config {
	return &Config{
		App: &App{
			Name: "go-project-initial",
			Mode: "release",
			Listen: ":9033",
			CacheExpire: 5 * time.Minute,
			InfoLogPath: "info.log",
			AccessLogPath: "access.log",
			ErrorLogPath: "error.log",
			StackCollectNum: 1024,
		},
		Database: &Database{
			Host: "127.0.0.1",
			Port: 3306,
			Username: "root",
			Password: "123456",
			DBName: "goadmin",
			Charset: "utf8mb4",
		},
		Redis: &Redis{
			Host: "127.0.0.1",
			Port: 6379,
			DB: 0,
			Password: "",
		},
		Log: &Log{
			RotationCount: 30,
			RotationTime: 24 * time.Hour,
			Level: "info",
		},
		Mail: &Mail{
			Host: "smtp.qq.com",
			Port: 465,
			Username: "1617189289@qq.com",
			Password: "krfnreuffszkfbbj",
		},
	}
}

func Parse() (err error) {
	Conf = Default()

	vp := viper.New()
	vp.SetConfigName(".env")                // name of config file (without extension)
	vp.SetConfigType("ini")                 // REQUIRED if the config file does not have the extension in the name
	vp.AddConfigPath(ProjectDir()) // optionally look for config in the working directory
	if err = vp.ReadInConfig(); err != nil {
		err = fmt.Errorf("Read config file err: %s\n", err)
	}
	if err = vp.Unmarshal(&Conf); err != nil {
		err = fmt.Errorf("Unmarshal config file err: %s\n", err)
	}
	return
}

//项目目录
func ProjectDir() (path string) {
	path, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	name := Conf.App.Name
	pos := strings.Index(path, name)
	path = path[0: pos + len(name)]
	return
}
