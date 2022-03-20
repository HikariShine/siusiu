#!/bin/bash
if [ "$1" == "shell" ];then
    /bin/bash
elif [ "$1" == "cp" ];then 
    cp $2 /tmp/    
else 
    searchsploit $*
fi
    