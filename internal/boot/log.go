package boot

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"go-project-initial/configs"
	"path/filepath"
	"time"
)

func init() {
	p := fmt.Sprintf(configs.ProjectDir()+"/logs/%s.log", "default")
	p, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}
	writer, _ := rotatelogs.New(
		//p+".%Y%m%d%H%M",
		p+".%Y%m%d",
		rotatelogs.WithLinkName(p),
		//rotatelogs.WithMaxAge(86400*time.Second),
		rotatelogs.WithRotationCount(uint(configs.Conf.Log.RotationCount)),
		rotatelogs.WithRotationTime(configs.Conf.Log.RotationTime),
	)

	log.SetLevel(log.DebugLevel)
	log.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			log.DebugLevel: writer,
			log.InfoLevel:  writer,
			log.WarnLevel:  writer,
			log.ErrorLevel: writer,
		},
		&log.TextFormatter{
			DisableQuote: true,
		},
	))
}

var logs map[string]*log.Logger

func NewLogger(logName string) *log.Logger {
	if logName == "" {
		logName = "default"
	}

	if logs == nil {
		logs = make(map[string]*log.Logger)
	}

	if l, ok := logs[logName]; ok {
		return l
	}

	p := fmt.Sprintf(configs.ProjectDir()+"/logs/%s.log", logName)
	p, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}

	writer, _ := rotatelogs.New(
		//p+".%Y%m%d%H%M",
		p+".%Y%m%d",
		rotatelogs.WithLinkName(p),
		//rotatelogs.WithMaxAge(86400*time.Second),
		rotatelogs.WithRotationCount(uint(configs.Conf.Log.RotationCount)),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	l := log.New()
	l.SetLevel(log.DebugLevel)
	l.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			log.DebugLevel: writer,
			log.InfoLevel:  writer,
			log.WarnLevel:  writer,
			log.ErrorLevel: writer,
		},
		&log.TextFormatter{
			DisableQuote: true,
		},
	))

	logs[logName] = l

	return l
}
