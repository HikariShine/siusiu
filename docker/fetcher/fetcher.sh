#!/bin/sh
function getdir(){
    for element in `ls $1`
    do
        dir_or_file=$1"/"$element
        if [ -d $dir_or_file ]
        then
            getdir $dir_or_file
        else
            echo $dir_or_file
        fi
    done
}

if [ $# -ne 1 ];then 
	echo "Usage: fetch.sh <dir>"
else 
	getdir $1 | sed s/.// 
fi 
