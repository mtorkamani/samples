package log

import (
	"fmt"
	"log"
	"time"
)

type Logger interface {
	Log(message string, args ...interface{})
	Error(err error)
	Fatal(err error)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) Log(message string, args ...interface{}) {
	go func() {
		lmsg := format("Info", message, args...)
		fmt.Println(lmsg)
	}()
}

func (l *logger) Fatal(err error) {
	go func() {
		lmsg := format("Fatal", err.Error())
		log.Fatalln(lmsg)
	}()
}

func (l *logger) Error(err error) {
	go func() {
		lmsg := format("Error", err.Error())
		fmt.Println(lmsg)
	}()
}

func format(head, message string, args ...interface{}) string {
	fmsg := message
	if len(args) > 0 {
		fmsg = fmt.Sprintf(message, args)
	}
	now := time.Now().Local()
	return fmt.Sprintf("[%s] - %s - %s", head, now.Format("02-Jan-2006 15:04:05"), fmsg)
}
