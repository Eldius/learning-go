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
	    echo " => Cleaning $D"
        go clean
        echo "*************"
        cd $CURR_DIR
    fi
done
