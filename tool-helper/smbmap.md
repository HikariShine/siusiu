基本使用

smbmap -H 192.168.31.74
smbmap -H 192.168.31.74 -u msfadmin -p msfadmin

因为具有可读权限，所以我们可以进一步访问共享目录下的文件：

python3 smbmap.py -H 192.168.31.74 -r <dir>

下载文件：

python3 smbmap.py -H 192.168.31.74 --download Reports/annual.txt

如果具有可写权限，还可以上传文件:

python3 smbmap.py -H 192.168.31.74 --upload README.md  Reports/README.md

