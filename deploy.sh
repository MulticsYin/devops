#!/usr/bin/env bash

echo -e "Start deploy projuct"
killall devops
cd $GOPATH/src

if [ ! -d "/home/multics/go/gopath/src/deploy" ]; then
    echo -e "git clone from https://github.com/MulticsYin/devops.git"
    git clone https://github.com/MulticsYin/devops.git
    cd $GOPATH/src/devops
else
    echo -e "Find old project, git pull from https://github.com/MulticsYin/devops.git"
    cd $GOPATH/src/devops
    git pull
fi

echo -e "Build project now"
go build . && ./devops &

