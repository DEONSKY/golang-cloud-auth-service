package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	App    string
	Module string
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
	cmd := fmt.Sprintf(format+": %s\n", date.Format(time.RFC3339Nano), l.App, l.Module, message)
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

func New(app string, module string) *Logger {
	logger := Logger{
		App:    app,
		Module: module,
	}

	return &logger
}

var GlobalLogger *Logger = New("AUTHENTICATION_SERVICE", "MainModule")

func main() {
	l := New("AUTHENTICATION_SERVICE", "LoggerModule")

	l.Info("Hello World!")
	l.Warning("Fire started!!")
	l.Error("WORLD ON FIRE!")
}
