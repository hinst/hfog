#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR
target="caddy"
echo "$target control: $1"
if [ "$1" = "start" ]
then
	nohup ./$target 2>>error-log.txt 1>>output-log.txt &
fi

if [ "$1" == "stop" ]
then
	path=$(realpath $target)
	pid=$(pidof $path)
	if [[ $? == 0 ]]
	then
		kill -s SIGINT $pid
		while kill -s 0 $pid
		do 
			sleep 1
		done
		echo stopped
	else
		echo probably not running
	fi
fi