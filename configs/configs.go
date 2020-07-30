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
	once.Do(func() {
		if err := Parse(); err != nil {
			panic(err)
		}
	})
}

type Config struct {
	App       *App
	Database  *Database
	Redis     *Redis
	Log       *Log
	Mail      *Mail
	AliyunOss *AliyunOss
}

type App struct {
	Name                 string
	Mode                 string
	Listen               string
	CacheExpire          time.Duration
	CacheCleanInterval   time.Duration
	InfoLogPath          string
	AccessLogPath        string
	ErrorLogPath         string
	StackCollectNum      int
	PasswordCost         int
	StorePrefix          string
	PostViewNum          int
	AdminEnv             string
	FileUploadEngineName string
}

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	Charset  string
}

type Redis struct {
	Host     string
	Port     int
	PORT     int
	DB       int
	Password string
}

type Log struct {
	RotationCount int
	RotationTime  time.Duration
	Level         string
}

type Mail struct {
	Host     string
	Port     int
	Username string
	Password string
}

type AliyunOss struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

var Conf *Config
var once sync.Once

func Default() *Config {
	return &Config{
		App: &App{
			Name:                 "go-project-initial",
			Mode:                 "release", //debug test release
			Listen:               ":9033",
			CacheExpire:          5 * time.Minute,
			CacheCleanInterval:   1 * time.Minute,
			InfoLogPath:          "info.log",
			AccessLogPath:        "access.log",
			ErrorLogPath:         "error.log",
			StackCollectNum:      1024,
			PasswordCost:         10,
			StorePrefix:          "/uploads",
			PostViewNum:          100,
			AdminEnv:             "prod", //local test prod
			FileUploadEngineName: "local-prefix",
		},
		Database: &Database{
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "123456",
			DBName:   "goadmin",
			Charset:  "utf8mb4",
		},
		Redis: &Redis{
			Host:     "127.0.0.1",
			Port:     6379,
			DB:       0,
			Password: "",
		},
		Log: &Log{
			RotationCount: 30,
			RotationTime:  24 * time.Hour,
			Level:         "info",
		},
		Mail: &Mail{
			Host:     "",
			Port:     0,
			Username: "",
			Password: "",
		},
		AliyunOss: &AliyunOss{
			Endpoint:        "",
			AccessKeyId:     "",
			AccessKeySecret: "",
			BucketName:      "",
		},
	}
}

func Parse() (err error) {
	Conf = Default()

	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("toml")
	vp.AddConfigPath(ProjectDir())
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
	path = path[0 : pos+len(name)]
	return
}
