#!/bin/bash

CURR_DIR=${PWD}

function build {
    f=$1

}

for D in *
do
    if [ -d "${D}" ]; then
        cd $CURR_DIR/$D
        echo "*************"
	    echo " => Building $D"
        FILE="$CURR_DIR/$D/build.sh"
        if test -f "$FILE"; then
            source $FILE
        else
            go test ./... && go build && echo " => Build $D with success" || echo " => Build $D failed"
        fi
        echo "*************"
        cd $CURR_DIR
    fi
done
