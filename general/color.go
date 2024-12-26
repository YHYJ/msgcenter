/*
File: color.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-11-03 09:26:14

Description: 定义输出样式
*/

package general

import "github.com/gookit/color"

var (
	InfoText    = color.Info.Render    // Info 文本
	NoteText    = color.Note.Render    // Note 文本
	LightText   = color.Light.Render   // Light 文本
	NoticeText  = color.Notice.Render  // Notice 文本
	SuccessText = color.Success.Render // Success 文本
	CommentText = color.Comment.Render // Comment 文本
	PrimaryText = color.Primary.Render // Primary 文本
	WarnText    = color.Warn.Render    // Warn 文本
	DangerText  = color.Danger.Render  // Danger 文本

	SecondaryText = color.Secondary.Render // Secondary 文本
)
