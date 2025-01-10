package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func getGoroutineID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	id := int64(-1)
	fmt.Sscanf(string(buf[:n]), "goroutine %d", &id)
	fmt.Println()
	return id
}

const ALL = "ALL"
const TRACE = "TRACE"
const DEBUG = "DEBUG"
const INFO = "INFO"
const WARN = "WARN"
const ERROR = "ERROR"
const FATAL = "FATAL"
const OFF = "OFF"

func Println(a ...any) (int, error) {
	callerFuncName := "unknow"
	pc, file, line, ok := runtime.Caller(1)
	fileNameSlice := strings.Split(file, "/")
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			callerFuncName = fn.Name()
		}
		temp := ""
		for _, v := range a {
			a, _ := v.(string)
			temp += a + " "
		}
		fmt.Printf(" %s [%d]	--- \033[1;34;40m %s:%s:%d \033[0m		:",
			time.Now().Format("2006-01-02 15:04:05.999"), getGoroutineID(), fileNameSlice[len(fileNameSlice)-1], callerFuncName, line)
		return fmt.Println(a...)
	}
	return 0, nil
}

func Printf(format string, a ...any) (int, error) {
	callerFuncName := "unknow"
	pc, file, line, ok := runtime.Caller(1)
	fileNameSlice := strings.Split(file, "/")
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			callerFuncName = fn.Name()
		}
		temp := ""
		for _, v := range a {
			a, _ := v.(string)
			temp += a + " "
		}
		fmt.Printf(" %s [%d]	--- \033[1;34;40m %s:%s:%d \033[0m		:",
			time.Now().Format("2006-01-02 15:04:05.999"), getGoroutineID(), fileNameSlice[len(fileNameSlice)-1], callerFuncName, line)
		return fmt.Printf(format, a...)
	}
	return 0, nil
}
