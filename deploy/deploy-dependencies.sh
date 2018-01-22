sudo docker stop some-zookeeper
sudo docker stop some-redis
sudo docker stop some-mysql
sudo docker rm some-zookeeper
sudo docker rm some-redis
sudo docker rm some-mysql
sudo docker run -v `pwd`/schema.sql:/docker-entrypoint-initdb.d/schema.sql --name some-mysql -e MYSQL_DATABASE=test -e MYSQL_ROOT_PASSWORD=root -d mysql
sudo docker run --name some-zookeeper --restart always -d zookeeper
sudo docker run --name some-redis -d redis
