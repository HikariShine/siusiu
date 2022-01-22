package routers

import (
	"siusiu/controllers"
	"siusiu/pkg/exec"

	"github.com/abiosoft/ishell"
)

//Init 初始化路由
func Init(shell *ishell.Shell) error {
	//第三方工具
	shell.AddCmd(&ishell.Cmd{
		Name: "crawlergo",
		Help: "使用chrome headless模式进行URL收集的浏览器爬虫",
		Func: func(c *ishell.Context) {
			exec.Docker("rickshang/crawlergo", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "nmap",
		Help: "主机发现、端口扫描、服务扫描、版本识别",
		Func: func(c *ishell.Context) {
			exec.Docker("uzyexe/nmap", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "sqlmap",
		Help: "SQL注入攻击工具",
		Func: func(c *ishell.Context) {
			exec.Docker("paoloo/sqlmap", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "dirsearch",
		Help: "目录爆破工具",
		Func: func(c *ishell.Context) {
			exec.Docker("rickshang/dirsearch:0.4.2", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "ffuf",
		Help: "模糊测试工具",
		Func: func(c *ishell.Context) {
			exec.Docker("rickshang/ffuf:1.3.1", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "tool-helper",
		Help: "获取工具的帮助文档",
		Func: func(c *ishell.Context) {
			exec.Docker("rickshang/tool-helper", c.Args)
		},
	})
	//未找到命令时
	shell.NotFound(controllers.NotFoundHandler)

	return nil
}
