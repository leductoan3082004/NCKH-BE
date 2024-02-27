docker rmi nckh-user-service:latest
docker build -t nckh-user-service:latest .
docker tag nckh-user-service:latest toan3082004/nckh-user-service:latest
docker push toan3082004/nckh-user-service:latest