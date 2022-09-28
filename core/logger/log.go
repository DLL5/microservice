package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// 全局log，设置log格式
var Log = log.New()

func InitLog() {
	Log.Out = os.Stdout
	log.SetFormatter(&log.TextFormatter{})
	//log.WithTime()
}
