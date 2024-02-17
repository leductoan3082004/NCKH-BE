APP_NAME=nckh
DEPLOY_CONNECT=root@165.22.54.162

docker rm -f ${APP_NAME}
docker image rm ${APP_NAME}
docker-compose -H "ssh://${DEPLOY_CONNECT}" down || true

ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} 'bash -s' < ./deploy/remove.sh
./deploy/remove.sh

echo "Docker building..."

docker build -t ${APP_NAME}:latest .
echo "Docker pushing..."
docker tag ${APP_NAME}:latest toan3082004/${APP_NAME}:latest
docker push toan3082004/${APP_NAME}:latest

echo "Deploying..."

docker-compose -H "ssh://${DEPLOY_CONNECT}"  up --detach
echo "Done"