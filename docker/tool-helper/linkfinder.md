## 输出到html文件
python linkfinder.py -i https://example.com/1.js -o results.html

## CLI/STDOUT 输出（不使用 jsbeautifier，这使得它非常快）：
python linkfinder.py -i https://example.com/1.js -o cli

## 分析整个域及其 JS 文件：
python linkfinder.py -i https://example.com -d

## Burp 输入（在目标中选择要保存的文件，右键单击，Save selected items将该文件作为输入提供）：
python linkfinder.py -i burpfile -b

## 枚举 JavaScript 文件的整个文件夹，同时查找以 /api/ 开头的端点，最后将结果保存到 results.html：
python linkfinder.py -i 'Desktop/*.js' -r ^/api/ -o results.html

