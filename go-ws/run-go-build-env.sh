pushd .
cd ../go-build-system
./create-go-build-system.sh
popd
sudo docker run -p 8080:8080 -p 8000:8000 -w /gows/src/  -v `pwd`:/gows/ -e "GOPATH=/gows" --link some-redis:redis --link some-mysql:mysql --link some-zookeeper:zookeeper -it gobs bash 
