package controllers

import (
	"fmt"
	"os"
	"siusiu/pkg/exec"
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

//GetShellPrompt 获取shell提示符
func GetShellPrompt() string {
	logo := "siusiu"
	pwd, err := os.Getwd()
	if err != nil {
		logrus.Error("os.Getwd failed,err:", err)
		return fmt.Sprintf("%s > ", logo)
	}
	return fmt.Sprintf("%s:%s > ", color.YellowString(logo), color.GreenString(pwd))
}

//NotFoundHandler 未找到命令时处理函数
func NotFoundHandler(c *ishell.Context) {
	args := c.RawArgs
	input := strings.Join(c.RawArgs, " ")
	if args[0] == "cd" {
		var dir string
		switch len(args) {
		case 1:
			dir = os.Getenv("HOME")
		default:
			dir = args[1]
		}
		if err := os.Chdir(dir); err != nil {
			logrus.Error("os.Chdir failed,err:", err)
			return
		}
		c.SetPrompt(GetShellPrompt())
		return
	}
	exec.CmdExec("/bin/bash", "-c", input)
}
