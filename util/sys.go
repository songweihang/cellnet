package util

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// 给定打印层数,一般3~5覆盖你的逻辑及封装代码范围
func StackToString(count int) string {

	const startStack = 2

	var sb strings.Builder

	for i := startStack; i < startStack+count; i++ {
		_, file, line, ok := runtime.Caller(i)

		var str string

		if ok {
			str = fmt.Sprintf("%s:%d", filepath.Base(file), line)
		} else {
			str = "??"
		}

		// 折叠??
		if str != "??" {
			if i > startStack {
				sb.WriteString(" -> ")
			}

			sb.WriteString(str)
		} else {
			break
		}

	}

	return sb.String()
}

func ErrorCatcher(errFunc func(error)) {

	err := recover()

	switch err.(type) {
	// 运行时错误
	case interface {
		RuntimeError()
	}:

		// 继续外抛， 方便调试
		panic(err)

	case error:
		errFunc(err.(error))
	case nil:
	default:
		panic(err)
	}
}