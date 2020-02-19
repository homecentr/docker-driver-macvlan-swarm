#!/usr/bin/env bash

NAME="homecentr/mdnp:latest"

docker plugin disable $NAME || echo "Plugin was not enabled"
docker plugin rm $NAME || echo "Plugin did not exist"

rm -rf ./rootfs || echo "Directory rootfs did not exist"

docker build . -t rootfsimage
id=$(docker create rootfsimage true)
sudo mkdir ./rootfs
sudo docker export "$id" | sudo tar -x -C ./rootfs
docker rm -vf "$id"
#docker rmi rootfsimage

docker plugin create $NAME .  
docker plugin enable $NAME