package responseWraper

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type ResponseData struct {
	Code        int
	Description string
	Data        any
}

const (
	nonCode    int = 0
	normalCode int = 200
)

func ResponseJson(c *gin.Context, Body any) {
	c.JSON(normalCode, &ResponseData{
		Code: nonCode,
		Data: Body,
	})
}

func ResponseJsonErrCode(c *gin.Context, code int) {
	c.JSON(normalCode, &ResponseData{
		Code:        code,
		Description: errCodeToExplain[code],
	})
}

type ErrCodeToExplain map[int]string

var errCodeToExplain = make(ErrCodeToExplain, 0)

func InitErrorCodeToExplain() {
	file, err := ioutil.ReadFile("./conf/errorCode.conf")
	if err != nil {
		log.Fatal("读取错误码配置文件失败")
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		l := strings.Split(line, "=")
		if len(l) != 2 {
			log.Println("Read ErrorCode failed! You should check and fix it!")
			continue
		}
		code, explain := l[0], l[1]
		code = strings.TrimSpace(code)
		explain = strings.TrimSpace(explain)
		codeInt, err := strconv.Atoi(code)
		if err != nil {
			log.Println("ErrorCode file exist wrong format! You should check and fix it!")
			continue
		}
		if _, ok := errCodeToExplain[codeInt]; ok {
			log.Println("ErrorCode exists repeated codeError! You should check and fix it!")
			continue
		}
		errCodeToExplain[codeInt] = explain
	}
}
