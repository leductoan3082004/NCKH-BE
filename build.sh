APP_NAME=nckh-user-service

docker rmi ${APP_NAME}:latest
docker build -t ${APP_NAME}:latest .
docker tag ${APP_NAME}:latest toan3082004/${APP_NAME}:latest
docker push toan3082004/${APP_NAME}:latest