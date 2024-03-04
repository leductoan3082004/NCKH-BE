docker rmi -f toan3082004/nckh:latest
docker rmi -f nckh:latest
docker rmi $(docker images -f "dangling=true" -q)