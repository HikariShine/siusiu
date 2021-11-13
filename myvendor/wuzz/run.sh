#!/bin/bash
base_path=$HOME/src
install_path=$base_path/wuzz #程序安装目录

function download {
    git clone --depth 1 https://github.com.cnpmjs.org/asciimoo/wuzz.git $1 && cd $1 && go build ./... && go build
}

#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #2.如果不存在就下载
    echo "[*] download..."
    download $install_path
    echo "[*] download success"
fi
current=$(pwd)
cd $install_path && git reset --hard && git pull origin master && go build ./... && go build && cd $current

if [ ! -x $install_path/wuzz ];then
    chmod +x $install_path/wuzz
fi


$install_path/wuzz $*
