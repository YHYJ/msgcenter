/*
File: stack.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-11-03 09:28:44

Description: 堆栈相关函数
*/

package general

import (
	"path/filepath"
	"runtime"
	"strings"
)

// GetCallerInfo 获取调用者信息
//
// 返回：
//   - 调用者所在文件名（不带后缀）
//   - 调用者的函数名
//   - 调用者所在行号
func GetCallerInfo() (string, string, int) {
	// runtime.Caller 的参数 skip 是要上升的堆栈数，0 表示 runtime.Caller 的调用者，1 表示 GetCallerInfo 的调用者，以此类推
	pc, fullFilePath, line, ok := runtime.Caller(2)
	if !ok {
		return "", "", 0
	}

	// 获取文件名
	file := strings.Split(filepath.Base(fullFilePath), ".")[0]

	// 获取函数名
	functionSplit := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	function := functionSplit[len(functionSplit)-1]

	return file, function, line
}
