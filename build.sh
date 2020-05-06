#!/bin/bash

CURR_DIR=${PWD}

function build {
    f=$1

}

RESULTS=()

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
            go test ./... && go build && execution_result=" => Build $D with success" || execution_result=" => Build $D failed"
        fi
        echo $execution_result
        RESULTS+=("$execution_result")
        echo "*************"
        cd $CURR_DIR
    fi
done
echo ""
echo ""
echo "** EXECUTION RESULTS: **********"
printf "%s\n" "${RESULTS[@]}"
echo "********************************"
echo ""
