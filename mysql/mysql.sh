#!/bin/sh

docker run \
 --name mysql \
 -e MYSQL_ROOT_PASSWORD=mysql \
 -e BIND-ADDRESS=0.0.0.0 \
 -p 3306:3306 \
 -d mysql:latest

mysql -u root -p -h 127.0.0.1 -P 3306

