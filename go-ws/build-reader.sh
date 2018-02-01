#!/bin/bash
echo "Building Reader....."
mkdir -p bin/reader
sudo docker run  -v `pwd`:/gows/ -w /gows/src/reader -e "GOBIN=/gows/bin/reader/" -e "GOPATH=/gows" -it gobs /bin/sh -c "go get; CGO_ENABLED=0 GOOS=linux go install -a -ldflags '-extldflags "-static"' ."
if [ ! $? -eq 0 ]; then
    echo "Failed to build reader $?"
fi
sudo cp src/reader/*.html ./bin/reader

