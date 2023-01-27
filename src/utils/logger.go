package main

import (
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

type colors func(format string, a ...interface{}) string

func (l *Logger) log(message string, color colors) {
	date := time.Now()
	format := color("[%s] | [%s] - [%s]")
	cmd := fmt.Sprintf(format+": %s\n", date.Format(time.RFC3339Nano), l.App, l.Module, message)
	io.WriteString(os.Stdout, cmd)
}

func (l *Logger) Info(message string) {
	l.log(message, color.BlueString)
}

func (l *Logger) Warning(message string) {
	l.log(message, color.YellowString)
}

func (l *Logger) Error(message string) {
	l.log(message, color.RedString)
}

func New(app string, module string) *Logger {
	logger := Logger{
		App:    app,
		Module: module,
	}

	return &logger
}

func main() {
	l := New("AUTHENTICATION_SERVICE", "LoggerModule")

	l.Info("Hello World!")
	l.Warning("Fire started!!")
	l.Error("WORLD ON FIRE!")
}
