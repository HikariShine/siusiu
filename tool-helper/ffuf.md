
## 目录爆破，默认的替换变量名为FUZZ
ffuf -w /path/to/wordlist -u https://target/FUZZ
### 递归式目录爆破
ffuf -w /path/to/wordlist -u https://target/FUZZ -maxtime-job 60 -recursion -recursion-depth 2

## 对http请求头的host字段进行fuzz，假设正常页面返回的响应报文大小4242 
ffuf -w /path/to/vhost/wordlist -u https://target -H "Host: FUZZ" -fs 4242

## 对get传参进行爆破
### 过滤掉响应报文大小为791bytes的请求
ffuf -w /Web-Content/burp-parameter-names.txt:PARAM -w val.txt:VALUE -u http://192.168.1.105/index.php?PARAM=VALUE -fs 791
### 过滤掉响应码为401的请求
ffuf -w /path/to/values.txt -u https://target/script.php?valid_name=FUZZ -fc 401

## 对post传参进行爆破
ffuf -w /path/to/postdata.txt -X POST -d "username=admin\&password=FUZZ" -u https://target/login.php -fc 401

## 暂停fuzz，并重新配置
在ffuf 执行期间，按下enter 进入交互模式，help查看帮助，resume继续，restart 重新开始

