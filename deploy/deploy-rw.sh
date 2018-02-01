sudo docker kill reader
sudo docker kill writer
sudo docker rm reader
sudo docker rm writer
sudo docker run -d -p 8080:8080 --name reader --link some-redis:redis --link some-mysql:mysql --link some-zookeeper:zookeeper  reader
sudo docker run -d -p 8000:8000 --name writer --link some-redis:redis --link some-mysql:mysql --link some-zookeeper:zookeeper  writer 
