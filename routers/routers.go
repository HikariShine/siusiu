package routers

import (
	"log"
	"os"
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
			exec.Docker("instrumentisto/nmap:7.92", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "sqlmap",
		Help: "SQL注入攻击工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/root/.local/share/sqlmap/output/", "-w", "/root/.local/share/sqlmap/output/", "rickshang/sqlmap:1.6.1"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "dirsearch",
		Help: "目录爆破工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/root/reports", "-w", "/root/reports", "rickshang/dirsearch:0.4.2"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "ffuf",
		Help: "模糊测试工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/ffuf:1.3.1"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "whatweb",
		Help: "web指纹识别",
		Func: func(c *ishell.Context) {
			exec.Docker("bberastegui/whatweb", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "tool-helper",
		Help: "获取工具的帮助文档",
		Func: func(c *ishell.Context) {
			exec.Docker("rickshang/tool-helper", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "stegseek",
		Help: "爆破隐写术密码",
		Func: func(c *ishell.Context) {
			//获取当前目录
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "-v", currentDir + ":/steg", "rickdejager/stegseek"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "steghide",
		Help: "隐写术工具",
		Func: func(c *ishell.Context) {
			// docker run -it --rm -v $(pwd):/src -w /src bartimar/steghide info doubletrouble.jpeg
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "-v", currentDir + ":/src", "-w", "/src", "bartimar/steghide"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "http3-client",
		Help: "支持http3的客户端",
		Func: func(c *ishell.Context) {
			exec.Docker("rickshang/http3-client", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "gopherus",
		Help: "ssrf漏洞gopher协议payload生成工具",
		Func: func(c *ishell.Context) {
			exec.Docker("rickshang/gopherus", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "firefox-decrypt",
		Help: "firefox浏览器密码提取工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/firefox-decrypt"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "hydra",
		Help: "弱口令爆破工具",
		Func: func(c *ishell.Context) {
			//1.查询本地是否有hydra
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/thc-hydra"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "wfuzz",
		Help: "web应用fuzz工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "ghcr.io/xmendez/wfuzz", "wfuzz"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "cewl",
		Help: "爬去网站关键字以生成字典",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "nocflame/cewl"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "davtest",
		Help: "webdav利用工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/davtest"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	//未找到命令时
	shell.NotFound(controllers.NotFoundHandler)

	return nil
}
