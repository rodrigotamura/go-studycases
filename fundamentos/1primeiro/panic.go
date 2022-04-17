package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func soma() int64 {
	return 1 + 2
}

func main() {
	time.RFC3339
	defer func() {
		if err := recover(); err != nil {
			// fmt.Print(strings.HasPrefix(string(runtime.FuncForPC(reflect.ValueOf(err).Pointer()).Name()), "main.main.func"))
			fmt.Print(string(runtime.FuncForPC(reflect.ValueOf(err).Pointer()).Name()), "main.main.func")
		}
	}()

	panic(soma())
}
