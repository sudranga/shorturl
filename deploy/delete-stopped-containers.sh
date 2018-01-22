sudo docker ps -aq -f status=exited | xargs sudo docker rm
