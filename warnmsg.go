/*
File: warnmsg.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-11-03 12:48:38

Description: 输出格式化后的警告信息
*/

package msgcenter

import (
	"github.com/gookit/color"
	"github.com/yhyj/msgcenter/general"
)

// PrintWarnMsg 格式化原始警告信息后输出
//
//   - 输出格式为：[函数名:行号] 警告信息
//
// 参数：
//   - warn: 原始警告信息
func PrintWarnMsg(warn string) {
	color.Printf("%s %s\n", general.WarningFlag, general.WarnText(warn))
}
