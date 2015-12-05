#!/bin/bash
clear
# determine which directory to watch
DIRNAME="."
trap ctrl_c INT
echo "Starting Server"
go build main.go
./main &

function ctrl_c() {
        echo "** Trapped CTRL-C"
        killall main
}

#print variable on a screen
inotifywait -m -r --event modify $DIRNAME | while read line; do
	echo "Restarting Server because $line"
	kill %1
	go build main.go
	./main &
done

