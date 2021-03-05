#!/bin/bash

sleep 30 

FILE=/root/started


if [ ! -f "$FILE" ]; then
	touch /root/started
	/bin/odin-engine -id master-node raft0 &
fi


if test -f "$FILE"; then
    /bin/odin
fi

