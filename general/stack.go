/*
File: stack.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-11-03 09:28:44

Description: 堆栈相关函数
*/

package general

import "runtime"

// GetCallerInfo 获取调用者信息
//
// 返回：
//   - 调用者的函数名
//   - 调用者所在行号
func GetCallerInfo() (string, int) {
	// runtime.Caller 的参数 skip 是要上升的堆栈数，0 表示 runtime.Caller 的调用者，1 表示 GetCallerInfo 的调用者，以此类推
	pc, _, line, ok := runtime.Caller(2)
	if !ok {
		return "", 0
	}
	function := runtime.FuncForPC(pc).Name()

	return function, line
}
