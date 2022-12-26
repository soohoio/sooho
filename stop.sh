make localnet-clean
docker rm $(docker ps -aq)
docker image rm $(docker image ls -qa)