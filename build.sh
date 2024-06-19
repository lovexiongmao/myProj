#!/bin/bash

abort() {
    echo >&2 '
 #################
 ###  ABORTED  ###
 #################
 '

    echo "An err occured, Exiting..." > &2
    exit 1   
}

trap abort EXIT

set -e 

export GO111MODULE="on"
export GOPROXY="https://goproxy.cn,direct"

RUN_NAME="myProj"
mkdir -p output/bin output/conf  output/docs
cp conf/*.json  output/conf
cp -r docs/* output/docs
go build -a -o output/bin/${RUN_NAME}

trap : 0

echo '
##############
###  DONE  ###
##############
'