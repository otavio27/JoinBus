/*
 MODULE: TTKLog.go
 AUTHOR: Leo Schneider <schleo@outlook.com>
 DATE  : August 2017
 INFO  : This module handles Log outputs
*/

package ttktools

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

//Log levels
const (
	DEBUG   = 0
	FATAL   = 1
	ERROR   = 2
	WARNING = 3
	INFO    = 4
)

// TTKLog Process application Logs
type ttklog struct {
	level   int
	logFile string
	console bool
	logchan chan logmsg
}

// LogMsg ...
type logmsg struct {
	module  string
	message string
	level   int
}

// Log Setup defaults
func Log(LogLevel int, LogFile string, Console bool) *ttklog {
	t := ttklog{}
	t.level = LogLevel
	t.logFile = LogFile
	t.console = Console
	t.logchan = make(chan logmsg)
	go t.logger()
	return &t
}

// Msg writes a message to the log queue
func (t *ttklog) Msg(ctx context.Context, module, message string, level int) {
	t.logchan <- logmsg{module, message, level}
}

// Logger reads the log queue and writes the messages to disk
func (t *ttklog) logger() {
	f, err := os.OpenFile(t.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print(GRAY + "[" + RED + "FATAL" + GRAY + "] " + NC + err.Error() + "\n")
		os.Exit(2)
	} else {
		log.SetOutput(f)
		for {
			m := <-t.logchan
			if m.level >= t.level {
				logLevel := [...]string{"DEBUG", "FATAL", "ERROR", "WARN ", "INFO "}
				colorprefix := [...]string{PURPLE, RED, RED, LGREEN, GREEN}
				prefix := logLevel[m.level]
				color := colorprefix[m.level]
				msg := m.module + " - " + m.message
				if t.console {
					i, err := strconv.ParseInt(strconv.Itoa(int(time.Now().Unix())), 10, 64)
					if err != nil {
						log.SetPrefix(prefix + " ")
						log.Println(err.Error())
					}
					fmt.Print(GRAY + "[" + color + prefix + GRAY + "] " + NC + time.Unix(i, 0).Format("15:04:05") + " - " + m.message + "\n")
				}
				log.SetPrefix(prefix + " ")
				log.Println(msg)
				if m.level == FATAL {
					os.Exit(1)
				}
			}
		}
	}
}
