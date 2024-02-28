APP_NAME=nckh-be-post

docker rmi -f toan3082004/${APP_NAME}:latest
docker rmi -f ${APP_NAME}:latest
docker rmi $(docker images -f "dangling=true" -q)