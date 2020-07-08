package funclog

import (
       "fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

type LogWriter struct{}

func (f LogWriter) Write(p []byte) (n int, err error) {
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	log.Printf("%s:%d %s: %s", filepath.Base(file), line, fnName, p)
	return len(p), nil
}

func New(prefix string) *log.Logger {
	return log.New(LogWriter{}, prefix, 0)
}


func info(depth int) (n string, err error) {
	pc, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	data := fmt.Sprintf("%s:%d %s", filepath.Base(file), line, fnName)
	return data, nil
}

func Parent() (n string, err error) {
        data, err := info(4)

	return data, err
}

func Me() (n string, err error) {
        data, err := info(3)

	return data, err
}

func SimpleStack() (err error) {
     parent, err := Parent()
     me, err := Me()

     data := fmt.Sprintf("%s -> %s", parent, me)

     log.Printf(data)

     return nil
}     