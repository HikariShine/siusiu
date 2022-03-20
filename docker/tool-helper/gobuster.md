## 目录爆破
-e 相当于dirsearch中的 --full-url
-x 相当于dirsearch中的 -e 用来指定后缀
-r 跟随重定向
-d 寻找备份文件
-c 设置cookie
-l 显示报文长度
-v 显示详细信息
--exclude-length 排除特定长度的响应
-b, --status-codes-blacklist 排除特定的响应码
gobuster dir -u http://192.168.1.108/site -w ./common.txt -e -x txt,php,html,jsp -t 100 -d -l --random-agent


