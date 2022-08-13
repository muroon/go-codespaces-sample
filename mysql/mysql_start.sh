#!/bin/sh

docker run \
 --name mysql \
 -e MYSQL_ROOT_PASSWORD=mysql \
 -e BIND-ADDRESS=0.0.0.0 \
 -p 3306:3306 \
 -d mysql:latest

