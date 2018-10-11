
docker ps -a | awk 'match($2,"practice00") {print $1}' | xargs docker rm
