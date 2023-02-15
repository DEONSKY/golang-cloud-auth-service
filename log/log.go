package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	App        string
	Module     string
	DateFormat string
}

type ProdLevelLog struct {
	App     string `json:"app"`
	Module  string `json:"module"`
	Message string `json:"message"`
}

type colors func(format string, a ...interface{}) string

func (l *Logger) debugLevelLog(message string, color colors) {
	date := time.Now()
	format := color("[%s] | [%s] - [%s]")
	cmd := fmt.Sprintf(format+": %s\n", date.Format(l.DateFormat), l.App, l.Module, message)
	io.WriteString(os.Stdout, cmd)
}

func (l *Logger) prodLevelLog(message string) {
	logObject := &ProdLevelLog{
		App:     l.App,
		Module:  l.Module,
		Message: message,
	}
	out, err := json.Marshal(logObject)
	if err != nil {
		fmt.Println(err, message)
	}
	fmt.Println(string(out))
}

func (l *Logger) log(message string, color colors) {
	// FIXME: this control should check for env variable for if it's prod (maybe staging too) it will print logs as prodLevelLog
	if true {
		l.debugLevelLog(message, color)
	} else {
		l.prodLevelLog(message)
	}
}

func (l *Logger) Trace(message string) {
	l.log(message, color.CyanString)
}

func (l *Logger) Info(message string) {
	l.log(message, color.GreenString)
}

func (l *Logger) Warning(message string) {
	l.log(message, color.YellowString)
}

func (l *Logger) Error(message string) {
	l.log(message, color.RedString)
}

func (l *Logger) Fatal(message string) {
	l.log(message, color.MagentaString)
	os.Exit(1)
}

func New(module string) *Logger {
	return &Logger{
		App:        "AUTHENTICATION_SERVICE",
		Module:     module,
		DateFormat: "2006-01-02T15:04:05.00000000000Z07:00",
	}
}

var GlobalLogger *Logger = New("MainModule")
