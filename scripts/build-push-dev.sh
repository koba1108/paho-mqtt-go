#!/bin/bash

if [ "$1" = "" ]; then
  echo "please enter image tag version. example=[0.0.1]"
  exit
fi

DockerfilePath="./build/docker/Dockerfile"
docker build -t gcr.io/ykoba-mumbai/paho-mqtt-go:"$1" -f "$DockerfilePath" . &&
  docker push gcr.io/ykoba-mumbai/paho-mqtt-go:"$1"
