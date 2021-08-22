package logger

import (
	"github.com/withmandala/go-log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	Logger = log.New(os.Stderr).WithColor()
}
