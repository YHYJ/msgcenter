/*
File: infomsg.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-11-03 12:50:35

Description: 输出格式化后的常规信息
*/

package msgcenter

import (
	"github.com/gookit/color"
	"github.com/yhyj/msgcenter/general"
)

// PrintInfoMsg 格式化原始常规信息后输出
//
//   - 输出格式为：常规信息
//
// 参数：
//   - info: 原始常规信息
func PrintInfoMsg(info string) {
	color.Printf("%s\n", general.LightText(info))
}
