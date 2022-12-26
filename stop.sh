make localnet-clean
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
docker image rm $(docker image ls -qa)