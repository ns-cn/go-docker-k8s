#!/bin/zsh

docker build -t hello:1.1 .
docekr tag hello:1.1 tangyujun/hello:latest
docker push tangyujun/hello:latest