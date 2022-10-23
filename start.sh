#!/bin/sh

echo "~~~~~ Starting database ~~~~~~~~~"
docker-compose up -d

TAG="latest"
IMAGE="wagerservice"
echo "~~~~~ Starting build developement entity-server ${IMAGE} ${TAG} ~~~~~~~~~"
docker build -t $IMAGE:$TAG  -f ./dev.dockerfile .
echo '~~~~~ Finish build developement entity-server ~~~~~~~~~'
docker run -d -p 8080:8080 $IMAGE:$TAG
echo '~~~~~ Finish run developement entity-server ~~~~~~~~~'