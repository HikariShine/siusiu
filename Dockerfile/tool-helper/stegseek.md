# 爆破某个隐写文件的密码
stegseek --crack <filename> <wordlist_filename>
stegseek --crack doubletrouble.jpeg rockyou.txt
# 爆破某个隐写文件的同时将隐写内容输出到指定文件中
stegseek --crack <filename> <worldlist_filename> -xf <outputfile>