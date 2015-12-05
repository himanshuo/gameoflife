#!/bin/bash
clear
# determine which directory to watch
DIRNAME="."
MYPID=""
echo "Starting Server"
go run main.go &
MYPID=$!
kill $MYPID
go run main.go &
MYPID=$!

#print variable on a screen
#inotifywait -m -r --event modify $DIRNAME | while read line; do
#	echo "Restarting Server because $line"
#	kill $MYPID
#	go run main.go &
#	MYPID=$!
#done