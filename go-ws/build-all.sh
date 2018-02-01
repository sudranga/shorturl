#!/bin/bash
echo "Building GO Build system...."
pushd .
cd ../go-build-system
./create-go-build-system.sh
popd

./build-reader.sh
./build-writer.sh
