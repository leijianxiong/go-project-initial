package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-project-initial/configs"
	"math"
	"net/http"
	"runtime"
)

const ErrnoFail = 400
const ErrnoSuccess = 0

func Fail(errmsg string, errnos ...int) gin.H {
	errno := ErrnoFail
	if len(errnos) > 0 {
		errno = errnos[0]
	}

	return gin.H{
		"errno":  errno,
		"errmsg": errmsg,
	}
}

func Success(data interface{}) gin.H {
	return gin.H{
		"errno":  ErrnoSuccess,
		"errmsg": "",
		"data":   data,
	}
}

func Recover(c *gin.Context) {
	if p := recover(); p != nil {
		errmsg := fmt.Sprintf("%s", p)
		pc, fn, line, _ := runtime.Caller(1)
		log.Errorf("api(%s)[error] in %s[%s:%d] %v", c.Request.URL.Path, runtime.FuncForPC(pc).Name(), fn, line, errmsg)

		buf := make([]byte, configs.Conf.App.StackCollectNum)
		buf = buf[:int(math.Min(float64(configs.Conf.App.StackCollectNum), float64(runtime.Stack(buf, true))))]
		log.Errorf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)

		c.JSON(http.StatusBadRequest, Fail(errmsg))
	}
}
