# siusiu （suitesuite）
一个用来管理suite 的suite，志在将渗透测试工程师从各种安全工具的学习和使用中解脱出来，减少渗透测试工程师花在安装工具、记忆工具使用方法上的时间和精力。

## Features

siusiu提供了一个shell控制台，通过该控制台，可以：

- 查看第三方安全工具列表
- 自动安装第三方安全工具
- 运行第三方安全工具
- 查看第三方安全工具的说明文档与使用样例（通过demos命令）

同时siusiu也支持非交互模式，便于siusiu被其他程序调用,例如:siusiu exec help

## Usage：
```
siusiu > help

Commands:
  clear                clear the screen
  crawlergo            使用chrome headless模式进行URL收集的浏览器爬虫
  dirsearch            目录爆破工具
  exit                 exit the program
  ffuf                 模糊测试工具
  firefox-decrypt      firefox浏览器密码提取工具
  gopherus             ssrf漏洞gopher协议payload生成工具
  help                 display help
  http3-client         支持http3的客户端
  hydra                弱口令爆破工具
  nmap                 主机发现、端口扫描、服务扫描、版本识别
  sqlmap               SQL注入攻击工具
  steghide             隐写术工具
  stegseek             爆破隐写术密码
  tool-helper          获取工具的帮助文档
  whatweb              web指纹识别
```

## Installation:

```
wget https://gitee.com/nothing-is-nothing/siusiu/raw/master/setup.sh
chmod +x setup.sh
./setup.sh
siusiu
```

## Screenshots

如果用户未安装pocsuite3，则自动下载 pocsuite3,然后自动运行 
![avatar](https://img-blog.csdnimg.cn/20211006160456729.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)

在siusiu控制台中运行sqlmap和dirsearch
![avatar](https://img-blog.csdnimg.cn/20211006160557298.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)


## Tested On  

- MacOS
- CentOS7
- Ubuntu

## Develop  

如果您有其他好的安全工具也想集成到siusiu中，可以按照如下步骤操作：  
step1.在siusiu安装目录（$HOME/src/siusiu）下创建对应的工具目录（建议以工具名命名,例如：dirsearch），并在该目录下创建该工具的shell脚本 run.sh，例如：

```shell
#!/bin/bash
base_path=$HOME/src
dirsearch_path=$base_path/dirsearch

function download {
    git clone https://github.com.cnpmjs.org/maurosoria/dirsearch.git $1
    cd $1
    pip3 install -r requirements.txt
}

#1.检查程序目录是否存在
if [ ! -d $dirsearch_path ]; then
    #2.如果不存在就下载
    echo "[*] download dirsearch..."
    download $dirsearch_path
fi
#运行dirsearch
python3 $dirsearch_path/dirsearch.py $*
```
step2. 在config.json 配置文件中添加对应工具，例如：
```
        {
            "Name": "dirsearch",
            "Help": "目录扫描器",
            "Run": "dirsearch/run.sh"
        },
```
其中name为工具名，help为工具描述，run为该工具的run.sh在myvendor目录下的相对路径

## 为工具编写demo文档

不知道你是否也曾有过这样的烦恼：每天疲于学习各种工具的使用方法,当真正需要使用某个工具的时候，却一时半会儿想不起某个工具怎么用，这时你翻开了你的笔记本，找呀找，终于找到了以前的笔记。  
关于这个问题，siusiu提供一种解决方案：将工具的使用文档或者常用demo集成在shell控制台中，需要时直接通过命令：demos+工具名 查看即可。  
你可以将你常用的一些命令demo，以markdown文档的方式写在 $HOME/src/siusiu/myvendor/demos 目录下，siusiu控制台会自动读取该目录。  
例如为sqlmap编写常用demo文档：

```markdown
# sqlmap demoes

```shell
# -m 批量扫描 —batch 全部采用默认行为，不向用户请求y/n,并且使用随机的user—agnet
sqlmap -m temp2.txt --batch --random-agent> result.txt

# 尝试获取所有数据库名
sqlmap -u url --dbs —-random-agent --batch

# 获取表名
sqlmap -u url --tables —-random-agent --batch

# 尝试获取所有用户:
sqlmap -u url --users --random-agent --batch

# 尝试获取账号密码:
sqlmap -u url --password --random-agent --batch

# 尝试获取当前用户:
sqlmap -u url --current-user --random-agent --batch

# 测试当前用户权限:
sqlmap -u url --is-dba --random-agent --batch

# 尝试写入木马,getshell
sqlmap -u url --os-shell --random-agent --batch

# 执行指定的sql语句
sqlmap.py -u url -v 1 --sql-query 'select top 20 * from City'
```

在siusiu控制台中通过 demos sqlmap.md 即可查看该文档。  

## QA

问：我喜欢作者怎么办？  
答：⁄(⁄ ⁄•⁄ω⁄•⁄ ⁄)⁄  

问：能不能一键日卫星？  
答：~~不能，至少目前现阶段不可以。抱歉。么么哒~~ 可以一键日卫星，我的小可爱。  

问：How to build and install siusiu on raspberry pi ?  
答：首先，请不要说英文，请用普通话。然后， I can not replay your question you just mentioned in mandarin excuse me. If you want to run it on embed platforms like raspberry pi 3 model b, you must need to know that these platforms have a ARM core, for example, pi 3b's cpu is Contex-A53 with a ARM v8 architecture, so, set GOOS=linux and GOARCH=aarch64 then run go build main.go in your shell, enjoyed it !

