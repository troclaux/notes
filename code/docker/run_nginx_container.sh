#!/bin/bash

path="$HOME/Documents/notes/code/docker"
echo "Script path: $path"
cd $path
docker image build -t nginx-image -f ./nginx.Dockerfile .
docker container run -p 8080:80 -v ~/Documents/notes/code/docker/volume:/arthur --name my-nginx -d nginx-image
docker exec -it my-nginx /bin/bash
# docker container run --p 8080:80 -name my-nginx --rm nginx-image
