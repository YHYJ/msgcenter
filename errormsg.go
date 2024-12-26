/*
File: errormsg.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-11-03 12:40:35

Description: 输出格式化后的错误信息
*/

package msgcenter

import (
	"github.com/gookit/color"
	"github.com/yhyj/msgcenter/general"
)

// PrintErrorMsg 格式化原始错误信息后输出
//
//   - 输出格式为：[函数名:行号] 错误信息
//
// 参数：
//   - err: 原始错误信息
func PrintErrorMsg(err error) {
	functionName, lineNo := general.GetCallerInfo()
	color.Printf("%s %s\n", general.SecondaryText("[", functionName, ":", lineNo-1, "]"), general.DangerText(err))
}
