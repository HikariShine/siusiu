## 安装 peda 
```
git clone https://github.com/longld/peda.git ~/peda
echo "source ~/peda/peda.py" >> ~/.gdbinit
```
## 启动gdb
gdb -q filename  
-q 表示不显示banner
## 运行可执行程序
```
gdb > r 运行
gdb > r < payload.txt  运行并向标准输入中输入数据
```
## fuzz
生成500个A
```
python -c "print('A'*500)"
```
生成一段特征字符，每4个字符各不相同
```
gdb > pattern create 100
```
自动搜索各个寄存器中特征字符的偏移量
```
gdb > pattern search 
```
## 查看汇编代码
查看所有函数名
```
gdb > info func 
gdb > info functions 
gdb > info functions 正则表达式
```
查看main函数的汇编代码
```
gdb > disas main
```
查看具体的函数名
```
gdb > disas 函数名 
```
## 查看寄存器的值
i register 
## 断点调试
```
gdb > break *0x0x0804873e 
gdb > r 运行
gdb > s 单步调试
gdb > del 1 删除第一个断点
gdb > break *0x08048750 重新下断点
gdb > break 函数名 在某个函数处下断点
```
## 退出
```
gdb > quit
gdb > q
```
参考资料:  
[100个gdb小技巧](https://wizardforcel.gitbooks.io/100-gdb-tips/content/option-format.html)
