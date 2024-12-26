/*
File: successmsg.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-11-03 12:50:35

Description: 输出格式化后的成功信息
*/

package msgcenter

import (
	"github.com/gookit/color"
	"github.com/yhyj/msgcenter/general"
)

// PrintSuccessMsg 格式化原始成功信息后输出
//
//   - 输出格式为：成功信息
//
// 参数：
//   - success: 原始成功信息
func PrintSuccessMsg(success string) {
	color.Printf("%s\n", general.SuccessText(success))
}
