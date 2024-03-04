APP_NAME=nckh
DEPLOY_CONNECT=root@165.22.54.162

./deploy/remove.sh

docker-compose -H "ssh://${DEPLOY_CONNECT}" down || true

ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} 'bash -s' < ./deploy/remove.sh
./deploy/remove.sh

echo "Deploying..."

docker-compose -H "ssh://${DEPLOY_CONNECT}"  up --detach
echo "Done"