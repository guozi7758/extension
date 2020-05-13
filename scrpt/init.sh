#!/bin/bash

init(){
	if test -e ./tg_run_$1
	then
		rm tg_run_$1
	fi
	if test -e ./nohup.out
	then
		rm ./nohup.out
	fi
	pid=`pgrep -l tg_run_$1 | awk '{print $1}'`
	if [ "" = "$pid" ];then
		echo "进程为空"
	else
		kill -9 $pid
	fi
}

init
