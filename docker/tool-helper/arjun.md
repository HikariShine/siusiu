//fuzz get传参
arjun -u https://api.example.com/endpoint

//fuzz form格式的post传参
arjun -u https://api.example.com/endpoint -m POST

//fuzz json格式的post传参
arjun -u https://api.example.com/endpoint -m JSON

//fuzz xml格式的post传参
arjun -u https://api.example.com/endpoint -m XML

//批量fuzz
arjun -i targes.txt

//以json格式导出结果
arjun -u https://api.example.com/endpoint -oJ result.json

//推送结果到burpsuite
-oB 127.0.0.1:8080

//指定注入点
arjun -u https://api.example.com/endpoint -m JSON --include='{"root":{"a":"b",$arjun$}}'
arjun -u https://api.example.com/endpoint -m XML --include='<?xml><root>$arjun$</root>'

//控制发包速率
arjun  -u https://api.example.com/endpoint --stable


