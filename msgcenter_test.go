/*
File: errormsg_test.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-12-26 15:58:23

Description: errormsg.go 的单元测试
*/

package msgcenter

import (
	"errors"
	"testing"
)

func TestPrintSuccessMsg(t *testing.T) {
	success := "I am success message"
	PrintSuccessMsg(success)
}

func TestPrintInfoMsg(t *testing.T) {
	info := "I am info message"
	PrintInfoMsg(info)
}

func TestPrintWarnMsg(t *testing.T) {
	warn := "I am warn message"
	PrintWarnMsg(warn)
}

func TestPrintErrorMsg(t *testing.T) {
	err := errors.New("I am error message")
	PrintErrorMsg(err)
}
