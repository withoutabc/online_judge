package util

import (
	"fmt"
	"runtime"
)

func Find() {
	// 输出当前代码的位置信息
	_, file, line, _ := runtime.Caller(0)
	fmt.Printf("第 1 次调用：%s 的第 %d 行。\n", file, line)
	// 获取调用栈信息
	pc := make([]uintptr, 64)
	n := runtime.Callers(1, pc)
	// 循环输出每一层调用者的信息
	for i, c := 0, 2; i < n; i, c = i+1, c+1 {
		// 获取当前层调用者的位置信息
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		// 输出位置信息和调用次数
		fmt.Printf("第 %d 次调用：%s 的第 %d 行。\n", c, file, line)
	}
}
