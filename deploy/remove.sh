APP_NAME=nckh-user-service

docker rmi -f toan3082004/nckh-user-service:latest
docker rmi -f nckh-user-service:latest
docker rmi $(docker images -f "dangling=true" -q)