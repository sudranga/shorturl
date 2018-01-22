#!/bin/bash
echo "Building GO Build system...."
pushd .
cd ../go-build-system
./create-go-build-system.sh
popd

sudo mkdir -p ./bin/reader
sudo mkdir -p ./bin/writer

echo "Building Reader....."
sudo docker run  -v `pwd`:/gows/ -w /gows/src/reader -e "GOBIN=/gows/bin/reader/" -e "GOPATH=/gows" -it gobs /bin/sh -c "CGO_ENABLED=0 GOOS=linux go install -a -ldflags '-extldflags "-static"' ."
if [ ! $? -eq 0 ]; then
    echo "Failed to build reader $?"
fi
sudo cp src/reader/*.html ./bin/reader

echo "Building Writer....."
sudo docker run -v `pwd`:/gows/ -w /gows/src/writer -e "GOBIN=/gows/bin/writer/" -e "GOPATH=/gows" -it gobs /bin/sh -c "CGO_ENABLED=0 GOOS=linux go install -a -ldflags '-extldflags "-static"' ."
if [ ! $? -eq 0 ]; then
    echo "Failed to build writer $?"
fi
sudo cp src/writer/*.html ./bin/writer
