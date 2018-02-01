#!/bin/bash
echo "Building Writer....."
mkdir -p build/writer
sudo docker run -v `pwd`:/gows/ -w /gows/src/writer -e "GOBIN=/gows/bin/writer/" -e "GOPATH=/gows" -it gobs /bin/sh -c "go get; CGO_ENABLED=0 GOOS=linux go install -a -ldflags '-extldflags "-static"' ."
if [ ! $? -eq 0 ]; then
    echo "Failed to build writer $?"
fi
sudo cp src/writer/*.html ./bin/writer
