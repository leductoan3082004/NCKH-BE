APP_NAME=nckh-be-post
./deploy/remove.sh

docker rmi ${APP_NAME}:latest
docker build -t ${APP_NAME}:latest .
docker tag ${APP_NAME}:latest toan3082004/${APP_NAME}:latest
docker push toan3082004/${APP_NAME}:latest