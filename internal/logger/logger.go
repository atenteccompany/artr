// MIT License
//
// Copyright (c) 2025 AtenTEC
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package logger

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	infoColor    = "\033[1;34m%v\033[0m"
	noticeColor  = "\033[1;33m%v\033[0m"
	warningColor = "\033[1;33m%v\033[0m"
	errorColor   = "\033[1;31m%v\033[0m"
	valColor     = "\033[0;36m%v\033[0m"
)

// Message struct gives you extensibility
type message struct {
	Level   string
	Content []string
	Time    time.Time
}

func Info(msg ...string)  { tlog("INFO", msg) }
func Error(msg ...string) { tlog("ERROR", msg) }

// func Debug(msg ...string)        { tlog("DEBUG", msg) }
func Process(msg ...string)      { tlog("PROCESS", msg) }
func StartSpinner(msg ...string) { tlog("STARTSPINNER", msg) }
func StopSpinner()               { tlog("STOPSPINNER", []string{}) }

func tlog(level string, content []string) {
	printLog(message{
		Level:   level,
		Content: content,
		Time:    time.Now(),
	})
}

func printLog(m message) {
	var prefix string
	switch m.Level {
	case "STARTSPINNER":
		startSpinner(m.Content[0])
		return
	case "STOPSPINNER":
		stopSpinner()
		return
	case "INFO":
		// prefix = fmt.Sprintf(infoColor, "🔴")
		prefix = fmt.Sprintf(infoColor, "[INFO]")
	case "ERROR":
		// prefix = fmt.Sprintf(errorColor, "🔴")
		prefix = fmt.Sprintf(errorColor, "[ERROR]")
	case "DEBUG":
		// prefix = fmt.Sprintf(warningColor, "🔴")
		prefix = fmt.Sprintf(warningColor, "[DEBUG]")
	case "PROCESS":
		// prefix = fmt.Sprintf(warningColor, "🔴")
		prefix = fmt.Sprintf(warningColor, "[EXEC] ")
	default:
		prefix = "⚪"
	}

	fmt.Printf("%s\t [%s] %s\n", prefix, m.Time.Format("15:04:05"), strings.Join(m.Content, ""))
}

// To-Do: delete function after development
func LogJSON(val any) {
	j, _ := json.MarshalIndent(val, "", "    ")
	fmt.Printf("%s\n", j)
}
